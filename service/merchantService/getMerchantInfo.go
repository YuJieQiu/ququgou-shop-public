package merchantService

import (
	"github.com/jinzhu/gorm"
	merEmun "github.com/ququgou-shop/modules/merchant/merEnum"
	"github.com/ququgou-shop/modules/merchant/model"
	uModel "github.com/ququgou-shop/modules/user/model"
)

//获取商户信息
func GetMerchantInfo(db *gorm.DB, q *GetMerchantInfoModel, imgServiceUrl string) (error, *MerInfoModel) {
	var (
		err        error
		info       MerInfoModel
		mer        model.Merchant
		merdetail  model.MerchantDetail
		merconf    model.MerchantConfig
		merlabel   []model.MerchantLabel
		meraddress model.MerchantAddress
	)

	var dbMer *gorm.DB
	if q.MerCode != "" {
		dbMer = db.Where("guid=? and status=?", q.MerCode, int(merEmun.Verified))
	} else if q.MerId > 0 {
		dbMer = db.Where("id=? and status=?", q.MerId, int(merEmun.Verified))
	} else {
		return err, nil
	}

	//TODO:测试 暂时不做验证
	//if q.IsWhereUserId {
	//	dbMer = dbMer.Where("user_id=? ", u.ID)
	//}

	err = dbMer.First(&mer).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).First(&merdetail).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).Order("sort").Find(&merlabel).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	//err=db.Where("merchant_id=?",mer.ID).Order("sort").Find(&merres).Error
	//if err!=nil {
	//	return err,nil
	//}

	err = db.Where("merchant_id=?", mer.ID).First(&merconf).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).First(&meraddress).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	labels := make(map[uint64]string)
	for _, v := range merlabel {
		labels[v.LabelId] = v.Text
	}

	//var merRes []MerResourceModel
	//for _,v:=range merres {
	//	merRes=append(merRes,MerResourceModel{
	//		Id:v.ID,
	//
	//	})
	//}
	logo := MerResourceModel{}
	err, merRes := getMerResourceInfo(db, &mer, imgServiceUrl)
	if err != nil {
		return err, nil
	}

	for _, v := range merRes {
		if v.IsLogo {
			logo = v
			//merRes = append(merRes[:i], merRes[i+1:]...)
			//merRes = append(merRes[:1])
			break
		}
	}

	info = MerInfoModel{
		Id:          mer.Guid,
		Name:        mer.Name,
		Description: mer.Description,
		BusinessTime: MerBusinessTime{
			StartTime: merdetail.BusinessStartTime,
			EndTime:   merdetail.BusinessEndTime,
		},
		Address: MerAddressModel{
			Address:   meraddress.Address,
			City:      meraddress.City,
			Region:    meraddress.Region,
			Town:      meraddress.Town,
			Latitude:  meraddress.Latitude,
			Longitude: meraddress.Longitude,
			Name:      meraddress.Name,
			Remark:    meraddress.Remark,
		},
		Label:         labels,
		Logo:          logo,
		Resources:     merRes,
		Phones:        merdetail.Phones,
		TypeId:        mer.TypeId,
		ActiveGoods:   merconf.ActiveGoods,
		ActiveService: merconf.ActiveService,
		ActiveComment: merconf.ActiveComment,
	}

	return nil, &info
}

//更新商户信息
func UpdateMerchantInfo(db *gorm.DB, q *MerInfoModel, u *uModel.User) (error, *model.Merchant) {

	var (
		err       error
		mer       model.Merchant
		merdetail model.MerchantDetail
		merconf   model.MerchantConfig
		merres    []model.MerchantResources
		//mertype model.MerchantType
		merlabel   []model.MerchantLabel
		meraddress model.MerchantAddress

		//newLable []model.MerchantLabel
	)

	//TODO://测试暂时不做验证
	err = db.Where("guid=? and status=?", q.Id, int(merEmun.Verified)).First(&mer).Error
	if err != nil {
		return err, nil
	}

	//err = db.Where("guid=? and user_id=? and status=?", q.Id, u.ID, int(merEmun.Verified)).First(&mer).Error
	//if err != nil {
	//	return err, nil
	//}

	err = db.Where("merchant_id=?", mer.ID).First(&merdetail).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).First(&merconf).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).First(&meraddress).Error
	if err != nil {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).Find(&merlabel).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	err = db.Where("merchant_id=?", mer.ID).Find(&merres).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return err, nil
	}

	mer.Name = q.Name
	mer.Description = q.Description

	//var labArr string
	//
	//for i,v:=range q.Label {
	//	labArr +=v+","
	//
	//	for k:=0;k<len(merlabel) ;k++  {
	//		if i==merlabel[k].ID {
	//			newLable=append(newLable,merlabel[k])
	//			merlabel=append(merlabel[:k],merlabel[k+1:]...)
	//			break
	//		}
	//	}
	//}
	//merdetail.Label=ext_struct.JsonStringArray(labArr)

	merdetail.BusinessStartTime = q.BusinessTime.StartTime
	merdetail.BusinessEndTime = q.BusinessTime.EndTime

	merdetail.Phones = q.Phones
	mer.TypeId = q.TypeId

	merconf.ActiveComment = q.ActiveComment
	merconf.ActiveGoods = q.ActiveGoods
	merconf.ActiveService = q.ActiveService

	meraddress.Name = q.Address.Name
	meraddress.Address = q.Address.Address
	meraddress.Remark = q.Address.Remark
	meraddress.City = q.Address.City
	meraddress.Region = q.Address.Region
	meraddress.Town = q.Address.Town
	meraddress.Latitude = q.Address.Latitude
	meraddress.Longitude = q.Address.Longitude
	meraddress.Active = true //默认启用

	tx := db.Begin()

	chErr := make(chan error)
	//defer close(chErr)

	if q.Logo.Id <= 0 && len(q.Resources) > 0 {
		q.Logo = q.Resources[0]
		q.Resources[0].IsLogo = true
	} else {
		q.Resources = append(q.Resources, q.Logo)
	}

	go merResourceProcess(tx, &merres, &q.Resources, &mer, chErr)
	go merLabelProcess(tx, &merlabel, &q.Label, &mer, chErr)

	count := 2

	for e := range chErr {
		count--
		if e != nil {
			tx.Rollback()
			return e, nil
		}
		if count == 0 {
			break
		}
	}

	err = tx.Save(&mer).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = tx.Save(&merdetail).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = tx.Save(&merconf).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = tx.Save(&meraddress).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	if err = tx.Commit().Error; err != nil {
		tx.Rollback()
		return err, nil
	}

	return nil, &mer
}

