package main

import (
	"ElasticLoad/bench"
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
			"How to start load emulation:\n\t./%s -start=1 -stop=100\n", fileName)

		os.Exit(1)
	}

	start := flag.Int("start", 0, "add start index number")
	finish := flag.Int("finish", 100, "finish index number")

	flag.Parse()

	le, err := bench.NewLoadEmulator()
	if err != nil {
		log.Fatal(err)
	}

	le.RunPutIndexEmulator(*start, *finish)
}
