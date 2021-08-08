package user

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/modules/user/model"
)

type ModuleUser struct {
	DB *gorm.DB
}

//创建用户
func (m *ModuleUser) CreateUser(user *model.User) (error, *model.User) {

	err := m.DB.Create(user).Error
	if err != nil {
		return err, nil
	}

	return nil, user
}

//获取单个用户信息
func (m *ModuleUser) GetSingleUser(q *GetSingleUserModel) (error, *model.User) {
	var (
		data model.User
		err  error
	)

	w := model.User{
		Guid:     q.Guid,
		UserName: q.UserName,
		//WechatOpenId: q.WechatOpenId,
		Mobile: q.Mobile,
		Type:   q.Type,
	}

	err = m.DB.Where(&w).First(&data).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &data
		}
		return err, nil
	}

	return nil, &data
}

//编辑用户信息
func (m *ModuleUser) EditUser(user *model.User) (error, *model.User) {

	var (
		e   model.User
		err error
	)

	err = m.DB.Where("id = ?", user.ID).First(&e).Error
	if err != nil {
		return err, nil
	}

	err = m.DB.Save(&user).Error

	if err != nil {
		return err, nil
	}

	return nil, user
}

//  UserLoginLog S
func (m *ModuleUser) CreateUserLoginLog(ul *model.UserLoginLog) (error, *model.UserLoginLog) {

	err := m.DB.Create(ul).Error
	if err != nil {
		log.Println(err)
		return err, nil
	}

	return nil, ul
}

//  UserLoginLog E
