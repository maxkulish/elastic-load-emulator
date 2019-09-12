package service

import (
	"ElasticLoad/config"
	"log"
	"time"
)

type Benchmark struct {
	Elastic *elastic.Elastic
	Config  *config.Config
}

func NewBenchmark() (*Benchmark, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	conn, err := elastic.NewElasticConn(cfg)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Benchmark{
		Elastic: conn,
		Config:  cfg,
	}, nil
}

func (b *Benchmark) RunPutIndexLoad(start, finish int) {

	for i := start; i <= finish; i++ {
		err := b.Elastic.PutIndex(start)
		time.Sleep(1 * time.Second)
		if err != nil {
			log.Printf("Stopped at: %d", i)
			break
		}
		log.Printf("sent index: %d", i)
	}

	log.Printf("finished. Started: %d, planned finish: %d", start, finish)
}
