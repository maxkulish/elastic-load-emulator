package main

import (
	"ElasticLoad/config"
	"ElasticLoad/elasticd"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := elasticd.NewElasticConn(cfg)
	if err != nil {
		log.Fatal(err)
	}

	event := PrepareEvent()

	for i := 1; i <= 100; i++ {
		err := PutIndex(i, event, conn.Client)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Printf("Stopped at: %d", i)
			break
		}
		log.Printf("sent index: %d", i)
	}

	log.Printf("finished. Started: %d, planned finish: %d", 0, 100)

}

func PutIndex(id int, event map[string]interface{}, client *elastic.Client) error {

	event["eventTime"] = time.Now().Format(time.RFC3339)

	// Add a document to the index
	_, err := client.Index().
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

func PrepareEvent() map[string]interface{} {

	// Open our jsonFile
	jsonFile, err := os.Open("./index.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(byteValue), &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
