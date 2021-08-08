package addressService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/user/model"
)

//创建地址信息
func CreateAddress(db *gorm.DB, m *model.Address) (error, *model.Address) {
	var (
		err error
	)

	err = db.Create(&m).Error
	if err != nil {
		return err, nil
	}

	return nil, m
}

//保存
func SaveAddress(db *gorm.DB, m *model.Address) (error, *model.Address) {
	var (
		err  error
		data *model.Address
	)

	if m.ID > 0 {
		err, data = EditAddress(db, m)
		if err != nil {
			return err, nil
		}
	} else {
		err, data = CreateAddress(db, m)
		if err != nil {
			return err, nil
		}
	}

	return nil, data
}

//编辑
func EditAddress(db *gorm.DB, m *model.Address) (error, *model.Address) {

	var (
		e   model.Address
		err error
	)

	err = db.Where("id = ? and user_id=?", m.ID, m.UserId).First(&e).Error
	if err != nil {
		return err, nil
	}

	e.Name = m.Name
	e.Phone = m.Phone
	e.City = m.City
	e.Region = m.Region
	e.Town = m.Town
	e.Address = m.Address
	e.Remark = m.Remark
	e.IsDefault = m.IsDefault

	err = db.Table("addresses").Where("id=?", m.ID).Update(&e).Error
	if err != nil {
		return err, nil
	}

	return nil, &e
}

type DeleteAddressModel struct {
	Id     uint64 `json:"id" form:"id"`
	UserId uint64 `json:"userId" form:"userId"`
}

func DeleteAddress(db *gorm.DB, m *DeleteAddressModel) error {
	var (
		err  error
		data model.Address
	)

	err = db.Where("id=? and  user_id=?", m.Id, m.UserId).First(&data).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if data.ID > 0 {
		err = db.Delete(&data).Error
		if err != nil {
			return err
		}
	}

	return nil
}

type GetAddressModel struct {
	Id     uint64 `json:"id" form:"id"`
	UserId uint64 `json:"userId" form:"userId"`
}

func GetAddress(db *gorm.DB, m *GetAddressModel) (error, *model.Address) {
	var (
		err  error
		data model.Address
	)

	err = db.Where("id=? and  user_id=?", m.Id, m.UserId).First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	return nil, &data
}

//获取用户地址列表
func GetUserAddressList(db *gorm.DB, u *model.User) (error, *[]model.Address) {

	var (
		err  error
		list []model.Address
	)

	err = db.Where("user_id=?", u.ID).Order(" is_default desc").Find(&list).Error
	if err != nil {
		return err, nil
	}

	return nil, &list
}

//获取用户地址列表
func GetUserAddressFirst(db *gorm.DB, u *model.User) (error, *model.Address) {

	var (
		err  error
		data model.Address
	)

	err = db.Where("user_id=?", u.ID).Order(" is_default desc").First(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	return nil, &data
}

//获取全部地址列表
func GetAddressList(db *gorm.DB) (error, *[]model.Address) {
	var (
		err  error
		list []model.Address
	)

	err = db.Order(" id desc").Find(&list).Error
	if err != nil {
		return err, nil
	}

	return nil, &list
}
