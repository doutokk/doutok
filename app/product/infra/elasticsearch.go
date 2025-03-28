package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/doutokk/doutok/app/product/biz/dal/model"
	"github.com/doutokk/doutok/app/product/biz/dal/mysql"
	"github.com/doutokk/doutok/app/product/conf"
	"log"
	"strconv"
	"strings"

	"github.com/doutokk/doutok/rpc_gen/kitex_gen/product"
	"github.com/elastic/go-elasticsearch/v8"
)

var esClient *elasticsearch.Client

// 初始化 Elasticsearch 客户端
func InitElasticsearch() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			conf.GetConf().ElasticSearch.Address,
		},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	esClient = client
}

// 搜索商品
func SearchProducts(ctx context.Context, name string, category string, page int32, pageSize int32) (*product.SearchProductsResp, error) {
	// 构造查询语句
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match": map[string]interface{}{
							"name": map[string]interface{}{
								"query":    name,  // 精确匹配
								"operator": "and", // 精确匹配
							},
						},
					},
					{
						"wildcard": map[string]interface{}{
							"name": map[string]interface{}{
								"value": fmt.Sprintf("%s*", name), // 前缀匹配
							},
						},
					},
					{
						"fuzzy": map[string]interface{}{
							"name": map[string]interface{}{
								"value":     name,   // 模糊匹配
								"fuzziness": "AUTO", // 模糊度
							},
						},
					},
				},
				"minimum_should_match": 1, // 强制返回符合至少一个条件的商品
			},
		},
		"from": (page - 1) * pageSize,
		"size": pageSize,
	}

	// 如果提供了类别，则添加类别筛选
	if category != "" {
		boolQuery := query["query"].(map[string]interface{})["bool"].(map[string]interface{})
		mustQueries := boolQuery["should"].([]map[string]interface{})
		boolQuery["must"] = append(mustQueries, map[string]interface{}{
			"match": map[string]interface{}{
				"categories": map[string]interface{}{
					"query": category,
				},
			},
		})
	}

	// 将查询转换为 JSON
	queryJSON, err := json.Marshal(query)
	if err != nil {
		return nil, fmt.Errorf("error marshalling query: %v", err)
	}

	// 执行 ES 搜索请求
	res, err := esClient.Search(
		esClient.Search.WithContext(ctx),
		esClient.Search.WithIndex("products"),
		esClient.Search.WithBody(strings.NewReader(string(queryJSON))),
		esClient.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, fmt.Errorf("error searching products: %v", err)
	}
	defer res.Body.Close()

	// 解析搜索结果
	var result struct {
		Hits struct {
			Total struct {
				Value int `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source product.Product `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	var total int64
	if err := mysql.DB.Model(&model.Product{}).Count(&total).Error; err != nil {
		return nil, err
	}

	// 构造返回结果
	return &product.SearchProductsResp{
		Item:  extractProducts(result.Hits.Hits),
		Total: int32(total),
	}, nil
}

// 插入商品
func InsertProduct(ctx context.Context, prod *product.Product) error {
	// 将商品数据转换为 JSON
	data, err := json.Marshal(prod)
	if err != nil {
		return fmt.Errorf("error marshalling product: %v", err)
	}

	// 执行 ES 插入请求
	res, err := esClient.Index(
		"products",                                                // 索引名称
		strings.NewReader(string(data)),                           // 商品数据
		esClient.Index.WithContext(ctx),                           // 上下文
		esClient.Index.WithDocumentID(strconv.Itoa(int(prod.Id))), // 文档 ID
	)
	if err != nil {
		return fmt.Errorf("error inserting product: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	log.Printf("Inserted product with ID: %d", prod.Id)
	return nil
}

func DeleteProduct(ctx context.Context, productId uint32) error {
	// 执行 ES 删除请求
	res, err := esClient.Delete(
		"products",                       // 索引名称
		strconv.Itoa(int(productId)),     // 文档 ID
		esClient.Delete.WithContext(ctx), // 上下文
	)
	if err != nil {
		return fmt.Errorf("error deleting product: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	log.Printf("Deleted product with ID: %d", productId)
	return nil
}

// 更新商品
func UpdateProduct(ctx context.Context, prod *product.Product) error {
	// 将商品数据转换为 JSON
	data, err := json.Marshal(prod)
	if err != nil {
		return fmt.Errorf("error marshalling product: %v", err)
	}

	// 执行 ES 更新请求
	res, err := esClient.Index(
		"products",                                                // 索引名称
		strings.NewReader(string(data)),                           // 商品数据
		esClient.Index.WithContext(ctx),                           // 上下文
		esClient.Index.WithDocumentID(strconv.Itoa(int(prod.Id))), // 文档 ID
	)
	if err != nil {
		return fmt.Errorf("error updating product: %v", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error response from Elasticsearch: %s", res.String())
	}

	log.Printf("Updated product with ID: %d", prod.Id)
	return nil
}

// 解析商品数据
func extractProducts(hits []struct {
	Source product.Product `json:"_source"`
}) []*product.Product {
	products := make([]*product.Product, len(hits))
	for i, hit := range hits {
		products[i] = &product.Product{
			Id:          hit.Source.Id,
			Name:        hit.Source.Name,
			Description: hit.Source.Description,
			Picture:     hit.Source.Picture,
			Price:       hit.Source.Price,
			Categories:  hit.Source.Categories,
		}
	}
	return products
}
