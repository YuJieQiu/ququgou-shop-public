package appConfigService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/appConfig/model"
	"github.com/ququgou-shop/service/productService"
)

//创建首页产品配置信息
func CreateHomeProductConfig(db *gorm.DB, m *model.HomeProductConfig) (error, *model.HomeProductConfig) {
	err := db.Create(&m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//获取首页产品配置列表
func GetHomeProductConfigList(db *gorm.DB, q *GetHomeProductConfigListModel, isPage bool) (err error,
	list []model.HomeProductConfig,
	count int) {

	q.PageSet()

	w := model.HomeProductConfig{
		AppConfigId: q.AppConfigId,
	}

	if isPage {
		err = db.Model(&model.HomeProductConfig{}).Where(&w).Count(&count).Error
		if err != nil {
			return err, list, count
		} else if count <= 0 {
			return nil, list, count
		}
	}

	if isPage {
		err = db.Where(&w).Order("id desc").Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error
	} else {
		err = db.Where(&w).Order("id desc").Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	return nil, list, count
}

//或首页配置和商品基本信息列表
func GetHomeConfigProductInfoList22(db *gorm.DB, q *GetHomeProductConfigListModel, imgServiceUrl string) (error, *[]model.AppConfig) {

	var (
		err  error
		list []model.AppConfig
	)

	c := make(chan *[]productService.ProductSmallInfoModel)
	cErr := make(chan error)
	defer close(c)
	defer close(cErr)

	go GetProductInfoListForConfigId(db, q.AppConfigId, c, cErr, imgServiceUrl)

	for i := 0; i < len(list); i++ {
		list[i].ProductSmallInfos = <-c
		err = <-cErr
		if err != nil {
			return err, nil
		}
	}

	return nil, &list
}

func GetProductInfoListForConfigId(db *gorm.DB,
	configId uint64,
	cList chan *[]productService.ProductSmallInfoModel,
	cErr chan error,
	imgServiceUrl string) {

	var (
		err        error
		productIds []uint64
		list       []productService.ProductSmallInfoModel
	)

	sql := ` select DISTINCT product_id from home_product_configs hpc INNER JOIN app_configs ac on hpc.app_config_id=ac.id
			where hpc.deleted_at is null and ac.deleted_at is null and ac.id=? `

	err = db.Raw(sql, configId).Pluck("product_id", &productIds).Error
	if err != nil {
		cList <- nil
		cErr <- err
		return
	}

	q := productService.GetProductSmallInfoListModel{
		ProductIds: productIds,
	}

	err, list, _ = productService.GetProductSmallInfoList(db, &q, false, "id", imgServiceUrl)

	if err != nil {
		cList <- nil
		cErr <- err
		return
	}

	cList <- &list
	cErr <- nil
	return
}

//获取首页ConfigProductInfoList
func GetHomeConfigProductInfoList(db *gorm.DB, q *GetHomeConfigProductInfoModel, isPage bool,
	imgServiceUrl string) (err error,
	list []productService.ProductSmallInfoModel,
	count int) {
	var (
		pIds []uint64
	)
	q.PageSet()

	tx := db.Table("home_product_configs").Select("home_product_configs.product_id").
		Joins("INNER JOIN products on home_product_configs.product_id = products.id").
		Where("home_product_configs.deleted_at IS NULL AND products.deleted_at IS NULL") //TODO:product 状态筛选

	if q.AppConfigId > 0 {
		tx = tx.Where("home_product_configs.app_config_id =? ", q.AppConfigId)
	} else {
		return err, nil, count
	}
	tx = tx.Order("home_product_configs.updated_at desc").Offset(q.Offset).Limit(q.Limit)
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
		ProductIds: pIds,
	}
	err, list, _ = productService.GetProductSmallInfoList(db, &sq, false, "stock", imgServiceUrl)

	if err != nil {
		return err, list, count
	}

	return err, list, count
}
