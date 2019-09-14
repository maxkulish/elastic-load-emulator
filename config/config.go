package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "/config.yml"
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("can't read current directory")
	}

	yamlFile, err := ioutil.ReadFile(dir + configFile)
	if err != nil {
		//CreateExampleConfig()
		return nil, errors.New("can't read config file")
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func CreateExampleConfig() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("can't read current directory")
	}

	// Create file
	file, err := os.Create(dir + configFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := `# Example ElasticSearch config. Put your data
elastic:
  proto: "http"
  host: "1.1.1.1"
  port: "9200"
  username: ""
  password: ""
  index: "index-name"`

	// Write data to file
	_, err = file.Write([]byte(config))
	if err != nil {
		log.Fatal("can't write data to config.yml file", err)
	}
}
