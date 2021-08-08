package resource

type (
	//分页
	QueryParamsPage struct {
		Page   int `form:"page" json:"page,omitempty"`
		Limit  int `form:"limit" json:"limit,omitempty"`
		Offset int `form:"offset" json:"offset,omitempty"`
	}

	GetResourceListModel struct {
		QueryParamsPage `gorm:"-"`
		MerId           uint64 `form:"merId" json:"merId" `
		UserId          uint64 `form:"userId" json:"userId"`
	}
)

func (page *QueryParamsPage) PageSet() {
	if page.Page <= 0 {
		page.Page = 1
	}

	if page.Limit <= 0 {
		page.Limit = 20
	}

	page.Offset = (page.Page - 1) * page.Limit
}
