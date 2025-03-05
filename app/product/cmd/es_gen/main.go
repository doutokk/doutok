package main

import (
	"context"
	"encoding/json"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/mysql"
	"github.com/doutokk/doutok/app/product/conf"
	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strconv"
	"strings"
)

func main() {
	// Initialize MySQL
	mysql.Init()

	// Fetch products from the database
	var products []model.Product
	if err := mysql.DB.Find(&products).Error; err != nil {
		log.Fatalf("Error fetching products from database: %v", err)
	}

	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.GetConf().ElasticSearch.Address,
		},
	}

	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}

	// Clear the Elasticsearch index
	_, err = esClient.DeleteByQuery(
		[]string{"products"},
		strings.NewReader(`{"query": {"match_all": {}}}`),
		esClient.DeleteByQuery.WithContext(context.Background()),
	)
	if err != nil {
		log.Fatalf("Error clearing Elasticsearch index: %v", err)
	}

	// Insert products into Elasticsearch
	for _, prod := range products {
		var categories []string
		for _, category := range prod.Categories {
			categories = append(categories, category.Name)
		}
		esProduct := &product.Product{
			Id:          uint32(prod.ID),
			Name:        prod.Name,
			Description: prod.Description,
			Picture:     prod.Picture,
			Price:       prod.Price,
			Categories:  categories,
		}

		data, err := json.Marshal(esProduct)
		if err != nil {
			log.Printf("Error marshalling product: %v", err)
			continue
		}

		_, err = esClient.Index(
			"products",                                                // Index name
			strings.NewReader(string(data)),                           // Document body
			esClient.Index.WithDocumentID(strconv.Itoa(int(prod.ID))), // Document ID
			esClient.Index.WithContext(context.Background()),          // Context
		)
		if err != nil {
			log.Printf("Error inserting product into Elasticsearch: %v", err)
			continue
		}

		log.Printf("Inserted product with ID: %s", strconv.Itoa(int(prod.ID)))
	}
}
