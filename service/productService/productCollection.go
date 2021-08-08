package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//商品收藏 S
//获取用户收藏商品
func GetUserProductCollectionList(db *gorm.DB, q *GetUserProductCollectionListModel,
	imgServiceUrl string) (error, []ProductSmallInfoModel, int) {
	var (
		err   error
		count int
		list  []ProductSmallInfoModel
		pIds  []uint64
	)

	q.PageSet()

	rows, err := db.Model(&model.ProductCollection{}).
		Select("product_collection.product_id").
		Where("user_id=?", q.UserId).Order("product_collection.created_at desc").Offset(q.Offset).Limit(q.Limit).Rows()

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

	sq := GetProductSmallInfoListModel{
		ProductIds:      pIds,
		Lat:             q.Lat,
		Lon:             q.Lon,
		ComputeDistance: false,
	}
	err, list, _ = GetProductSmallInfoList(db, &sq, false, "", imgServiceUrl)

	if err != nil {
		return err, list, count
	}

	return err, list, count
}

//用户添加收藏商品
func UserProductCollectionAdd(db *gorm.DB, q *UserProductCollectionAddModel) error {
	var (
		err error
		m   model.ProductCollection
		p   model.Product
	)

	m.UserId = q.UserId

	err = db.Where("guid=?", q.ProductCode).First(&p).Error
	if err != nil {
		return err
	}
	m.ProductId = p.ID

	err = db.Create(&m).Error

	if err != nil {
		return err
	}
	return nil
}

//用户移除收藏商品
func UserProductCollectionRemove(db *gorm.DB, q *UserProductCollectionRemoveModel) error {
	var (
		err error
		p   model.Product
	)

	err = db.Where("guid=?", q.ProductCode).First(&p).Error
	if err != nil {
		return err
	}

	err = db.Delete(model.ProductCollection{}, "user_id=? and product_id=?", &q.UserId, p.ID).Error

	if err != nil {
		return err
	}
	return nil
}

//获取用户收藏商品数量
func GetUserProductCollectionCount(db *gorm.DB, userId uint64) (error, int) {
	var (
		err   error
		count int
	)

	err = db.Model(&model.ProductCollection{}).Where("user_id = ? and deleted_at is null",
		userId).Count(&count).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, count
	}
	return nil, count
}

//商品收藏 E
