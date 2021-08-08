package appConfigService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/appConfig/model"
	rModel "github.com/ququgou-shop/modules/resource/model"
)

func CreateAppConfig(db *gorm.DB, appConf *model.AppConfig) (error, *model.AppConfig) {
	var (
		r rModel.Resource
	)

	if appConf.ResourceId > 0 {
		db.Where("id=?", appConf.ResourceId).First(&r)
	}

	if r.ID > 0 {
		j := ext_struct.JsonImageString{
			Guid: r.Guid,
			Path: r.Path,
			Url:  r.Url,
		}
		appConf.ImageJson = j
	}

	err := db.Create(appConf).Error
	if err != nil {
		return err, nil
	}
	return nil, appConf
}

func GetAppConfigList(db *gorm.DB, q *GetAppConfigListModel, isPage bool, imgServiceUrl string) (err error, list []model.AppConfig, count int) {

	q.PageSet()

	w := model.AppConfig{
		Name: q.Name,
		Type: q.Type,
	}

	if isPage {
		err = db.Model(&model.AppConfig{}).Where(&w).Count(&count).Error
		if err != nil {
			return err, list, count
		} else if count <= 0 {
			return nil, list, count
		}
	}

	if isPage {
		err = db.Where(&w).Order("sort").Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error
	} else {
		err = db.Where(&w).Order("sort").Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	for i := 0; i < len(list); i++ {
		list[i].ImageJson.Url = imgServiceUrl + list[i].ImageJson.Url
	}

	return nil, list, count
}

func SaveAppConfig(db *gorm.DB, a *AppConfigSaveModel) error {
	var (
		ids []uint64
		err error
	)

	for index, i := range a.AppConfigs {
		if a.AppConfigs[index].Type == "" || a.AppConfigs[index].Type != a.Type {
			a.AppConfigs[index].Type = a.Type
		}

		if i.ID > 0 {
			ids = append(ids, i.ID)
		}
	}

	tx := db.Begin()

	if len(ids) > 0 {
		err = tx.Where("id not in (?) and type=?", ids, a.Type).Delete(&model.AppConfig{}).Error
	} else {
		err = tx.Where("type=?", a.Type).Delete(&model.AppConfig{}).Error
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, i := range a.AppConfigs {
		if i.ID > 0 {
			//i.Sort = index
			err, _ = EditAppConfig(tx, &i)
		} else {
			//i.Sort = index
			err, _ = CreateAppConfig(tx, &i)
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func EditAppConfig(db *gorm.DB, a *model.AppConfig) (error, *model.AppConfig) {

	var (
		e   model.AppConfig
		r   rModel.Resource
		err error
	)

	err = db.Where("id = ?", a.ID).First(&e).Error
	if err != nil {
		return err, nil
	}

	e.Name = a.Name
	e.Type = a.Type
	e.Description = a.Description
	e.LinkUrl = a.LinkUrl
	e.Sort = a.Sort
	e.ResourceId = a.ResourceId
	e.Status = a.Status
	e.Code = a.Code
	e.Text = a.Text
	e.LinkType = a.LinkType

	if e.ResourceId != a.ResourceId {

		e.ResourceId = a.ResourceId

		db.Where("id=?", a.ResourceId).First(&r)

		if r.ID > 0 {
			j := ext_struct.JsonImageString{
				Guid: r.Guid,
				Path: r.Path,
				Url:  r.Url,
			}
			e.ImageJson = j
		}

	}

	err = db.Table("app_configs").Where("id=?", a.ID).Update(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

func GetAppConfigMapList(db *gorm.DB, q *GetAppConfigListModel, isPage bool,
	imageServerUrl string) (err error, data map[string][]model.AppConfig, count int) {

	data = make(map[string][]model.AppConfig)

	err, list, count := GetAppConfigList(db, q, isPage, imageServerUrl)

	if err != nil {
		return err, data, count
	}

	for i := 0; i < len(list); i++ {
		if _, ok := data[list[i].Type]; ok {
			data[list[i].Type] = append(data[list[i].Type], list[i])
		} else {
			data[list[i].Type] = append(data[list[i].Type], list[i])
		}
	}

	return nil, data, count
}

func GetAppConfigTypeList(db *gorm.DB) (error, *[]model.AppConfigType) {
	var (
		err  error
		list []model.AppConfigType
	)
	err = db.Where("status = ?", 1).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}
	return nil, &list
}
