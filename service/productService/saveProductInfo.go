package productService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/product/model"
	modelResource "github.com/ququgou-shop/modules/resource/model"
)

//创建商品信息

func CreateProductInfo(db *gorm.DB, m *CreateProductInfoModel) (error, *model.Product) {

	var (
		err  error
		rid  []uint64
		rArr []modelResource.Resource
	)

	//商品 SKU 必须
	if m.SKU == nil || len(m.SKU) <= 0 {
		return ErrProductSKUEmpty, nil
	}

	p := model.Product{
		Guid:              utils.CreateUUID(),
		MerId:             m.MerId,
		BrandId:           m.BrandId,
		Name:              m.Name,
		TypeId:            m.TypeId,
		Status:            m.Status,
		Description:       m.Description,
		Keywords:          m.Keywords,
		OriginalPrice:     m.OriginalPrice,
		MinPrice:          m.MinPrice,
		MaxPrice:          m.MaxPrice,
		CurrentPrice:      m.CurrentPrice,
		Sales:             m.Sales,
		ProductType:       m.ProductType,
		Width:             m.Width,
		Height:            m.Height,
		Depth:             m.Depth,
		Weight:            m.Weight,
		Integral:          m.Integral,
		Active:            m.Active,
		IsSingle:          m.IsSingle,
		RecommendPriority: m.RecommendPriority,
	}

	for _, v := range m.Resources {
		rid = append(rid, v.ResourceId)
	}
	for _, v := range m.SKU {
		rid = append(rid, v.ResourceId)
	}

	rid = utils.RemoveRepeatedElemenForInt(rid)

	//查询Resource
	err = db.Where(rid).Find(&rArr).Error
	if err != nil {
		return err, nil
	}

	//开启事务 S
	tx := db.Begin()

	//创建product
	err = tx.Create(&p).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductContent(tx, p.ID, m.Content.Content, m.Content.ContentRemark, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductTags(tx, p.ID, m.Tags, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductCategory(tx, p.ID, m.CategoryIds, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductProperty(tx, p.ID, m.Property, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductResources(tx, p.ID, m.Resources, rArr, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductPaymentType(tx, p.ID, m.PaymentTypeIds, true)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = createProductSKU(tx, p.ID, &p, m, rArr)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	//修改 ImageJson 保存
	var images []ext_struct.JsonImage
	for _, v := range rArr {
		for _, rv := range m.Resources {
			if rv.ResourceId == v.ID {
				j := ext_struct.JsonImage{
					Guid: v.Guid,
					Path: v.Path,
					Url:  v.Url,
				}
				images = append(images, j)
			}
		}
	}
	p.ImageJson = images
	err = tx.Save(&p).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	//事务提交
	tx.Commit()

	return nil, &p
}

//更新商品信息
func UpdateProductInfo(db *gorm.DB, m *UpdateProductInfoModel) (error, *model.Product) {
	var (
		err  error
		rid  []uint64
		rArr []modelResource.Resource
		sku  []model.ProductSKU
		p    model.Product
	)
	//更新商品逻辑

	//商品 SKU 必须
	if m.SKU == nil || len(m.SKU) <= 0 {
		return ErrProductSKUEmpty, nil
	}

	//查询product
	if err := db.Where("id = ? and mer_id = ?", m.Id, m.MerId).First(&p).Error; err != nil {
		return err, nil
	}

	p.BrandId = m.BrandId
	p.Name = m.Name
	p.TypeId = m.TypeId
	p.Status = m.Status
	p.Description = m.Description
	p.Keywords = m.Keywords
	p.OriginalPrice = m.OriginalPrice
	p.MinPrice = m.MinPrice
	p.MaxPrice = m.MaxPrice
	p.CurrentPrice = m.CurrentPrice
	p.Sales = m.Sales
	p.ProductType = m.ProductType
	p.Width = m.Width
	p.Height = m.Height
	p.Depth = m.Depth
	p.Weight = m.Weight
	p.Integral = m.Integral
	p.Active = m.Active
	p.IsSingle = m.IsSingle
	p.RecommendPriority = m.RecommendPriority

	for _, v := range m.Resources {
		rid = append(rid, v.ResourceId)
	}
	for _, v := range m.SKU {
		rid = append(rid, v.ResourceId)
	}

	rid = utils.RemoveRepeatedElemenForInt(rid)

	//查询Resource
	err = db.Where(rid).Find(&rArr).Error
	if err != nil {
		return err, nil
	}

	if err = db.Where("product_id = ?", p.ID).Find(&sku).Error; err != nil {
		return err, nil
	}

	//开启事务 S
	tx := db.Begin()

	err = saveProductContent(tx, p.ID, m.Content.Content, m.Content.ContentRemark, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductTags(tx, p.ID, m.Tags, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductCategory(tx, p.ID, m.CategoryIds, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductProperty(tx, p.ID, m.Property, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductResources(tx, p.ID, m.Resources, rArr, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = saveProductPaymentType(tx, p.ID, m.PaymentTypeIds, false)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	err = updateProductSKU(tx, p.ID, &p, m, sku, rArr)
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	//修改 ImageJson 保存
	var images []ext_struct.JsonImage
	for _, v := range rArr {
		for _, rv := range m.Resources {
			if rv.ResourceId == v.ID {
				j := ext_struct.JsonImage{
					Guid: v.Guid,
					Path: v.Path,
					Url:  v.Url,
				}
				images = append(images, j)
			}
		}
	}
	p.ImageJson = images

	//保存更新
	err = tx.Save(&p).Error
	if err != nil {
		tx.Rollback()
		return err, nil
	}

	//事务提交
	tx.Commit()

	return nil, &p
}

//TODO: 优惠信息绑定、运费模板类型绑定 等等等.....

//创建 SKU
func createProductSKU(db *gorm.DB, productId uint64, p *model.Product, productCreateModel *CreateProductInfoModel, rArr []modelResource.Resource) error {
	var (
		err error
		//dbSKU []model.ProductSKU
	)
	sku := productCreateModel.SKU

	//sku
	for k, v := range sku {

		//if !productCreateModel.IsSingle {
		//	//attributeInfo, _ = json.Marshal(v.AttributeInfo)
		//}

		img := ext_struct.JsonImageString{}

		for _, rv := range rArr {
			if v.ResourceId == rv.ID {
				img = ext_struct.JsonImageString{
					Guid: rv.Guid,
					Path: rv.Path,
					Url:  rv.Url,
				}
			}
		}

		f := model.ProductSKU{
			Guid:              utils.CreateUUID(),
			ProductId:         productId,
			Name:              v.Name,
			Code:              v.Code,
			BarCode:           v.BarCode,
			OriginalPrice:     v.OriginalPrice,
			Price:             v.Price,
			Stock:             v.Stock,
			Sort:              k,
			Width:             If(v.Width == 0, p.Width, v.Width).(float32),
			Weight:            If(v.Weight == 0, p.Weight, v.Weight).(float32),
			Height:            If(v.Height == 0, p.Height, v.Height).(float32),
			Depth:             If(v.Depth == 0, p.Depth, v.Depth).(float32),
			ResourceId:        v.ResourceId,
			AttributeInfo:     v.AttributeInfo,
			IsSingleAttribute: v.IsSingleAttribute,
		}

		f.ImageJson = img

		//创建sku
		if err = db.Create(&f).Error; err != nil {
			return err
		}

		//商品 不是单规格
		if !productCreateModel.IsSingle {

			//单规格 则需要先创建 单规格的属性 (单规格使用的是AttributeID=0 的信息)
			//单规格 是指 只有一种属性 例如：颜色: 红 黄 蓝 绿 ....
			//多规格 是指 有多种属性 例如：颜色、尺码 等
			//if v.IsSingleAttribute {
			//	//先查询是否存在
			//	//err, av := shop.GetAttributeValue(db, &shop.GetAttributeValueModel{
			//	//	Name:        v.AttributeInfo[0].ValueName,
			//	//	MerId:       productCreateModel.MerId,
			//	//	AttributeId: v.AttributeInfo[0].Aid,
			//	//})
			//	//
			//	//if err != nil && !gorm.IsRecordNotFoundError(err) {
			//	//	chErr <- err
			//	//	return
			//	//}
			//
			//	//if err = db.Where("product_id = ?", p.ID).Find(&dbSKU).Error; err != nil {
			//	//	return err
			//	//}
			//
			//	av := &model.AttributeValue{
			//		Name:        v.AttributeInfo[0].ValueName,
			//		MerId:       productCreateModel.MerId,
			//		AttributeId: v.AttributeInfo[0].Aid,
			//	}
			//
			//	err = db.Where("name=? and attribute_id=? and (mer_id = ? or is_system = ? )",
			//		av.Name, av.AttributeId, av.MerId, true).First(&av).Error
			//
			//	if err != nil {
			//		if err != gorm.ErrRecordNotFound {
			//			return err
			//		} else {
			//			err = nil
			//		}
			//	}
			//
			//	if av != nil && av.ID > 0 {
			//		v.AttributeValueIds = []uint64{av.ID}
			//		v.AttributeInfo[0].Vid = av.ID
			//	} else {
			//		err, av = shop.CreateAttributeValue(db, &model.AttributeValue{
			//			MerId:       productCreateModel.MerId,
			//			Name:        v.AttributeInfo[0].ValueName,
			//			AttributeId: v.AttributeInfo[0].Aid,
			//			Sort:        k,
			//			Status:      0,
			//		})
			//		if err != nil {
			//			return err
			//		}
			//		v.AttributeValueIds = []uint64{av.ID}
			//		v.AttributeInfo[0].Vid = av.ID
			//	}
			//
			//	//更新sku AttributeInfo ？？有必要在最后更新
			//	f.AttributeInfo = v.AttributeInfo
			//	if err = db.Save(&f).Error; err != nil {
			//		return err
			//	}
			//}

			if len(v.AttributeInfo) > 0 {
				for sort, att := range v.AttributeInfo {
					var attAid uint64 = 0

					//不是单规格
					if !v.IsSingleAttribute && att.Aid > 0 {
						attAid = att.Aid
					}

					if att.Vid <= 0 {
						//获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录
						attValue := &model.AttributeValue{}

						err = db.Where("name=? and attribute_id=? and (mer_id = ? or is_system = ? )",
							att.ValueName, attAid, productCreateModel.MerId, true).First(attValue).Error

						if err != nil {
							if err != gorm.ErrRecordNotFound {
								return err
							} else {
								err = nil
							}
						}

						if attValue == nil || attValue.ID <= 0 {
							//创建
							err, attValue = createAttributeValue(db, &model.AttributeValue{
								MerId:       productCreateModel.MerId,
								Name:        att.ValueName,
								Status:      0,
								Sort:        sort,
								AttributeId: attAid,
							})
						}

						if err != nil {
							return err
						}

						v.AttributeInfo[sort].Vid = attValue.ID
					}

					//Aid Vid 如果 vid 为0  通过数据库查询是否存在，存在返回 vid

					//这个sku 是单规格的，不用添加 Attribute，默认 AttributeID=0
				}

			}

			//更新sku AttributeInfo ？？有必要在最后更新
			f.AttributeInfo = v.AttributeInfo
			if err = db.Save(&f).Error; err != nil {
				return err
			}

			//更新 通过 AttributeInfo 来获取规格属性信息

			//添加sku 规格属性组
			ag := model.AttributeGroup{
				ProductId: p.ID,
				Sort:      k,
			}

			err = db.Create(&ag).Error
			if err != nil {
				return err
			}

			//添加sku 规格属性
			for _, vv := range v.AttributeInfo {
				aog := model.AttributeValueGroup{
					AttributeGroupId: ag.ID,
					AttributeValueId: vv.Vid,
				}

				if err = db.Create(&aog).Error; err != nil {
					return err
				}
			}

			ags := model.AttributeGroupSKU{
				ProductSkuId:     f.ID,
				AttributeGroupId: ag.ID,
			}

			if err = db.Create(&ags).Error; err != nil {
				return err
			}
		}

		//添加商品sku价格信息 记录价格变化
		pp := model.ProductPriceLog{
			ProductId:       p.ID,
			ProductSkuId:    f.ID,
			Price:           f.Price,
			OperatorAdminId: 0,
		}
		if err = db.Create(&pp).Error; err != nil {
			return err
		}
	}
	return err
}

//更新 SKU
func updateProductSKU(db *gorm.DB, productId uint64, p *model.Product, productUpdateModel *UpdateProductInfoModel, dbSKU []model.ProductSKU, rArr []modelResource.Resource) error {
	var (
		err error
		//dbSKU        []model.ProductSKU
		deleteSKUIds []uint64 //需要删除的 sku id
	)

	//if err = db.Where("product_id = ?", productId).Find(&dbSKU).Error; err != nil {
	//	chErr <- err
	//	return
	//}
	skuModel := productUpdateModel.SKU
	index := 0
	b := false
	for _, v := range skuModel {
		i := 0

		var sku *model.ProductSKU

		for k, vv := range dbSKU {
			if v.Id == vv.ID {
				b = true
				i = k
				break
			}
		}

		img := ext_struct.JsonImageString{}
		//图片资源
		for _, rv := range rArr {
			if v.ResourceId == rv.ID {
				img = ext_struct.JsonImageString{
					Guid: rv.Guid,
					Path: rv.Path,
					Url:  rv.Url,
				}
			}
		}

		if b {
			//更新

			dbSKU[i].Name = v.Name
			dbSKU[i].Code = v.Code
			dbSKU[i].BarCode = v.BarCode
			dbSKU[i].OriginalPrice = v.OriginalPrice
			dbSKU[i].Price = v.Price
			dbSKU[i].Stock = v.Stock
			dbSKU[i].Sort = index
			dbSKU[i].Width = If(v.Width == 0, p.Width, v.Width).(float32)
			dbSKU[i].Weight = If(v.Weight == 0, p.Weight, v.Weight).(float32)
			dbSKU[i].Height = If(v.Height == 0, p.Height, v.Height).(float32)
			dbSKU[i].Depth = If(v.Depth == 0, p.Depth, v.Depth).(float32)
			dbSKU[i].ResourceId = v.ResourceId
			dbSKU[i].AttributeInfo = v.AttributeInfo
			dbSKU[i].ImageJson = img
			dbSKU[i].IsSingleAttribute = v.IsSingleAttribute

			if err = db.Save(&dbSKU[i]).Error; err != nil {
				return err
			}

			//删除 规格属性组记录
			agSku := model.AttributeGroupSKU{}
			if err = db.Where("product_sku_id= ?", dbSKU[i].ID).First(&agSku).Error; err != nil {
				if err != gorm.ErrRecordNotFound {
					return err
				} else {
					err = nil
				}
			}
			if agSku.ID > 0 {
				if err = db.Delete(model.AttributeGroup{}, "id = ?", agSku.AttributeGroupId).Error; err != nil {
					return err
				}
				if err = db.Delete(model.AttributeValueGroup{}, "attribute_group_id = ?", agSku.AttributeGroupId).Error; err != nil {
					return err
				}
				if err = db.Delete(model.AttributeGroupSKU{}, "id = ?", agSku.ID).Error; err != nil {
					return err
				}
			}

			sku = &dbSKU[i]
		} else {
			//创建
			f := model.ProductSKU{
				Guid:              utils.CreateUUID(),
				ProductId:         productId,
				Name:              v.Name,
				Code:              v.Code,
				BarCode:           v.BarCode,
				OriginalPrice:     v.OriginalPrice,
				Price:             v.Price,
				Stock:             v.Stock,
				Sort:              index,
				Width:             If(v.Width == 0, p.Width, v.Width).(float32),
				Weight:            If(v.Weight == 0, p.Weight, v.Weight).(float32),
				Height:            If(v.Height == 0, p.Height, v.Height).(float32),
				Depth:             If(v.Depth == 0, p.Depth, v.Depth).(float32),
				ResourceId:        v.ResourceId,
				AttributeInfo:     v.AttributeInfo,
				IsSingleAttribute: v.IsSingleAttribute,
			}
			f.ImageJson = img

			if err = db.Create(&f).Error; err != nil {
				return err
			}
			sku = &f
		}

		if !productUpdateModel.IsSingle {

			//单规格 则需要先创建 单规格的属性 (单规格使用的是AttributeID=0 的信息)
			//单规格 是指 只有一种属性 例如：颜色: 红 黄 蓝 绿 ....
			//多规格 是指 有多种属性 例如：颜色、尺码 等
			//if v.IsSingleAttribute {
			//
			//	//先查询是否存在
			//	err, av := shop.GetAttributeValue(db, &shop.GetAttributeValueModel{
			//		Name:        v.AttributeInfo[0].ValueName,
			//		MerId:       productUpdateModel.MerId,
			//		AttributeId: v.AttributeInfo[0].Aid,
			//	})
			//
			//	if err != nil && !gorm.IsRecordNotFoundError(err) {
			//		return err
			//	}
			//
			//	if av != nil && av.ID > 0 {
			//		v.AttributeValueIds = []uint64{av.ID}
			//		v.AttributeInfo[0].Vid = av.ID
			//	} else {
			//		err, av = shop.CreateAttributeValue(db, &model.AttributeValue{
			//			MerId:       productUpdateModel.MerId,
			//			Name:        v.AttributeInfo[0].ValueName,
			//			AttributeId: v.AttributeInfo[0].Aid,
			//			Sort:        index,
			//			Status:      0,
			//		})
			//		if err != nil {
			//			return err
			//		}
			//		v.AttributeValueIds = []uint64{av.ID}
			//		v.AttributeInfo[0].Vid = av.ID
			//	}
			//	sku.AttributeInfo = v.AttributeInfo
			//	if err = db.Save(&sku).Error; err != nil {
			//		return err
			//	}
			//}

			if len(v.AttributeInfo) > 0 {
				for sort, att := range v.AttributeInfo {
					var attAid uint64 = 0

					//不是单规格
					if !v.IsSingleAttribute && att.Aid > 0 {
						attAid = att.Aid
					}

					if att.Vid <= 0 {
						//获取匹配的第一条记录, 否则根据给定的条件创建一个新的记录
						attValue := &model.AttributeValue{}

						err = db.Where("name=? and attribute_id=? and (mer_id = ? or is_system = ? )",
							att.ValueName, attAid, productUpdateModel.MerId, true).First(attValue).Error

						if err != nil {
							if err != gorm.ErrRecordNotFound {
								return err
							} else {
								err = nil
							}
						}

						if attValue == nil || attValue.ID <= 0 {
							//创建
							err, attValue = createAttributeValue(db, &model.AttributeValue{
								MerId:       productUpdateModel.MerId,
								Name:        att.ValueName,
								Status:      0,
								Sort:        sort,
								AttributeId: attAid,
							})
						}

						if err != nil {
							return err
						}

						v.AttributeInfo[sort].Vid = attValue.ID
					}

					//Aid Vid 如果 vid 为0  通过数据库查询是否存在，存在返回 vid

					//这个sku 是单规格的，不用添加 Attribute，默认 AttributeID=0
					//更新sku AttributeInfo
					sku.AttributeInfo = v.AttributeInfo
					if err = db.Save(&sku).Error; err != nil {
						return err
					}
				}

			}

			//添加sku 规格属性组
			ag := model.AttributeGroup{
				ProductId: p.ID,
				Sort:      index,
			}
			err = db.Create(&ag).Error
			if err != nil {
				return err
			}

			//添加sku 规格属性
			for _, vv := range v.AttributeInfo {
				aog := model.AttributeValueGroup{
					AttributeGroupId: ag.ID,
					AttributeValueId: vv.Vid,
				}

				if err = db.Create(&aog).Error; err != nil {
					return err
				}
			}

			ags := model.AttributeGroupSKU{
				ProductSkuId:     sku.ID,
				AttributeGroupId: ag.ID,
			}
			if err = db.Create(&ags).Error; err != nil {
				return err
			}
		}

		//添加商品sku价格信息
		pp := model.ProductPriceLog{
			ProductId:       p.ID,
			ProductSkuId:    sku.ID,
			Price:           sku.Price,
			OperatorAdminId: 0,
		}
		if err = db.Create(&pp).Error; err != nil {
			return err
		}

		b = false
		i = 0
		index++
	}

	//删除无效的sku
	for _, v := range dbSKU {
		b := false
		for _, vv := range skuModel {
			if v.ID == vv.Id {
				b = true
				break
			}
		}
		if !b {
			//删除
			deleteSKUIds = append(deleteSKUIds, v.ID)
		}
	}

	if len(deleteSKUIds) > 0 {
		if err = db.Where("id in (?)", deleteSKUIds).Delete(&model.ProductSKU{}).Error; err != nil {
			return err
		}
		//删除 规格属性组记录
		if err = db.Delete(model.AttributeGroupSKU{}, "product_sku_id in (?)", deleteSKUIds).Error; err != nil {
			return err
		}
		if err = db.Delete(model.AttributeGroup{}, "id in (select attribute_group_id from attribute_group_sku where product_sku_id in (?))", deleteSKUIds).Error; err != nil {
			return err
		}
		if err = db.Delete(model.AttributeValueGroup{}, "attribute_group_id in (select attribute_group_id from attribute_group_sku where product_sku_id in (?))", deleteSKUIds).Error; err != nil {
			return err
		}
	}

	return err
}

//保存 Product Resources
func saveProductResources(db *gorm.DB, productId uint64, resources []CreateProductResourceModel, rArr []modelResource.Resource, isNewRecord bool) error {
	var (
		err error
	)

	if !isNewRecord {
		//更新 先删除原先记录
		err = db.Where("product_id = ?", productId).Delete(model.ProductResource{}).Error
		if err != nil {
			return err
		}
	}

	//新增
	for i, v := range resources {

		for _, r := range rArr {
			if r.ID == v.ResourceId {
				err = db.Save(&model.ProductResource{
					ResourceId:   r.ID,
					ResourceGuid: r.Guid,
					ProductId:    productId,
					Sort:         i,
					Cover:        v.Cover,
					Type:         v.Type,
					Position:     v.Position,
				}).Error

				if err != nil {
					return err
				}
				break
			}
		}
	}

	return err
}

//保存 Product Property
func saveProductProperty(db *gorm.DB, productId uint64, property []CreateProductPropertyModel, isNewRecord bool) error {
	var (
		err error
	)

	if !isNewRecord {
		//更新 先删除原先记录
		err = db.Where("product_id = ?", productId).Delete(model.PropertyValue{}).Error
		if err != nil {
			return err
		}
	}

	//新增
	for _, v := range property {
		err = db.Save(&model.PropertyValue{
			ProductId:  productId,
			PropertyId: v.PropertyId,
			Name:       v.Name,
			Title:      v.Title,
		}).Error

		if err != nil {
			return err
		}
	}

	return err
}

//保存 Product Category
func saveProductCategory(db *gorm.DB, productId uint64, categoryIds []uint64, isNewRecord bool) error {
	var (
		err error
	)

	if !isNewRecord {
		//更新 先删除原先记录
		err = db.Where("product_id = ?", productId).Delete(model.CategoryProduct{}).Error
		if err != nil {
			return err
		}
	}

	//新增
	for _, v := range categoryIds {
		err = db.Save(&model.CategoryProduct{
			ProductId:  productId,
			CategoryId: v,
		}).Error

		if err != nil {
			return err
		}
	}
	return err
}

//保存 ProductTags
func saveProductTags(db *gorm.DB, productId uint64, tagsModel []CreateProductTagsModel, isNewRecord bool) error {
	var (
		err error
	)

	if !isNewRecord {
		//更新 先删除原先记录
		err = db.Where("product_id = ?", productId).Delete(model.ProductTags{}).Error
		if err != nil {
			return err
		}
	}

	//新增
	for k, v := range tagsModel {
		err = db.Save(&model.ProductTags{
			TagId:     v.TagId,
			ProductId: productId,
			Sort:      k,
		}).Error

		if err != nil {
			return err
		}
	}

	return err
}

//保存 产品Content
func saveProductContent(db *gorm.DB, productId uint64, content string, remark string, isNewRecord bool) error {
	var (
		err error
		c   model.ProductContent
	)

	//创建新记录
	if isNewRecord {
		c.Content = content
		c.Remark = remark
		c.ProductId = productId

		err = db.Create(&c).Error
	} else { //更新
		err = db.Model(&c).Where("product_id=?",
			productId).Updates(model.ProductContent{Content: content, Remark: remark}).Error

		//db.Where("product_id=?", productId).First(&c)
		//
		//c.Content = content
		//c.Remark = remark
		//c.ProductId = productId

		//err = db.Update(&c).Error
	}

	//if db.Where("product_id=?", productId).First(&c).RecordNotFound() {
	//	err = db.Create(&c).Error
	//} else {
	//	err = db.Save(&c).Error
	//}
	return err
}

func saveProductPaymentType(db *gorm.DB, productId uint64, payTypeIds []uint64, isNewRecord bool) error {
	var (
		err error
		//c   paymentModel.PaymentType
	)

	//不是创建新记录
	if !isNewRecord {
		//更新 先删除 在创建
		err = db.Where(" product_id=?", productId).Delete(&model.ProductPaymentType{}).Error
	}

	for _, id := range payTypeIds {
		var c model.ProductPaymentType
		c.ProductId = productId
		c.PaymentTypeId = id
		err = db.Create(&c).Error
		if err != nil {
			return err
		}
	}

	return err
}

//模拟三元表达式
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
