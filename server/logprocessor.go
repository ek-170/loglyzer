package logprocessor

import (
    "bufio"
    "log"
    "os"
    "strings"
    "sync"

    "github.com/olivere/elastic/v7"
)

const bulkSize = 100 // バルクサイズ（適宜調整）

type LogEntry struct {
    Line int    `json:"line"`
    Log  string `json:"log"`
}

func ProcessLogFile(logFilePath string, client *elastic.Client) error {
    // ファイルを開く
    file, err := os.Open(logFilePath)
    if err != nil {
        return err
    }
    defer file.Close()

    // ElasticsearchへのBulk Request用のチャネル
    bulkChannel := make(chan *elastic.BulkService, bulkSize) // バッファサイズを調整

    // Elasticsearchへの送信用のgoroutineを制御するためのWaitGroup
    var wg sync.WaitGroup

    // ファイルを読み込み、ログをElasticsearchに送信
    scanner := bufio.NewScanner(file)
    var lineNumber int
    var multiLineLog []string

    for scanner.Scan() {
        line := scanner.Text()
        lineNumber++

        if strings.TrimSpace(line) == "" {
            // 空行は無視
            continue
        }

        if strings.HasPrefix(line, "ERROR") {
            // エラーログの開始
            multiLineLog = append(multiLineLog, line)
        } else if len(multiLineLog) > 0 {
            // エラーログの終了
            multiLineLog = append(multiLineLog, line)
            logEntry := LogEntry{Line: lineNumber, Log: strings.Join(multiLineLog, "\n")}
            SendToBulkChannel(bulkChannel, logEntry)
            multiLineLog = nil
        } else {
            // 通常のログ
            logEntry := LogEntry{Line: lineNumber, Log: line}
            SendToBulkChannel(bulkChannel, logEntry)
        }

        // バルクサイズごとにBulk Requestを送信
        if len(bulkChannel) >= bulkSize {
            SendBulkRequests(client, bulkChannel, &wg)
        }
    }

    // 最後のバルク Request を送信
    SendBulkRequests(client, bulkChannel, &wg)

    if err := scanner.Err(); err != nil {
        return err
    }

    // ElasticsearchへのBulk Requestの送信完了待ち
    wg.Wait()
    return nil
}

func SendToBulkChannel(bulkChannel chan *elastic.BulkService, logEntry LogEntry) {
    bulkService := elastic.NewBulkIndexRequest().Index("myapp-logs").Type("_doc").Doc(logEntry)
    bulkChannel <- bulkService
}

func SendBulkRequests(client *elastic.Client, bulkChannel chan *elastic.BulkService, wg *sync.WaitGroup) {
    defer wg.Done()

    // ElasticsearchへのBulk Requestを作成
    ctx := context.Background()
    bulk := client.Bulk()
    for i := 0; i < bulkSize; i++ {
        select {
        case req, ok := <-bulkChannel:
            if !ok {
                return // チャネルがクローズされたら終了
            }
            bulk.Add(req)
        default:
            break // チャネルが空の場合は終了
        }
    }

    // ElasticsearchへのBulk Requestを送信
    _, err := bulk.Do(ctx)
    if err != nil {
        log.Printf("ElasticsearchへのBulk Requestの送信に失敗しました: %v", err)
    }
}