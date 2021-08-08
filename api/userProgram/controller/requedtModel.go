package controller

import "github.com/ququgou-shop/library/base_model"

type (
	GetMerProductInfoListModel struct {
		base_model.QueryParamsPage `gorm:"-"`
		MerCode                    string `json:"merCode" form:"merCode"`
		MerId                      uint64 `json:"merId" form:"merId"`
	}
)
