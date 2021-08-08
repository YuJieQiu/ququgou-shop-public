package appConfig

import (
	"github.com/jinzhu/gorm"
)

type ModuleAppConfig struct {
	DB *gorm.DB
}
