package merchantService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/utils"
	merEmun "github.com/ququgou-shop/modules/merchant/merEnum"
	"github.com/ququgou-shop/modules/merchant/model"
)

//获取用户商户信息
func GetMerchantForUser(db *gorm.DB, userId uint64) (error, *GetMerchantForUserModel) {
	var (
		err     error
		meruser model.MerchantUser
		mer     model.Merchant
		data    GetMerchantForUserModel
	)

	err = db.Where("status=? and active =? and user_id=?", int(merEmun.MerUserStatusPass), true, userId).
		First(&meruser).Error

	if err != nil {
		return err, nil
	}

	err = db.Where("id=? and status=? and active =?",
		meruser.MerchantId,
		int(merEmun.MerUserStatusPass),
		true).
		First(&mer).Error

	if err != nil {
		return err, nil
	}

	data = GetMerchantForUserModel{
		MerId:       mer.ID,
		MerGuid:     mer.Guid,
		Name:        mer.Name,
		Description: mer.Description,
		TypeId:      mer.TypeId,
		IsAdmin:     meruser.IsAdmin,
	}

	return nil, &data
}

//获取商户详细地址信息
func QueryMerchantAddressInfoList(db *gorm.DB, q *QueryMerchantAddressInfoListModel) (err error, list []MerchantAddressInfoModel) {

	sqlSelect := `merchants.name as mer_name,
					merchant_addresses.merchant_id,
					merchant_details.phones,
					merchants.description,
					merchant_addresses.city,merchant_addresses.region,
					merchant_addresses.town,
					merchant_addresses.address,merchant_addresses.latitude,
					merchant_addresses.longitude,
				  merchant_addresses.name,merchant_addresses.remark`

	err = db.Table("merchants").Select(sqlSelect).
		Joins("inner join merchant_addresses on merchants.id = merchant_addresses.merchant_id").
		Joins("inner join merchant_details on merchants.id = merchant_details.merchant_id").
		Where("merchants.deleted_at is null and merchant_addresses.deleted_at is null and merchant_details.deleted_at is null and merchants.id in ( ? )",
			q.MerIds).
		Scan(&list).Error

	if err != nil {
		return err, list
	}

	return err, list
}

//获取商户列表
func GetMerchantList(db *gorm.DB, q *GetMerchantListModel, imgServiceUrl string) (error, *[]MerchantListModel, int) {

	var (
		err   error
		list  []MerchantListModel
		count int
	)
	q.PageSet()
	sqlSelect := `merchants.id as mer_id,merchants.name as name,merchants.description,
				  merchant_addresses.latitude,merchant_addresses.longitude,
  				  resources.url as image`

	tx := db.Table("merchants").Select(sqlSelect).
		Joins("inner join merchant_addresses on merchants.id = merchant_addresses.merchant_id").
		Joins("inner join merchant_details on merchants.id = merchant_details.merchant_id").
		Joins("LEFT JOIN merchant_resources on merchants.id=merchant_resources.merchant_id and merchant_resources.is_logo= ? ").
		Joins("LEFT JOIN resources on merchant_resources.resources_id=resources.id").
		Where(`merchants.deleted_at is null and merchant_addresses.deleted_at is null 
					 and merchant_details.deleted_at is null 
					 and merchant_resources.deleted_at is null `, utils.BoolToInt(true))

	if len(q.Text) > 0 {
		tx = tx.Where("merchants.name like ?", "%"+q.Text+"%")
	}

	err = tx.Count(&count).
		Offset(q.Offset).Limit(q.Limit).
		Scan(&list).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err, &list, count
	}

	computeDistance := false
	//判断是否需要根据经纬度 计算距离
	if q.Lat > 0 && q.Lon > 0 && q.ComputeDistance {
		computeDistance = true
	}
	for i := 0; i < len(list); i++ {
		list[i].Image = imgServiceUrl + list[i].Image
		if computeDistance {
			_, km := utils.Distance(utils.Coord{Lat: q.Lat, Lon: q.Lon},
				utils.Coord{Lat: list[i].Latitude, Lon: list[i].Longitude})
			list[i].Distance = km
		}
	}

	return nil, &list, count
}
