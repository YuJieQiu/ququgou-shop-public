package resource

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/resource/model"
)

type ModuleResource struct {
	DB *gorm.DB
}

func (m *ModuleResource) Create(r *model.Resource) error {

	if err := m.DB.Create(r).Error; err != nil {
		return err
	}

	return nil
}

//验证是否存在md5 code
func exist(db *gorm.DB, md5code string) (bool, error, *model.Resource) {

	var (
		m model.Resource
	)

	if err := db.Where("hash_code = ?", md5code).First(&m).Error; err != nil && err != gorm.ErrRecordNotFound {
		return false, err, nil
	}

	res := m.ID > 0

	return res, nil, &m
}

func GetResourceList(db *gorm.DB, q *GetResourceListModel, isPage bool, imageServerUrl string) (err error, list []model.Resource, count int) {

	q.PageSet()

	tx := db.Model(&list)

	if q.MerId > 0 {
		tx = tx.Where(" mer_id = ? ", q.MerId)
	} else if q.UserId > 0 {
		tx = tx.Where(" user_id = ? ", q.UserId)
	}

	tx = tx.Order("updated_at desc")

	if isPage {

		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	for i := 0; i < len(list); i++ {
		list[i].Url = imageServerUrl + list[i].Url
	}

	return nil, list, count
}
