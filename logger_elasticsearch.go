package main

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

type ElasticsearchLogger struct {
	Url   string
	Index string
}

func (l *ElasticsearchLogger) Log(data []CSPRequest) {
	client, err := elastic.NewClient(elastic.SetURL(l.Url), elastic.SetSniff(false))
	if err != nil {
		log.Print(err.Error())
		return
	}

	bulk := client.Bulk().Index(l.Index).Type("_doc")
	for _, d := range data {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(d.Report))
	}

	res, err := bulk.Do(context.TODO())
	if err != nil {
		log.Print(err.Error())
		return
	}

	if res.Errors {
		log.Print("Bulk errors.")
	}
}
