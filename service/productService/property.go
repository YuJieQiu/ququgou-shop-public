package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/product/model"
)

//Property S
//创建属性
func CreateProperty(db *gorm.DB, m *model.Property) (error, *model.Property) {

	err := db.Create(m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//编辑属性
func EditProperty(db *gorm.DB, m *model.Property) (error, *model.Property) {

	var (
		e   model.Property
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
	e.Status = m.Status
	e.Sort = m.Sort

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

//获取属性列表
func GetPropertyList(db *gorm.DB, q *GetPropertyListModel, isPage bool) (err error,
	list []model.Property,
	count int) {

	q.PageSet()

	w := model.Property{
		Name:  q.Name,
		MerId: q.MerId,
	}

	tx := db.Model(&list).Where("mer_id = ? or is_system = ? ", q.MerId, true)

	tx = tx.Where(&w).Order("sort")

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	//if isPage {
	//	err = db.Model(&model.Property{}).Where(&w).Count(&count).Error
	//
	//	if err != nil || count <= 0 {
	//		return err, list, count
	//	}
	//}
	//
	//if isPage {
	//	err = db.Where(&w).Order("sort").Offset(q.Offset).
	//		Limit(q.Limit).Find(&list).Error
	//} else {
	//	err = db.Where(&w).Order("sort").Find(&list).Error
	//}

	if err != nil {
		return err, list, count
	}

	return nil, list, count
}

//Property E

//PropertyValue S
//创建属性
func CreatePropertyValue(db *gorm.DB, m *model.PropertyValue) (error, *model.PropertyValue) {

	err := db.Create(m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//编辑属性
func EditPropertyValue(db *gorm.DB, m *model.PropertyValue) (error, *model.PropertyValue) {

	var (
		e   model.PropertyValue
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
	e.Status = m.Status
	e.Sort = m.Sort

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

//获取属性列表
func GetPropertyValueList(db *gorm.DB, q *GetPropertyListModel, isPage bool) (err error,
	list []model.PropertyValue,
	count int) {

	q.PageSet()

	w := model.PropertyValue{
		Name:       q.Name,
		PropertyId: q.PropertyId,
	}

	if isPage {
		err = db.Model(&model.PropertyValue{}).Where(&w).Count(&count).Error

		if err != nil || count <= 0 {
			return err, list, count
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

	return nil, list, count
}

//PropertyValue E
