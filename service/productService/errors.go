package productService

import "errors"

var (
	//
	ErrUserNameExist = errors.New("user name exist")

	ErrProductSKUEmpty = errors.New("product sku empty")

	ErrProductPutaway = errors.New("商品已下架")
)
