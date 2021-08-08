package user

type (
	GetSingleUserModel struct {
		Guid         string `form:"guid" gorm:"-" json:"guid,omitempty"`
		UserName     string `form:"userName" gorm:"-" json:"userName,omitempty"`
		WechatOpenId string `form:"wechatOpen_id" gorm:"-" json:"wechatOpen_id,omitempty"`
		Mobile       string `form:"mobile" gorm:"-" json:"mobile,omitempty"`
		Type         byte   `json:"type"`
	}
)
