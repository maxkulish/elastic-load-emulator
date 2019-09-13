package bench

import (
	"ElasticLoad/config"
	"ElasticLoad/elasticd"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	indexFile = "./index.json"
)

type LoadEmulator struct {
	Config *config.Config
	Client *elasticd.Elastic
}

func NewLoadEmulator() (*LoadEmulator, error) {
	// load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	el, err := elasticd.NewElasticConn(cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &LoadEmulator{
		Config: cfg,
		Client: el,
	}, nil

}

func (le *LoadEmulator) RunPutIndexEmulator(start, finish int) {

	// Prepare event once
	event := prepareEvent()

	sTime := time.Now()

	// run one loop
	for i := start; i <= finish; i++ {
		err := le.Client.PutIndex(i, event)
		//time.Sleep(2 * time.Millisecond)
		if err != nil {
			log.Printf("Stopped at: %d", i)
			break
		}
		log.Printf("sent index: %d", i)
	}

	log.Printf("Done! Spent: %f seconds", time.Since(sTime).Seconds())
	log.Printf(" Started: %d, finished: %d", start, finish)
}

func prepareEvent() map[string]interface{} {

	// Open our jsonFile
	jsonFile, err := os.Open(indexFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalf("error loading index file <%s>. Error: %+v", indexFile, err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("error while reading json data from file")
	}

	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
