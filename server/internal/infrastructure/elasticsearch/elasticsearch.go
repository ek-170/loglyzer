package logprocessor

import (
    "github.com/olivere/elastic/v7"
)

func CreateElasticsearchClient() (*elastic.Client, error) {
    // Elasticsearchのクライアントを作成
    client, err := elastic.NewClient(elastic.SetSniff(false))
    if err != nil {
        return nil, err
    }
    return client, nil
}
