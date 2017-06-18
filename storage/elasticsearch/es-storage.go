package main

import (
	"context"
	"fmt"
	"time"

	"github.com/doubtingben/beeper/result"
	"gopkg.in/olivere/elastic.v5"
)

func main() {
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	// Ping with context
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s", code, info.Version.Number)
	esversion, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s", esversion)

	// Check if Index exits, if not, create it
	t := time.Now()
	index := "beep-" + t.Format("2006-01-02")
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		fmt.Println("Creating index")
		createIndex, err := client.CreateIndex(index).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("Index creation not ackowledged")
		}
	} else {
		fmt.Println("Index already exists")
	}

	// Index a tweet (using JSON serialization)
	result1 := result.Result{HTTPStatus: "200", HTTPStatusCode: 200,
		InstanceName: "someInstance", DNSLookup: 43, ServerProcessing: 65}
	put1, err := client.Index().
		Index(index).
		Type("result").
		Id("1").
		BodyJson(result1).
		Do(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed reult %s to index %s, type %s\n",
		put1.Id, put1.Index, put1.Type)

	// Flush to make sure the documents got written.
	_, err = client.Flush().Index(index).Do(ctx)
	if err != nil {
		panic(err)
	}
}
