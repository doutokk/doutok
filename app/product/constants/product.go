package constants

import "time"

const (
	ProductCategoryKeyPattern = "products:category:%s:page:%d:size:%d"
	ProductKeyPattern         = "search_products:page:%d:pageSize:%d"
	Expire                    = 2 * time.Minute
)
