package searchService

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/api/userProgram/enum"
	"github.com/ququgou-shop/modules/product/productEnum"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/productService"
)

//获取搜索商品列表
func GetSearchProductList(db *gorm.DB, q *productService.GetSearchProductListModel,
	u *model.User, elasticEnabled bool, imgServiceUrl string) (error, []productService.ProductSmallInfoModel, int) {

	q.PageSet()

	//使用 Elastic

	//参数 Text 可以是商品名 、类别名、属性名...
	//1、需要分词
	//2、需要建立 数据文件 建商品名、类别名、属性名... 放入进去
	//3、建立索引
	//4、 匹配索引
	//q.SearchSortType

	//是否启用
	if elasticEnabled {
		return ElasticsearchProduct(db, q, imgServiceUrl)
	} else {
		//没有启用 使用数据库查询
		return DBSearchProduct(db, q, imgServiceUrl)
	}

}

//数据库 like 查询 产品列表
func DBSearchProduct(db *gorm.DB, q *productService.GetSearchProductListModel, imgServiceUrl string) (error,
	[]productService.ProductSmallInfoModel, int) {
	var (
		err   error
		count int
		list  []productService.ProductSmallInfoModel
		pIds  []uint64
	)

	tx := db.Table("category_products").Select("category_products.product_id").
		Joins("INNER JOIN products on category_products.product_id = products.id").
		Where("category_products.deleted_at IS NULL AND products.deleted_at IS NULL and products.status=" +
			strconv.Itoa(int(productEnum.ProductStatusPutaway)))

	if q.CategoryId > 0 {
		tx = tx.Where("category_products.category_id =? ", q.CategoryId)
	}

	if q.MerId > 0 {
		tx = tx.Where("products.mer_id =? ", q.MerId)
	}

	if q.Text != "" {
		tx = tx.Where("name like ? ", "%"+q.Text+"%")
	}

	switch enum.SearchSortType(q.SearchSortType) {
	case enum.SearchSortTypeDefault:
		tx = tx.Order("products.updated_at desc")
		break
	case enum.SearchSortTypeSalesASC:
		tx = tx.Order("products.sales")
		break
	case enum.SearchSortTypeSalesDESC:
		tx = tx.Order("products.sales desc")
		break
	case enum.SearchSortTypePriceASC:
		tx = tx.Order("products.current_price")
		break
	case enum.SearchSortTypePriceDESC:
		tx = tx.Order("products.current_price desc")
		break
	case enum.SearchSortTypeDIST:
		//TODO:距离排序
		tx = tx.Order("products.updated_at desc")
		break
	default:
		tx = tx.Order("products.updated_at desc")
		break
	}

	tx = tx.Group("category_products.product_id")
	tx = tx.Offset(q.Offset).Limit(q.Limit)

	rows, err := tx.Rows()

	if err != nil {
		return err, nil, count
	}

	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			break
		}
		pIds = append(pIds, id)
	}

	if len(pIds) <= 0 {
		return nil, list, count
	}

	sq := productService.GetProductSmallInfoListModel{
		ProductIds:      pIds,
		Lat:             q.Lat,
		Lon:             q.Lon,
		ComputeDistance: true,
	}
	err, list, _ = productService.GetProductSmallInfoList(db, &sq, false, "", imgServiceUrl)

	if err != nil {
		return err, list, count
	}

	return err, list, count
}

//通过elasticsearch 搜索产品里列表
func ElasticsearchProduct(db *gorm.DB, q *productService.GetSearchProductListModel, imgServiceUrl string) (error, []productService.ProductSmallInfoModel, int) {
	//var (
	//	err   error
	//	count int
	//	list  []product.ProductSmallInfoModel
	//	pIds  []uint64
	//)
	//
	//pq := &elasticsearch.SearchProductModel{
	//	From:           q.Offset,
	//	Size:           q.Limit,
	//	Distance:       q.Distance,
	//	CategoryId:     q.CategoryId,
	//	QueryText:      q.Text,
	//	Lat:            q.Lat,
	//	Lon:            q.Lon,
	//	SearchSortType: q.SearchSortType,
	//}
	//
	//err, data := elasticsearch.SearchProductList(elasticClientBase.ElasticClient, pq)
	//
	//if err != nil {
	//	return err, list, count
	//}
	//
	//if len(data) <= 0 {
	//	return nil, list, count
	//}
	//
	//for _, i := range data {
	//	pIds = append(pIds, i.ID)
	//}
	//
	//sq := shop.GetProductSmallInfoListModel{
	//	ProductIds:      pIds,
	//	Lat:             q.Lat,
	//	Lon:             q.Lon,
	//	ComputeDistance: true,
	//}
	//err, list, count = product.GetProductSmallInfoList(db, &sq, false, "", config.Config.ImgService.QiniuUrl)
	//
	//if err != nil && err != gorm.ErrRecordNotFound {
	//	return err, list, count
	//}

	//return nil, list, count
	return nil, nil, 0
}