//商户图片资源处理
func merResourceProcess(tx *gorm.DB, res *[]model.MerchantResources, resModel *[]MerResourceModel, mer *model.Merchant, cerr chan error) {
	var (
		err       error
		newRes    []model.MerchantResources
		deleteRes []model.MerchantResources
		updateRes []model.MerchantResources
	)

	index := 0
	b := false
	for _, v := range *resModel {

		for _, vv := range *res {
			if v.Id == vv.ID {
				vv.Sort = index
				b = true
				deleteRes = append(deleteRes, vv)
				break
			}
		}

		if b {
			b = false
			continue
		}

		newRes = append(newRes, model.MerchantResources{
			MerchantId:  mer.ID,
			ResourcesId: v.Id,
			Cover:       v.Cover,
			Type:        v.Type,
			IsLogo:      v.IsLogo,
			Sort:        index,
		})

		b = false
		index++
	}

	for _, v := range *res {
		b = false
		for _, vv := range updateRes {
			if v.ID == vv.ID {
				b = true
				break
			}
		}

		if !b {
			deleteRes = append(deleteRes, v)
			continue
		}
	}

	for _, v := range newRes {
		err = tx.Create(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	for _, v := range updateRes {
		err = tx.Save(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	for _, v := range deleteRes {
		err = tx.Delete(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	cerr <- nil
	return
}

//商户标签处理
func merLabelProcess(tx *gorm.DB, merlabels *[]model.MerchantLabel, nlables *map[uint64]string, mer *model.Merchant, cerr chan error) {
	var (
		newLables    []model.MerchantLabel
		deleteLables []model.MerchantLabel
		updateLables []model.MerchantLabel
		err          error
	)

	index := 0
	b := false
	for i, v := range *nlables {

		for _, vv := range *merlabels {
			if i == vv.LabelId && v == vv.Text {
				vv.Sort = index
				b = true
				updateLables = append(updateLables, vv)
				break
			}
		}

		if b {
			b = false
			continue
		}

		newLables = append(newLables, model.MerchantLabel{
			MerchantId: mer.ID,
			LabelId:    i,
			Text:       v,
			Sort:       index,
		})

		b = false
		index++
	}

	for _, v := range *merlabels {
		b = false
		for _, vv := range updateLables {
			if v.ID == vv.ID {
				b = true
				break
			}
		}

		if !b {
			deleteLables = append(deleteLables, v)
			continue
		}
	}

	for _, v := range newLables {
		err = tx.Create(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	for _, v := range updateLables {
		err = tx.Save(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	for _, v := range deleteLables {
		err = tx.Delete(&v).Error
		if err != nil {
			cerr <- err
			return
		}
	}

	cerr <- nil
	return
}

//获取商户图片信息
func getMerResourceInfo(db *gorm.DB, mer *model.Merchant, imgServiceUrl string) (error, []MerResourceModel) {
	var (
		err  error
		data []MerResourceModel
	)

	err = db.Table("merchant_resources").Select("resources.id,resources.type,merchant_resources.cover,merchant_resources.sort,merchant_resources.is_logo,resources.url").Joins(" left join resources on merchant_resources.resources_id = resources.id ").
		Where("merchant_resources.merchant_id=? and merchant_resources.deleted_at is null", mer.ID).Order("merchant_resources.sort").Find(&data).Error

	if err != nil {
		return err, nil
	}

	for i := 0; i < len(data); i++ {
		data[i].Url = imgServiceUrl + data[i].Url
	}

	return nil, data
}

//Method Chaining where  Scopes S

func getMerchantInfoForMerchantWeb(u *uModel.User) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("and user_id=? ", u.ID)
	}

}

//Method Chaining where  Scopes E
