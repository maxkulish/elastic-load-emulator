package main

import (
	"ElasticLoad/bench"
	"ElasticLoad/config"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {

	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fileName := filepath.Base(os.Args[0])
		fmt.Printf("cd to directory from which you are planning to start the program.\n"+
			"Place 'config.yml' and 'index.json' files in this directory.\n"+
			"\nStart load emulation:\n\t./%s -start=1 -stop=100\n"+
			"\nCreate example config and index:\n\t./%s -example\n", fileName, fileName)

		os.Exit(1)
	}

	example := flag.Bool("example", false, "generate example index.json and config.yml")
	start := flag.Int("start", 0, "add start index number")
	stop := flag.Int("stop", 100, "finish index number")

	flag.Parse()

	if *example {
		log.Println("Example config.yml and index.json created")
		bench.CreateExampleIndexFile()
		config.CreateExampleConfig()
		os.Exit(1)
	}

	le, err := bench.NewLoadEmulator()
	if err != nil {
		log.Fatal(err)
	}

	le.RunPutIndexEmulator(*start, *stop)
}
