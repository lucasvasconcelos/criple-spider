package main

import (
	"flag"
	"github.com/lucasvasconcelos/criple-spider/config"
	"github.com/lucasvasconcelos/criple-spider/crawler"
	"log"
	"net/url"
	"os"
)

var AppConfig *config.Spec = config.GetConfig()

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) >= 1 {
		switch args[0] {
		default:
			config.PrintUsage()
			os.Exit(1)
		}

	}

	_, err := url.Parse(AppConfig.StartUrl)

	if err != nil {
		config.PrintUsage()
		log.Fatal(err)
	}

	crawler.StartCrawling(AppConfig.StartUrl)
}
