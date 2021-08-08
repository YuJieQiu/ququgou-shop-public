package adminService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/base_model"
	"github.com/ququgou-shop/modules/adminuser/model"
)

//创建管理员用户
//TODO:用户名密码验证 、 密码加密等处理
func CreateUser(db *gorm.DB, user *model.AdminUser) (error, *model.AdminUser) {

	err := db.Create(user).Error
	if err != nil {
		return err, nil
	}

	return nil, user
}

//获取管理员用户信息 //用户名密码登录使用
func GetAdminUserInfo(db *gorm.DB, q *GetAdminUserInfoModel) (error, *model.AdminUser) {
	var (
		err error
		u   model.AdminUser
	)
	param := &model.AdminUser{
		UserName: q.UserName,
	}

	err = db.Where(param).First(&u).Error

	if err != nil {
		return err, nil
	}

	if u.PassWord != q.PassWord {
		return ErrPassWord, nil
	}

	return nil, &u
}

//获取管理员用户信息
func GetSingleAdminUserInfo(db *gorm.DB, id uint64) (error, *model.AdminUser) {
	var (
		err error
		u   model.AdminUser
	)

	param := &model.AdminUser{
		IDAutoModel: base_model.IDAutoModel{id},
	}

	err = db.Where(param).First(&u).Error
	if err != nil {
		return err, nil
	}

	return nil, &u
}
