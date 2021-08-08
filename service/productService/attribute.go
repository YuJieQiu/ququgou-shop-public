package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/product/model"
	modelResource "github.com/ququgou-shop/modules/resource/model"
)

//Attribute
//创建规格
func CreateAttribute(db *gorm.DB, q *model.Attribute) (error, *model.Attribute) {

	var (
		err error
		res *modelResource.Resource
	)

	if q.ResourceId > 0 {
		err, res = GetResource(db, q.ResourceId)
		if err != nil {
			return err, nil
		}
		q.ImageJsonSingleModel.ImageJson = ResourceConverImageJsonSingleModel(res)
	}

	err = db.Create(q).Error
	if err != nil {
		return err, nil
	}

	return nil, q
}

//获取规格
func GetAttribute(db *gorm.DB, q *GetAttributeModel) (err error,
	data *model.Attribute) {

	w := model.Attribute{
		Name:  q.Name,
		MerId: q.MerId,
	}

	err = db.Where("name=? and (mer_id = ? or is_system = ? )",
		q.Name, q.MerId, true).First(&w).Error
	if err != nil {
		return err, nil
	}
	return nil, &w
}

//编辑规格
func EditAttribute(db *gorm.DB, q *model.Attribute) (error, *model.Attribute) {

	var (
		e   model.Attribute
		res *modelResource.Resource
		err error
	)

	err = db.Where("id = ?", q.ID).First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}

	e.Name = q.Name
	e.Status = q.Status
	e.Sort = q.Sort
	e.CategoryId = q.CategoryId
	e.Remark = q.Remark

	if q.ResourceId == 0 && e.ResourceId != 0 {
		e.ImageJson = ext_struct.JsonImageString{}
	} else if q.ResourceId != 0 && e.ResourceId != q.ResourceId {
		err, res = GetResource(db, q.ResourceId)
		if err != nil {
			return err, nil
		}
		e.ImageJsonSingleModel.ImageJson = ResourceConverImageJsonSingleModel(res)
	}

	e.ResourceId = q.ResourceId

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

//获取规格列表 isValues: 是否获取value
func GetAttributeList(db *gorm.DB, q *GetAttributeListModel, isPage bool, isValues bool) (err error, list []model.Attribute, count int) {

	q.PageSet()

	w := model.Attribute{
		Name:       q.Name,
		CategoryId: q.CategoryId,
	}

	tx := db.Model(&list).Where("mer_id = ? or is_system = ? ", q.MerId, true)

	tx = tx.Where(&w).Order("sort")

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	if isValues {
		var ids []uint64
		var values []model.AttributeValue
		for _, v := range list {
			ids = append(ids, v.ID)
		}
		ids = utils.RemoveRepeatedElemenForInt(ids)
		db.Where("attribute_id in (?)", ids).Find(&values)
		for i := 0; i < len(list); i++ {
			for v := 0; v < len(values); v++ {
				if list[i].ID == values[v].AttributeId {
					list[i].Values = append(list[i].Values, values[v])
					continue
				}
			}
		}
	}

	return nil, list, count
}

//Attribute E

//AttributeValue S
//创建规格属性
func CreateAttributeValue(db *gorm.DB, q *model.AttributeValue) (error, *model.AttributeValue) {
	return createAttributeValue(db, q)
}

func createAttributeValue(db *gorm.DB, q *model.AttributeValue) (error, *model.AttributeValue) {
	err := db.Create(q).Error
	if err != nil {
		return err, nil
	}

	return nil, q
}

//获取规格
func GetAttributeValue(db *gorm.DB, q *GetAttributeValueModel) (err error,
	data *model.AttributeValue) {

	w := model.AttributeValue{
		Name:        q.Name,
		MerId:       q.MerId,
		AttributeId: q.AttributeId,
	}

	err = db.Where("name=? and attribute_id=? and (mer_id = ? or is_system = ? )",
		q.Name, q.AttributeId, q.MerId, true).First(&w).Error
	if err != nil {
		return err, nil
	}
	return nil, &w
}

//编辑属性
func EditAttributeValue(db *gorm.DB, q *model.AttributeValue) (error, *model.AttributeValue) {

	var (
		e   model.AttributeValue
		res *modelResource.Resource
		err error
	)

	err = db.Where("id = ?", q.ID).First(&e).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}

	e.Name = q.Name
	e.Status = q.Status
	e.Sort = q.Sort
	e.AttributeId = q.AttributeId
	e.Remark = q.Remark

	if q.ResourceId == 0 && e.ResourceId != 0 {
		e.ImageJson = ext_struct.JsonImageString{}
	} else if q.ResourceId != 0 && e.ResourceId != q.ResourceId {
		err, res = GetResource(db, q.ResourceId)
		if err != nil {
			return err, nil
		}
		e.ImageJsonSingleModel.ImageJson = ResourceConverImageJsonSingleModel(res)
	}

	e.ResourceId = q.ResourceId

	err = db.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

//获取属性列表
func GetAttributeValueList(db *gorm.DB, q *GetAttributeValueListModel, isPage bool) (err error,
	list []model.AttributeValue,
	count int) {

	q.PageSet()

	w := model.AttributeValue{
		Name:        q.Name,
		AttributeId: q.AttributeId,
		MerId:       q.MerId,
	}
	tx := db.Model(&list).Where(&w).Order("sort")

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

//AttributeValue E
