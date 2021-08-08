package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/resource/model"
)

//
func ResourceConverImageJsonSingleModel(resource *model.Resource) ext_struct.JsonImageString {

	j := ext_struct.JsonImageString{
		Guid: resource.Guid,
		Path: resource.Path,
		Url:  resource.Url,
		Type: resource.Type,
	}

	return j
}

func GetResource(db *gorm.DB, resourceId uint64) (error, *model.Resource) {
	var (
		err error
		res model.Resource
	)

	if err = db.First(&res, resourceId).Error; err != nil {
		return err, nil
	}

	return nil, &res
}
