package appConfigService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/banner/model"
	rModel "github.com/ququgou-shop/modules/resource/model"
)

//创建 banner
func CreateBanner(db *gorm.DB, banner *model.Banner) (error, *model.Banner) {
	var (
		r rModel.Resource
	)

	if banner.ResourceId > 0 {
		db.Where("id=?", banner.ResourceId).First(&r)

		if r.ID > 0 {
			j := ext_struct.JsonImageString{
				Guid: r.Guid,
				Path: r.Path,
				Url:  r.Url,
			}
			banner.ImageJson = j
		}
	}

	err := db.Create(banner).Error
	if err != nil {
		return err, nil
	}

	return nil, banner
}

//获取banner列表
func GetBannerList(db *gorm.DB, q *GetBannerListModel, isPage bool, imgServiceUrl string) (err error, list []model.Banner, count int) {

	q.QueryParamsPage.PageSet()

	w := model.Banner{
		Name: q.Name,
	}

	if isPage {
		err = db.Where(&w).Order("sort").Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error
	} else {
		err = db.Where(&w).Order("sort").Find(&list).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, list, count
	}

	for i := 0; i < len(list); i++ {
		list[i].ImageJson.Url = imgServiceUrl + list[i].ImageJson.Url
	}

	return nil, list, count
}

func SaveBanner(db *gorm.DB, banner *BannerSaveModel) error {
	var (
		ids []uint64
		err error
	)

	for _, i := range banner.Banners {
		if i.ID > 0 {
			ids = append(ids, i.ID)
		}
	}

	tx := db.Begin()

	if len(ids) > 0 {
		err = tx.Where("id not in (?)", ids).Delete(&model.Banner{}).Error
	} else {
		err = tx.Where("").Delete(&model.Banner{}).Error
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	for _, i := range banner.Banners {
		if i.ID > 0 {
			//i.Sort = index
			err, _ = EditBanner(tx, &i)
		} else {
			//i.Sort = index
			err, _ = CreateBanner(tx, &i)
		}
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

//编辑banner

func EditBanner(db *gorm.DB, banner *model.Banner) (error, *model.Banner) {

	var (
		e   model.Banner
		r   rModel.Resource
		err error
	)

	err = db.Where("id = ?", banner.ID).First(&e).Error
	if err != nil {
		return err, nil
	}

	e.Name = banner.Name
	e.Type = banner.Type
	e.Description = banner.Description
	e.LinkUrl = banner.LinkUrl
	e.Position = banner.Position
	e.Sort = banner.Sort
	e.BackgroundColor = banner.BackgroundColor
	e.FontColor = banner.FontColor

	if e.ResourceId != banner.ResourceId {

		e.ResourceId = banner.ResourceId

		db.Where("id=?", banner.ResourceId).First(&r)

		if r.ID > 0 {
			j := ext_struct.JsonImageString{
				Guid: r.Guid,
				Path: r.Path,
				Url:  r.Url,
			}
			e.ImageJson = j
		}
	}

	err = db.Table("banners").Where("id=?", banner.ID).Update(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}
