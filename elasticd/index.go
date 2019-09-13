package elasticd

import (
	"ElasticLoad/config"
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

type Elastic struct {
	Client *elastic.Client
	event  map[string]interface{}
}

func NewElasticConn(conf *config.Config) (*Elastic, error) {

	ctx := context.Background()

	hostURL := fmt.Sprintf("%s://%s:%s",
		conf.ElasticParams.Proto,
		conf.ElasticParams.Host,
		conf.ElasticParams.Port)

	client, err := elastic.NewClient(elastic.SetURL(hostURL))
	if err != nil {
		log.Fatal("elastic connection error: ", err)
		return nil, err
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(hostURL).Do(ctx)
	if err != nil {
		log.Fatalf("connection created but elastic response: %d. %+v", code, err)
		return nil, err
	}

	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	return &Elastic{Client: client}, nil
}

func (e *Elastic) PutIndex(id int, event map[string]interface{}) error {

	event["eventTime"] = time.Now().Format(time.RFC3339)

	// Add a document to the index
	_, err := e.Client.Index().
		Index("bench").
		Id(strconv.Itoa(id)).
		BodyJson(event).
		Refresh("false").
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	return err
}

func (e *Elastic) DeleteIndex(idx string) error {
	_, err := e.Client.DeleteIndex(idx).Do(context.Background())
	if err != nil {
		// Handle error
		log.Fatal("Delete index error:", err)
	}
	return err
}
