package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "./config.yml"
)

type Config struct {
	ElasticParams *ElasticParams `yaml:"elastic"`
}

type ElasticParams struct {
	Proto     string `yaml:"proto"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	IndexName string `yaml:"index"`
}

func NewConfig() (*Config, error) {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		createExampleConfig()
		return nil, errors.New("can't read config file")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func createExampleConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("can't create example config file")
	}

	// Create file
	file, err := os.Create(dir + "/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	config := `# Example Elasticsearch config. Put your data
elastic:
  proto: "http"
  host: "1.1.1.1"
  port: "9200"
  username: ""
  password: ""
  index: "index-name"`

	// Write data to file
	_, err = file.Write([]byte(config))
	log.Fatal(err)
}
