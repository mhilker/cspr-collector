package main

import (
	"context"
	"github.com/olivere/elastic"
	"log"
)

type ElasticsearchOutput struct {
	Url   string
	Index string
}

func (o *ElasticsearchOutput) Write(data []CSPRequest) {
	client, err := elastic.NewClient(elastic.SetURL(o.Url), elastic.SetSniff(false))
	if err != nil {
		log.Print(err.Error())
		return
	}

	bulk := client.Bulk().Index(o.Index).Type("_doc")
	for _, d := range data {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(d.Report))
	}

	res, err := bulk.Do(context.TODO())
	if err != nil {
		log.Print(err.Error())
		return
	}

	if !res.Errors {
		return
	}

	log.Print("Bulk errors.")
	for _, items := range res.Items {
		for _, i := range items {
			if i.Error != nil {
				log.Print(i.Error.Reason)
			}
		}
	}
}
