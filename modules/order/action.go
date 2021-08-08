package order

import (
	"github.com/jinzhu/gorm"
)

type ModuleOrder struct {
	DB            *gorm.DB
	ImgServiceUrl string
}
