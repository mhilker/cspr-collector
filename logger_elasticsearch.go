package main

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

type ElasticsearchLogger struct {
	Url string
}

func (l *ElasticsearchLogger) Log(data []CSPRequest) {
	client, err1 := elastic.NewClient(elastic.SetURL(l.Url), elastic.SetSniff(false))
	if err1 != nil {
		log.Println(err1.Error())
		return
	}

	bulk := client.Bulk().Index("csp-violations").Type("_doc")
	for _, d := range data {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(d.Report))
	}

	res, err2 := bulk.Do(context.TODO())
	if err2 != nil {
		log.Println(err2.Error())
		return
	}

	if res.Errors {
		log.Println("Bulk errors")
	}
}
