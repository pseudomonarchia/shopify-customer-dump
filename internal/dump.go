package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	goshopify "github.com/bold-commerce/go-shopify/v4"
)

const (
	limit    = 250
	orderBy  = "created_at asc"
	cacheDir = ".cache"
)

func Dump(shop shopifyConfLi, wg *sync.WaitGroup) {
	defer wg.Done()

	err := initialCacheDir(shop.Name)
	if err != nil {
		log.Fatalf("Error creating cache directory: %v", err)
	}

	ctx := context.TODO()
	store := fmt.Sprintf("%s.myshopify.com", shop.Name)
	client, err := goshopify.NewClient(goshopify.App{}, store, shop.AccessToken)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	totalCount, err := client.Customer.Count(ctx, nil)
	if err != nil {
		log.Fatalf("Error getting customers count: %v", err)
	}

	offset := 0
	currentCount := 0
	pageOpt := &goshopify.ListOptions{Limit: limit, Order: orderBy}

	for {
	work:
		customers, next, err := client.Customer.ListWithPagination(ctx, pageOpt)
		if err != nil {
			log.Printf("Error getting customers: %v", err)
			time.Sleep(1 * time.Second)

			goto work
		}

		customersJson, err := json.Marshal(customers)
		if err != nil {
			log.Fatalf("Error marshalling customers: %v", err)
		}

		offset = currentCount * len(customers)
		currentCount++

		filepath := fmt.Sprintf(
			"%s/%s/%s-%d-%d.json",
			cacheDir,
			shop.Name,
			shop.Name,
			offset,
			offset+len(customers),
		)

		err = write2File(customersJson, filepath)
		if err != nil {
			log.Fatalf("Error writing customers: %v", err)
		}

		log.Printf(
			"Shop: %-17s / Total: %7d, Offset: %7d, Current: %d\n",
			shop.Name,
			totalCount,
			offset,
			currentCount,
		)

		if next.NextPageOptions != nil {
			pageOpt = next.NextPageOptions
		} else {
			log.Printf("Dumped %s", shop.Name)
			break
		}
	}

	filepath := fmt.Sprintf(
		"%s/%s/end.txt",
		cacheDir,
		shop.Name,
	)

	err = write2File(nil, filepath)
	if err != nil {
		log.Fatalf("Error writing end: %v", err)
	}
}

func write2File(data []byte, path string) error {
	return os.WriteFile(path, data, 0644)
}

func initialCacheDir(shopName string) error {
	err := os.MkdirAll(cacheDir, 0755)
	if err != nil {
		return err
	}

	storeDir := fmt.Sprintf("%s/%s", cacheDir, shopName)
	err = os.MkdirAll(storeDir, 0755)
	if err != nil {
		return err
	}

	return nil
}
