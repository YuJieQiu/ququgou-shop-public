package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/model"
	modelResource "github.com/ququgou-shop/modules/resource/model"
)

// Category
//创建分类
func CreateCategory(db *gorm.DB, data *model.Category) (error, *model.Category) {
	var (
		err error
		res *modelResource.Resource
	)

	if data.ResourceId > 0 {
		err, res = GetResource(db, data.ResourceId)
		if err != nil {
			return err, nil
		}
		data.ImageJsonSingleModel.ImageJson = ResourceConverImageJsonSingleModel(res)
	}

	err = db.Create(data).Error
	if err != nil {
		return err, nil
	}

	return nil, data
}

func CategoryChildLoad(db *gorm.DB, list []model.Category) []model.Category {
	var (
		data []model.Category
	)

	for i := 0; i < len(list); i++ {
		if list[i].Pid == 0 {
			d := list[i]
			for _, v := range list {
				if v.Pid == list[i].ID {
					d.Child = append(d.Child, v)
				}
			}
			data = append(data, d)
		}
	}

	return data
}

//获取分类列表
func GetCategoryList(db *gorm.DB, q *GetCategoryListModel, isPage bool, imageServerUrl string) (err error, list []model.Category, count int) {

	q.PageSet()

	//w := model.Category{
	//	Name:   q.Name,
	//	Status: q.Status,
	//	Pid:    q.Pid,
	//	MerId:  q.MerId,
	//}
	//w.ID = q.Id

	var tx *gorm.DB
	if q.IsSystem {
		tx = db.Model(&list).Where("mer_id = ? or is_system = ? ", 0, true)
	} else {
		tx = db.Model(&list).Where("mer_id = ? or is_system = ? ", q.MerId, true)
	}

	if q.Name != "" {
		tx = tx.Where("name =? ", q.Name)
	}

	if q.Id > 0 {
		tx = tx.Where("id =? or pid =?", q.Id, q.Id)
	}

	if q.Pid > 0 {
		tx = tx.Where("  pid =?", q.Pid)
	}

	tx = tx.Order("sort")

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error

	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
	}

	for i := 0; i < len(list); i++ {
		list[i].ImageJson.Url = imageServerUrl + list[i].ImageJson.Url
	}

	return nil, list, count
}

//获取商家商品分类列表
func GetMerchantCategoryList(db *gorm.DB, q *GetMerchantCategoryListModel, isPage bool, imageServerUrl string) (err error, list []model.Category, count int) {

	q.PageSet()

	var tx *gorm.DB

	tx = db.Model(&list).Where("mer_id = ? or is_system = ? ", q.MerId, false)

	if q.Name != "" {
		tx = tx.Where("name =? ", q.Name)
	}

	if q.Id > 0 {
		tx = tx.Where("id =? or pid =?", q.Id, q.Id)
	}

	if q.Pid > 0 {
		tx = tx.Where("  pid =?", q.Pid)
	}

	tx = tx.Order("sort")

	if isPage {
		err = tx.Count(&count).Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error
	} else {
		err = tx.Find(&list).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
	}

	for i := 0; i < len(list); i++ {
		if len(list[i].ImageJson.Url) > 0 {
			list[i].ImageJson.Url = imageServerUrl + list[i].ImageJson.Url
		}
	}

	return nil, list, count
}

type MerchantCategoryListSaveModel struct {
	MerId        uint64           `json:"merId"`
	CategoryList []model.Category `json:"categoryList"`
}

//商家商品分类保存
func MerchantCategoryListSave(db *gorm.DB, q *MerchantCategoryListSaveModel) error {

	if q.MerId <= 0 {
		return nil
	}

	var (
		ids           []uint64
		category      *model.Category
		childCategory *model.Category
		err           error
	)

	tx := db.Begin()

	for _, v := range q.CategoryList {

		v.IsSystem = false
		v.MerId = q.MerId

		if v.ID > 0 {
			ids = append(ids, v.ID)
			err, category = EditCategory(tx, &v)
		} else {
			err, category = CreateCategory(tx, &v)
			ids = append(ids, category.ID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		if v.Child != nil && len(v.Child) > 0 {
			for _, c := range v.Child {
				c.IsSystem = false
				c.MerId = q.MerId

				if c.ID > 0 {
					ids = append(ids, c.ID)
					c.Pid = v.ID
					err, childCategory = EditCategory(tx, &c)
				} else {
					c.Pid = category.ID
					err, childCategory = CreateCategory(tx, &c)
					ids = append(ids, childCategory.ID)
				}
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	if len(ids) > 0 {
		err = tx.Where("id not in (?) and mer_id = ? and is_system = ? ", ids, q.MerId, false).Delete(&model.Category{}).Error
	} else {
		err = tx.Where(" mer_id = ? and is_system = ?", q.MerId, false).Delete(&model.Category{}).Error
	}

	tx.Commit()
	return nil
}

type SystemCategoryListSaveModel struct {
	CategoryList []model.Category `json:"categoryList"`
}

//系统分类修改保存
func SystemCategoryListSave(db *gorm.DB, q *SystemCategoryListSaveModel) error {
	var (
		ids           []uint64
		category      *model.Category
		childCategory *model.Category
		err           error
	)

	tx := db.Begin()

	for _, v := range q.CategoryList {

		v.IsSystem = true

		if v.ID > 0 {
			ids = append(ids, v.ID)
			err, category = EditCategory(tx, &v)
		} else {
			err, category = CreateCategory(tx, &v)
			ids = append(ids, category.ID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		if v.Child != nil && len(v.Child) > 0 {
			for _, c := range v.Child {

				c.IsSystem = true

				if c.ID > 0 {
					ids = append(ids, c.ID)
					c.Pid = v.ID
					err, childCategory = EditCategory(tx, &c)
				} else {
					c.Pid = category.ID
					err, childCategory = CreateCategory(tx, &c)
					ids = append(ids, childCategory.ID)
				}
				if err != nil {
					tx.Rollback()
					return err
				}
			}
		}
	}

	if len(ids) > 0 {
		err = tx.Where("id not in (?) and mer_id = ? and is_system = ? ", ids, 0, true).Delete(&model.Category{}).Error
	} else {
		err = tx.Where(" mer_id = ? and is_system = ?", 0, true).Delete(&model.Category{}).Error
	}

	tx.Commit()
	return nil
}

//编辑分类
func EditCategory(tx *gorm.DB, m *model.Category) (error, *model.Category) {
	var (
		e model.Category

		res *modelResource.Resource
		err error
	)

	if m.IsSystem {
		err = tx.Where("id = ?", m.ID).First(&e).Error
	} else {
		err = tx.Where("id = ? and mer_id=?", m.ID, m.MerId).First(&e).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}

	e.Name = m.Name
	e.Status = m.Status
	e.Sort = m.Sort
	e.Pid = m.Pid
	e.Remark = m.Remark
	e.MerId = m.MerId
	e.IsSystem = m.IsSystem

	if m.ResourceId == 0 && e.ResourceId != 0 {
		e.ImageJsonSingleModel.ImageJson = ext_struct.JsonImageString{}
	} else if m.ResourceId != 0 && e.ResourceId != m.ResourceId {
		err, res = GetResource(tx, m.ResourceId)
		if err != nil {
			return err, nil
		}
		e.ImageJsonSingleModel.ImageJson = ResourceConverImageJsonSingleModel(res)
	}

	e.ResourceId = m.ResourceId

	//.Where("id = ?", m.ID).Update(&e)

	err = tx.Save(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

// Category E
