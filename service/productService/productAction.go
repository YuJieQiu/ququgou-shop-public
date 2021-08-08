package productService

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/product/productEnum"
	modelResource "github.com/ququgou-shop/modules/resource/model"

	"strconv"
)

//获取商品信息列表 //用户商品列表展示使用
func GetProductSmallInfoList(db *gorm.DB,
	q *GetProductSmallInfoListModel,
	isPage bool,
	order string,
	imgServiceUrl string) (err error, list []ProductSmallInfoModel, count int) {

	q.PageSet()

	//sql 语句查询

	sqlSelectCount := " select count(p.id) as count "

	sqlSelect := `select p.guid,p.type_id,p.current_price,p.brand_id,p.name,p.description,p.keywords,p.original_price,p.min_price,p.max_price,p.sales,r.url as image,
				(select count(stock) from product_sku where product_sku.deleted_at is null and product_id=p.id) as stock,
				m.id as mer_id,m.name as mer_name,md.latitude,md.longitude`

	//pr.cover=1 保证 一个product 只有 一个封面图片
	sql := ` from
		products p LEFT JOIN product_resources pr on p.id=pr.product_id and pr.cover=1 and pr.type=0
		LEFT JOIN merchants m on m.id=p.mer_id
		LEFT JOIN merchant_addresses md on md.merchant_id=p.mer_id
		LEFT JOIN resources r on pr.resource_id=r.id
		where p.active=1 and p.status=1 and p.deleted_at is null 
			  and pr.deleted_at is null and r.deleted_at is null
			  and md.deleted_at is null
		`

	//where 条件

	idsStr := utils.ArrayToString(q.ProductIds, ",")
	if q.ProductIds != nil && len(q.ProductIds) > 0 {
		sql = sql + " and p.id in ( " + idsStr + " ) "
	}

	if q.MerId > 0 {
		sql = sql + " and p.mer_id= " + strconv.FormatUint(q.MerId, 10)
	}

	switch order {
	case "id":
		order = " Order By p.id desc "
	case "stock":
		order = " Order By stock desc "
	case "created_at":
		order = " Order By p.created_at desc "
	default:
		if len(idsStr) > 0 {
			order = fmt.Sprintf(" Order By field ( p.id, %v )", idsStr)
		} else {
			order = " Order By p.id"
		}
	}

	if isPage {
		err = db.Raw(sqlSelectCount + sql).Count(&count).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, list, count
			}
			return err, list, count
		}
	}

	if isPage {
		err = db.Raw(sqlSelect + sql + order).Offset(q.Offset).Limit(q.Limit).Scan(&list).Error
	} else {
		err = db.Raw(sqlSelect + sql + order).Scan(&list).Error
	}

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, list, count
		}
		return err, list, count
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

	return nil, list, count
}

//Product Info E
//获取商品列表
func GetProductList(db *gorm.DB, q *GetProductListModel, isPage bool, imgServiceUrl string) (
	err error, list []model.Product, count int) {

	//获取商品列表逻辑 ：根据现有业务 这个方法获取商品基本信息，展示在前端，如需要更详细的信息应该有个单个商品详细信息

	q.PageSet()

	w := model.Product{
		Name:  q.Name,
		MerId: q.MerId,
	}

	if q.Status != int(productEnum.ProductStatusDefault) {
		w.Status = q.Status
	}

	if isPage {
		err = db.Model(&model.Product{}).Where(&w).Count(&count).Error

		if err != nil || count <= 0 {
			return err, list, count
		}
	}

	if isPage {
		err = db.Where(&w).Order("created_at desc").Offset(q.Offset).
			Limit(q.Limit).Find(&list).Error
	} else {
		err = db.Where(&w).Order("created_at desc").Find(&list).Error
	}

	if err != nil {
		return err, list, count
	}

	//TODO:图片后面优化
	for i := 0; i < len(list); i++ {
		list[i].SetImagesUrl(imgServiceUrl)
		//time
		list[i].CreatedTime = ext_struct.JsonTime(list[i].CreatedAt)
		list[i].StatusText = productEnum.ProductStatus(list[i].Status).Text()
	}

	return nil, list, count
}

