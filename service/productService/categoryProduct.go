package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//获取分类商品列表
func GetCategoryProductList(db *gorm.DB, q *GetCategoryProductListModel, imgServiceUrl string) (error,
	[]ProductSmallInfoModel, int) {
	var (
		err   error
		count int
		cps   []model.CategoryProduct
		list  []ProductSmallInfoModel
		pIds  []uint64
	)

	q.PageSet()

	tx := db.Model(&cps).Where("category_id = ? ", q.CategoryId)

	err = tx.Count(&count).Offset(q.Offset).
		Limit(q.Limit).Find(&cps).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
	}

	if len(cps) == 0 {
		return err, list, count
	}

	for _, v := range cps {
		pIds = append(pIds, v.ProductId)
	}

	sq := GetProductSmallInfoListModel{}
	err, list, _ = GetProductSmallInfoList(db, &sq, false, "stock", imgServiceUrl)

	if err != nil {
		return err, list, count
	}

	return err, list, count
}
