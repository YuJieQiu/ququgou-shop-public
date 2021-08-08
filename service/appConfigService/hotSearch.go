package appConfigService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/appConfig"
	"github.com/ququgou-shop/modules/appConfig/model"
)

var appConfigModule = appConfig.ModuleAppConfig{}

//获取首页热搜词
func GetHomeHotSearchList(db *gorm.DB) (error, []model.HotSearchConfig) {
	var hotSearchList []model.HotSearchConfig
	appConfigModule.DB = db

	//TODO 获取缓存
	//if bol := cache.WebCache.GetData(cache.HomeHotSearch, &hotSearchList); bol {
	//	return nil, hotSearchList
	//}

	//获取首页热搜配置
	hotSearchReq := GetHotSearchConfigListModel{IsHome: true}
	hotSearchReq.Page = 1
	hotSearchReq.Limit = 5

	err, hotSearchList, _ := GetHotSearchConfigList(db, &hotSearchReq, true)
	if err != nil {
		return err, hotSearchList
	}

	//一分钟
	//var outTime time.Duration = 60 * time.Second //一分钟

	//TODO 保存缓存
	//go cache.WebCache.SetData(cache.HomeHotSearch, hotSearchList, outTime)

	return nil, hotSearchList
}

//获取热搜词
func GetHotSearchList(db *gorm.DB) (error, []model.HotSearchConfig) {

	var hotSearchList []model.HotSearchConfig
	appConfigModule.DB = db

	//TODO 获取缓存
	//if bol := cache.WebCache.GetData(cache.HotSearch, &hotSearchList); bol {
	//	return nil, hotSearchList
	//}

	//获取首页热搜配置
	hotSearchReq := GetHotSearchConfigListModel{IsHome: false}
	hotSearchReq.Page = 1
	hotSearchReq.Limit = 20

	err, hotSearchList, _ := GetHotSearchConfigList(db, &hotSearchReq, true)
	if err != nil {
		return err, hotSearchList
	}

	//一分钟
	//var outTime time.Duration = 60 * time.Second //一分钟

	//TODO 保存缓存
	//go cache.WebCache.SetData(cache.HotSearch, hotSearchList, outTime)

	return nil, hotSearchList
}

//热搜索配置
func HotSearchConfigListSave(db *gorm.DB, data *HotSearchConfigListSaveModel) error {
	var (
		err             error
		ids             []uint64
		hotSearchConfig *model.HotSearchConfig
	)

	tx := db.Begin()

	for _, v := range data.List {

		if v.ID > 0 {
			ids = append(ids, v.ID)
			err, _ = EditHotSearchConfig(tx, &v, data.UserId)
		} else {
			err, hotSearchConfig = createHotSearchConfig(tx, &v)
			ids = append(ids, hotSearchConfig.ID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if len(ids) > 0 {
		err = tx.Where("id not in (?) ", ids).Delete(&model.HotSearchConfig{}).Error
	} else {
		err = tx.Where(" deleted_at is null").Delete(&model.HotSearchConfig{}).Error
	}

	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

func EditHotSearchConfig(tx *gorm.DB, data *model.HotSearchConfig, userId uint64) (error,
	*model.HotSearchConfig) {

	var (
		e   model.HotSearchConfig
		err error
	)

	err = tx.Where("id = ?", data.ID).First(&e).Error

	if err != nil {
		return err, nil
	}

	e.Text = data.Text
	e.IsHome = data.IsHome
	e.OpenUserId = userId
	e.Sort = data.Sort

	err = tx.Table("hot_search_config").Where("id = ?", data.ID).Update(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

func createHotSearchConfig(db *gorm.DB, m *model.HotSearchConfig) (error, *model.HotSearchConfig) {
	var (
		err error
	)

	err = db.Create(m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

func GetHotSearchConfigList(db *gorm.DB, q *GetHotSearchConfigListModel, isPage bool) (error,
	[]model.HotSearchConfig, int) {
	var (
		err   error
		list  []model.HotSearchConfig
		count int
	)
	q.PageSet()

	var tx *gorm.DB

	if q.IsHome {
		tx = db.Model(&list).Where("is_home = ?", true).Order("sort")
	} else {
		tx = db.Model(&list).Order("sort")
	}

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil, count
		}
		return err, nil, count
	}

	return nil, list, count
}
