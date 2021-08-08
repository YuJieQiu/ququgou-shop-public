package productEnum

type ProductStatus int

const (
	ProductStatusDefault  ProductStatus = 0  //默认 未上架
	ProductStatusPutaway                = 1  //上架
	ProductStatusUnShelve               = -1 //下架
)

func (s ProductStatus) Text() string {
	switch s {
	case ProductStatusDefault:
		return "未上架"
	case ProductStatusPutaway:
		return "已上架"
	case ProductStatusUnShelve:
		return "已下架"

	default:
		return ""
	}
}
