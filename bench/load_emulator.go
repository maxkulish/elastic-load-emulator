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
	indexFile = "/index.json"
)

type LoadEmulator struct {
	Config *config.Config
	Client *elasticd.Elastic
}

func NewLoadEmulator() (*LoadEmulator, error) {
	// load config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("error while creating NewConfig", err)
	}

	el, err := elasticd.NewElasticConn(cfg)
	if err != nil {
		log.Fatal("error while creating NewElasticConnection", err)
	}

	return &LoadEmulator{
		Config: cfg,
		Client: el,
	}, nil

}

func (le *LoadEmulator) RunPutIndexEmulator(start, finish int) {

	// Prepare event once for all inputs
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("can't create example index file")
	}

	// Open index jsonFile
	jsonFile, err := os.Open(dir + indexFile)
	if err != nil {
		log.Fatalf("error loading index file <%s>. Error: %+v", dir+indexFile, err)
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

func CreateExampleIndexFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("can't create example index file")
	}

	// Create file
	file, err := os.Create(dir + indexFile)
	if err != nil {
		log.Fatal("error while creating index.json file", err)
	}
	defer file.Close()

	index := `
{
 "created_at": "Wed Oct 10 20:19:24 +0000 2018",
 "id": 1050118621198921728,
 "id_str": "1050118621198921728",
 "text": "To make room for more expression, we will now count all emojis as equal—including those with gender‍‍‍ ‍‍and skin t… https://t.co/MkGjXf9aXm",
 "user": {},  
 "entities": {}
}`

	// Write data to file
	_, err = file.Write([]byte(index))
	if err != nil {
		log.Fatal("can't write data to index.json file", err)
	}
}
