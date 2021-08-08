package merchantService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/label/model"
)

//创建标签
func CreateLabel(db *gorm.DB, m *model.Label) (error, *model.Label) {
	var (
		err error
	)

	err = db.Create(m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//获取标签列表
func GetLabelList(db *gorm.DB, q *GetLabelListModel, isPage bool) (err error,
	list []model.Label,
	count int) {

	q.PageSet()

	w := model.Label{
		Type: q.Type,
	}

	tx := db.Model(&list).Where("type = ? or is_system = ? ", q.Type, true)

	tx = tx.Where(&w).Order("id")

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
func GetLabel(db *gorm.DB, q *GetLabelModel) (err error,
	data *model.Label) {

	w := model.Label{
		Text: q.Text,
		Type: q.Type,
	}

	err = db.Where("name=? and (type = ? or is_system = ? )",
		q.Text, q.Type, true).First(&w).Error
	if err != nil {
		return err, nil
	}
	return nil, &w
}

//编辑标签
func EditLabel(db *gorm.DB, m *model.Label) (error, *model.Label) {

	var (
		e   model.Label
		err error
	)

	err = db.Where("id = ?", m.ID).First(&e).Error
	if err != nil {
		return err, nil
	}

	e.Text = m.Text

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}
