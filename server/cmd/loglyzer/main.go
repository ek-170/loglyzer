package main

import (
    "fmt"
    "log"
    "github.com/ek-170/loglyzer/internal/logprocessor"
)

const logFilePath = "/usr/app/logs/myapp.log"

func main() {
    // Elasticsearchのクライアントを作成
    client, err := logprocessor.CreateElasticsearchClient()
    if err != nil {
        log.Fatalf("Elasticsearchクライアントの作成に失敗しました: %v", err)
    }

    err = logprocessor.ProcessLogFile(logFilePath, client)
    if err != nil {
        log.Fatalf("ログファイルの処理に失敗しました: %v", err)
    }

    fmt.Println("ログの送信が完了しました")
}