//get product all info (获取product 所有信息)
func GetProductInfo(db *gorm.DB, q *GetProductInfoModel, imgServiceUrl string) (error, *ProductDomainModel) {
	var (
		err                 error
		product             ProductDomainModel
		content             model.ProductContent
		sku                 []model.ProductSKU
		productResources    []model.ProductResource
		resources           []modelResource.Resource
		propertyValue       []model.PropertyValue
		property            []model.Property
		productTags         []model.ProductTags
		tags                []model.Tags
		category            []model.Category
		productCategory     []model.CategoryProduct
		attributeValue      []model.AttributeValue
		attributes          []model.Attribute
		productPaymentTypes []model.ProductPaymentType
		productId           uint64
		//attributeGroup      []model.AttributeGroup
		//attributeValueGroup []model.AttributeValueGroup
	)

	productId = q.ProductId

	//product
	if err = db.Where("id=? and mer_id=?", productId, q.MerId).First(&product).Error; err != nil {
		return err, nil
	}

	//TODO:图片后面优化
	product.SetImagesUrl(imgServiceUrl)

	//product content
	db.Where("product_id = ?", productId).First(&content)
	product.Content = &content

	// sku
	if err = db.Where("product_id = ?", productId).Find(&sku).Error; err != nil {
		return err, nil
	}

	//TODO:图片后面优化
	for i := 0; i < len(sku); i++ {
		sku[i].SetImagesUrl(imgServiceUrl)
	}

	//sku AttributeValues

	product.SKU = sku

	// property
	db.Where("product_id = ?", productId).Find(&propertyValue)

	if len(propertyValue) > 0 {
		var propertyIds []uint64
		for _, v := range propertyValue {
			propertyIds = append(propertyIds, v.PropertyId)
		}
		propertyIds = utils.RemoveRepeatedElemenForInt(propertyIds)
		db.Where("id in (?)", propertyIds).Find(&property)

		for i := 0; i < len(property); i++ {
			for v := 0; v < len(propertyValue); v++ {
				if propertyValue[v].PropertyId == property[i].ID {
					property[i].Values = append(property[i].Values, propertyValue[v])
					continue
				}
			}
		}
	}
	product.Property = property

	//resources
	db.Where("product_id = ?", productId).Find(&productResources)
	if len(productResources) > 0 {
		var rids []uint64
		for _, v := range productResources {
			rids = append(rids, v.ResourceId)
		}
		rids = utils.RemoveRepeatedElemenForInt(rids)
		db.Where("id in (?)", rids).Find(&resources)

		for i := 0; i < len(productResources); i++ {
			for v := 0; v < len(resources); v++ {
				if productResources[i].ResourceId == resources[v].ID {

					//TODO:图片后面优化
					resources[v].SetImagesUrl(imgServiceUrl)

					productResources[i].Resource = resources[v]
					continue
				}
			}
		}
	}
	product.Resources = productResources

	//tags
	db.Where("product_id = ?", productId).Find(&productTags)
	if len(productTags) > 0 {
		var tids []uint64
		for _, v := range productTags {
			tids = append(tids, v.TagId)
		}
		tids = utils.RemoveRepeatedElemenForInt(tids)
		db.Where("id in (?)", tids).Find(&tags)

		for i := 0; i < len(productTags); i++ {
			for v := 0; v < len(tags); v++ {
				if productTags[i].TagId == tags[v].ID {
					productTags[i].Tag = tags[v]
					continue
				}
			}
		}
	}
	product.Tags = productTags

	//category
	db.Where("product_id = ?", productId).Find(&productCategory)
	if len(productCategory) > 0 {
		var cids []uint64
		for _, v := range productCategory {
			cids = append(cids, v.CategoryId)
		}
		cids = utils.RemoveRepeatedElemenForInt(cids)
		db.Where("id in (?)", cids).Find(&category)

		for _, k := range category {
			product.CategoryIds = append(product.CategoryIds, k.ID)
			product.CategoryInfos = append(product.CategoryInfos,
				ProductDomainCategoryInfoModel{
					Id:   k.ID,
					Name: k.Name,
				})
		}

		//for i := 0; i < len(productCategory); i++ {
		//	for v := 0; v < len(category); v++ {
		//		if productCategory[i].CategoryId == category[v].ID {
		//			productCategory[i].Category = category[v]
		//			continue
		//		}
		//	}
		//}
	}
	//product.CategoryProduct = productCategory
	//if productCategory != nil {
	//	for _, k := range productCategory {
	//		product.CategoryIds = append(product.CategoryIds, k.ID)
	//		product.CategoryInfos = append(product.CategoryInfos,
	//			domainmodel.ProductDomainCategoryInfoModel{
	//				Id:   k.CategoryId,
	//				Name: k.Category.Name,
	//			})
	//	}
	//
	//}

	//TODO: 暂时不返回 Attributes 参数
	//attribute Value
	db.Joins("JOIN attribute_value_groups ON attribute_value_groups.attribute_value_id =  attribute_values.id ").
		Joins("JOIN attribute_groups ON attribute_groups.id = attribute_value_groups.attribute_group_id").
		Where("attribute_groups.product_id = ? and attribute_values.deleted_at is null and attribute_value_groups.deleted_at is null", productId).
		Find(&attributeValue)

	if len(attributeValue) > 0 {
		var aids []uint64
		for _, v := range attributeValue {
			aids = append(aids, v.AttributeId)
		}
		aids = utils.RemoveRepeatedElemenForInt(aids)
		db.Where("id in (?)", aids).Find(&attributes)

		for i := 0; i < len(attributes); i++ {
			for v := 0; v < len(attributeValue); v++ {
				if attributeValue[v].AttributeId == attributes[i].ID {
					attributes[i].Values = append(attributes[i].Values, attributeValue[v])
					continue
				}
			}
		}
	}

	product.Attributes = attributes

	db.Where("product_id=?", productId).Find(&productPaymentTypes)
	if len(productPaymentTypes) > 0 {
		for _, i := range productPaymentTypes {
			tt := fmt.Sprint(i.PaymentTypeId)
			product.PaymentTypeIds = append(product.PaymentTypeIds, tt)
		}
	}

	return nil, &product
}

//更新商品状态
func UpdateProductStatus(db *gorm.DB, m *UpdateProductStatusModel) error {
	//商品下架
	err := db.Model(&model.Product{}).Where("id = ? and mer_id=?", m.ProductId, m.MerId).
		Update("status", int(m.Status)).Error

	return err
}
