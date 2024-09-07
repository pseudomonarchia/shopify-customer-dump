package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/pseudomonarchia/shopify-customer-dump/internal"
)

const (
	confFileName = "config.shopify.yaml"
	logFileName  = "logs.txt"
)

func main() {
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	conf, err := internal.ReadConf(confFileName)
	if err != nil {
		fmt.Println(err)
	}

	wg := sync.WaitGroup{}
	for _, shop := range conf.Shopify {
		wg.Add(1)
		go internal.Dump(shop, &wg)
	}

	wg.Wait()
}
