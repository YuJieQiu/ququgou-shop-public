package merchantService

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	merEmun "github.com/ququgou-shop/modules/merchant/merEnum"
	"github.com/ququgou-shop/modules/merchant/model"
	userModel "github.com/ququgou-shop/modules/user/model"
)

//创建申请信息
func CreateMerApplyInfo(db *gorm.DB, q *CreateMerApplyInfoModel, u *userModel.User) error {
	var (
		err   error
		apply model.MerchantApply
	)

	err = db.Where("user_id = ?", u.ID).First(&apply).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	if apply.ID > 0 {
		return ErrApplyRepeat
	}

	var resourcesIds []uint64
	//resourcesIdsStr := ""
	if len(q.Resources) > 0 {
		for _, i := range q.Resources {
			resourcesIds = append(resourcesIds, i.Id)
		}
	}
	resourcesIdsStr, _ := json.Marshal(resourcesIds)

	apply = model.MerchantApply{
		UserId:       u.ID,
		Name:         q.Name,
		Description:  q.Description,
		TypeId:       0,
		City:         q.City,
		Region:       q.Region,
		Town:         q.Town,
		Address:      q.Address,
		Longitude:    q.Longitude,
		Latitude:     q.Latitude,
		Status:       0,
		Remark:       q.Remark,
		ResourcesIds: string(resourcesIdsStr),
		Phone:        q.Phone,
	}

	err = db.Create(&apply).Error

	return err
}

//获取申请信息
func GetMerApplyInfo(db *gorm.DB, u *userModel.User, imageServerUrl string) (error, *MerApplyInfoModel) {
	var (
		err       error
		data      MerApplyInfoModel
		resources []MerResourceModel
	)

	err = db.Table("merchant_apply").Where(" user_id = ? ", u.ID).First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}
	data.StatusText = merEmun.MerApplyStatus(data.Status).Text()

	if len(data.ResourcesIds) > 0 {
		var resourcesIds []uint64
		bytes := []byte(data.ResourcesIds)
		_ = json.Unmarshal(bytes, &resourcesIds)

		err := db.Table("resources").Where("id in (?)", resourcesIds).Find(&resources).Error
		if err != nil {

		} else {
			for i := 0; i < len(resources); i++ {
				resources[i].Url = imageServerUrl + resources[i].Url
			}
			data.Resources = resources
		}
	}

	return nil, &data
}

//通过审核
func MerApplyVerified(db *gorm.DB, id uint64) error {
	var (
		err       error
		applyInfo model.MerchantApply
		mer       model.Merchant
		ma        model.MerchantAddress
		mc        model.MerchantConfig
		md        model.MerchantDetail
		mu        model.MerchantUser
		//mr        []model.MerchantResources
	)

	err = db.Where("id=?", id).First(&applyInfo).Error
	if err != nil {
		return err
	}

	if merEmun.MerApplyStatus(applyInfo.Status) != merEmun.MerApplyStatusReview {
		return err
	}

	mer.Guid = utils.CreateUUID()
	mer.Name = applyInfo.Name
	mer.Description = applyInfo.Description
	mer.Status = 0
	mer.TypeId = applyInfo.TypeId
	mer.Active = true //TODO: 后面让用户选择是否激活

	ma.Remark = applyInfo.Remark
	ma.Latitude = applyInfo.Latitude
	ma.Longitude = applyInfo.Longitude
	ma.City = applyInfo.City
	ma.Address = applyInfo.Address
	ma.Town = applyInfo.Town
	ma.Active = true

	md.Phones = applyInfo.Phone

	mu = model.MerchantUser{
		UserId:  applyInfo.UserId,
		IsAdmin: true,
		Status:  0,
		Active:  true,
	}

	tx := db.Begin()

	err = tx.Create(&mer).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	ma.MerchantId = mer.ID

	mc.MerchantId = mer.ID

	md.MerchantId = mer.ID

	mu.MerchantId = mer.ID

	tx.Create(&ma)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Create(&mc)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Create(&md)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Create(&mu)
	if err != nil {
		tx.Rollback()
		return err
	}

	if len(applyInfo.ResourcesIds) > 0 {
		var resourcesIds []uint64
		bytes := []byte(applyInfo.ResourcesIds)
		_ = json.Unmarshal(bytes, &resourcesIds)

		for index, i := range resourcesIds {
			mr := model.MerchantResources{
				ResourcesId: i,
				Sort:        index,
				MerchantId:  mer.ID,
			}
			tx.Create(&mr)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	applyInfo.Status = int(merEmun.MerApplyStatusVerified)
	tx.Save(&applyInfo)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
