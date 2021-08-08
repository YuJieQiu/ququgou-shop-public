package searchService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/product/productEnum"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/productService"

	"strconv"
)

type QueryRecommendProductListModel struct {
	base_model.QueryParamsPage `gorm:"-"`
	UserId                     uint64  `json:"userId"`
	CategoryId                 uint64  `json:"categoryId" form:"categoryId"` //分类Id
	SearchKey                  string  `json:"searchKey" form:"searchKey"`   //搜索key
	Lat                        float64 `form:"lat" json:"lat"`               //维度
	Lon                        float64 `form:"lon" json:"lon"`               //经度
	Type                       int     `json:"type"`
	Source                     int     `json:"source"` //来源 1、首页 2、购物车 3、个人中心
}

//产品推荐 查询

//TODO:推荐产品列表 Elasticseacrch
//推荐产品列表 数据
func GetRecommendProductListByDB(db *gorm.DB, q *QueryRecommendProductListModel, u *model.User, imgServiceUrl string) (error,
	[]productService.ProductSmallInfoModel, int) {

	var (
		err   error
		count int
		list  []productService.ProductSmallInfoModel
		pIds  []uint64
	)

	q.PageSet()

	//推荐基本逻辑
	//1、根据用户的经纬度 只有通过Elasticseacrch 的有

	//2、根据用户平时浏览的分类类型

	//3、根据商品的推荐权重

	//4、根据价格、销量

	tx := db.Table("products").Select("products.id").
		Where("products.deleted_at IS NULL and products.status=" +
			strconv.Itoa(int(productEnum.ProductStatusPutaway)))

	if q.Source == 1 { //首页
		tx = tx.Order(`products.recommend_priority desc, products.sales desc, products.current_price ,products.updated_at desc`)

	} else if q.Source == 2 { //购物车
		tx = tx.Order(`products.current_price `)

	} else if q.Source == 3 { //个人中心
		tx = tx.Order(`products.sales desc,products.updated_at desc`)

	} else { //否则
		tx = tx.Order(`products.updated_at desc`)
	}

	tx = tx.Offset(q.Offset).Limit(q.Limit)

	rows, err := tx.Rows()

	if err != nil {
		//return err, nil, count
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
		//return nil, list, count
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
