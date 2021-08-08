package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//Tags 标签 S
func CreateTags(db *gorm.DB, m *model.Tags) (error, *model.Tags) {

	err := db.Create(m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//编辑标签
func EditTags(db *gorm.DB, m *model.Tags) (error, *model.Tags) {

	var (
		e   model.Tags
		err error
	)

	err = db.Where("id = ?", m.ID).First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}

	e.Name = m.Name

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

//获取标签
func GetTagsList(db *gorm.DB, q *GetTagsListModel, isPage bool) (err error,
	list []model.Tags,
	count int) {

	q.PageSet()

	w := model.Tags{
		MerId: q.MerId,
	}

	tx := db.Model(&list).Where(&w).Order("id desc")

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	return nil, list, count
}

//获取标签
func GetTag(db *gorm.DB, q *GetTagModel) (err error,
	data *model.Tags) {

	w := model.Tags{
		Name:  q.Name,
		MerId: q.MerId,
	}

	err = db.Where(&w).First(&w).Error
	if err != nil {
		return err, nil
	}
	return nil, &w
}

//Tags E
