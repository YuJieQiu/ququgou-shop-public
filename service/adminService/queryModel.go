package adminService

type (
	GetAdminUserInfoModel struct {
		UserName string `form:"userName" gorm:"-" json:"userName,omitempty"`
		PassWord string `form:"passWord" gorm:"-" json:"passWord,omitempty"`
	}

	GetSingleAdminUserInfoModel struct {
		Id uint64 `form:"id" json:"id"`
	}
)
