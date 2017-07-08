package store

import (
	"context"
	"fmt"
	"time"

	"github.com/doubtingben/beeper/result"
	"gopkg.in/olivere/elastic.v5"
)

// Save to store the result in elasticsearch
func Save(result *result.Result) error {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		return err
	}

	// Ping with context
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		return err
	}
	fmt.Printf("Elasticsearch version %s", esversion)

	// Check if Index exits, if not, create it
	t := time.Now()
	index := "beep-" + t.Format("2006-01-02")
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		return err
	}
	if !exists {
		fmt.Println("Creating index")
		createIndex, err := client.CreateIndex(index).Do(ctx)
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			fmt.Println("Index creation not ackowledged")
		}
	} else {
		fmt.Println("Index already exists")
	}

	// Index a result (using JSON serialization)
	put1, err := client.Index().
		Index(index).
		Type("result").
		Id("1").
		BodyJson(result).
		Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("Indexed reult %s to index %s, type %s\n",
		put1.Id, put1.Index, put1.Type)

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index(index).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}
