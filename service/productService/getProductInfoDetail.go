package productService

import (
	"github.com/jinzhu/gorm"
	merModel "github.com/ququgou-shop/modules/merchant/model"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/product/productEnum"
)

//获取单个商品详细信息
func GetProductDetailInfoSingle(db *gorm.DB,
	q *GetProductDetailInfoSingleModel,
	imgServiceUrl string) (error, *ProductDetailInfoSingle) {

	var (
		res           ProductDetailInfoSingle
		p             model.Product
		rs            *[]ProductResourceModel
		content       *model.ProductContent
		skuDS         *[]ProductSKUDetailModel
		attrs         *[]ProductAttributeModel
		tags          *[]ProductTagsModel
		deliveryTypes *[]ProductDeliveryTypeModel
		merInfo       *merModel.Merchant
		exitCollected bool
		err           error
	)

	//查看商品是否存在 和 获取商品信息
	err = db.Where("guid = ?", q.Guid).First(&p).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return err, nil
	}

	chSkuDS := make(chan *[]ProductSKUDetailModel)
	chErr := make(chan error)
	chAttrs := make(chan *[]ProductAttributeModel)
	chRs := make(chan *[]ProductResourceModel)
	chContent := make(chan *model.ProductContent)
	chTags := make(chan *[]ProductTagsModel)
	chDeliveryTypes := make(chan *[]ProductDeliveryTypeModel)
	chMerInfo := make(chan *merModel.Merchant)
	chExitCollected := make(chan bool)

	//defer close(chSkuDS)
	//defer close(chRs)
	//defer close(chAttrs)
	//defer close(chContent)
	//defer close(chTags)
	//defer close(chDeliveryTypes)
	//defer close(chErr)

	//获取商品sku 信息
	go getProductSKUDetailInfo(db, p.ID, imgServiceUrl, chSkuDS, chErr)

	//获取商品所有规格选项信息
	go getProductAttributeInfo(db, p.ID, chAttrs, chErr)

	//获取商品图片信息
	go getProductResourceInfo(db, p.ID, imgServiceUrl, chRs, chErr)

	//获取商品富文本信息
	go getProductContentInfo(db, p.ID, chContent, chErr)

	//获取标签
	go getProductTags(db, p.ID, chTags, chErr)

	//获取商品交付方式
	go getProductDeliveryTypes(db, p.ID, chDeliveryTypes, chErr)

	//获取商品商户信息
	go getProductMerInfo(db, p.MerId, chMerInfo, chErr)

	//获取商品收藏信息
	go exitProductCollected(db, p.ID, q.UserId, chExitCollected, chErr)

	skuDS = <-chSkuDS
	//err = <-chErr
	//if err != nil {
	//	return err, nil
	//}

	attrs = <-chAttrs
	//err = <-chErr
	//if err != nil {
	//	return err, nil
	//}

	rs = <-chRs
	//err = <-chErr
	//if err != nil {
	//	return err, nil
	//}

	content = <-chContent
	//err = <-chErr
	//if err != nil {
	//	return err, nil
	//}

	tags = <-chTags
	//err = <-chErr
	//
	//if err != nil {
	//	return err, nil
	//}
	deliveryTypes = <-chDeliveryTypes

	merInfo = <-chMerInfo

	exitCollected = <-chExitCollected
	//err = <-chErr
	//
	//if err != nil {
	//	return err, nil
	//}
	//if err != nil {
	//	return err, nil
	//}
	//err = <-chErr
	//chErr数量
	chErrCount := 8
	for e := range chErr {
		chErrCount--
		if e != nil {
			return e, nil
		}
		//fmt.Println(len(chErr))
		if chErrCount == 0 {
			break
		}
	}

	//TODO:商品类型判断
	//获取商品类型 如：是否是预定类型商品

	stock := 0

	for _, i := range *skuDS {
		stock += i.Stock
	}

	//商品信息封装打包出去
	res = ProductDetailInfoSingle{
		Guid:                p.Guid,
		TypeId:              p.TypeId,
		BrandId:             p.BrandId,
		Name:                p.Name,
		MerId:               p.MerId,
		MerName:             merInfo.Name,
		Description:         p.Description,
		Keywords:            p.Keywords,
		OriginalPrice:       p.OriginalPrice,
		CurrentPrice:        p.CurrentPrice,
		MinPrice:            p.MinPrice,
		Sales:               p.Sales,
		IsSingle:            p.IsSingle,
		Stock:               stock,
		Attributes:          attrs,
		SkuInfo:             skuDS,
		Resources:           rs,
		Content:             content.Content,
		Tags:                tags,
		ProductDeliveryType: deliveryTypes,
		Status:              p.Status,
		StatusText:          productEnum.ProductStatus(p.Status).Text(),
		Collected:           exitCollected,
	}

	return nil, &res
}

//先获取商品所有规格属性信息
//然后获取商品sku信息

func getProductAttributeInfo(db *gorm.DB, productId uint64, chPa chan *[]ProductAttributeModel, chErr chan error) {

	var (
		attr []productAttrOptionModel
		pa   []ProductAttributeModel
		err  error
	)

	sql := `SELECT a.id AS attribute_id, a.name AS attribute_name, ao.id AS attribute_option_id, ao.name AS attribute_option_name, ag.id AS attribute_group_id
FROM products p
	INNER JOIN attribute_groups ag ON p.id = ag.product_id
	INNER JOIN attribute_value_groups aog ON aog.attribute_group_id = ag.id
	INNER JOIN attribute_values ao ON ao.id = aog.attribute_value_id
	INNER JOIN attributes a ON a.id = ao.attribute_id
WHERE p.id = ?
	AND ag.deleted_at IS NULL
	AND aog.deleted_at IS NULL`

	err = db.Raw(sql, productId).Find(&attr).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		chPa <- &pa
		chErr <- err
		close(chPa)
		return
	}

	for _, v := range attr {

		p := ProductAttributeModel{}
		b := false

		if p.Options == nil {
			p.Options = []ProductAttributesOptionModel{}
		}

		p.Options = append(p.Options, ProductAttributesOptionModel{
			VId:  v.AttributeOptionId,
			Name: v.AttributeOptionName,
		})

		if attr != nil && len(attr) > 0 {
			for i := 0; i < len(pa); i++ {
				if pa[i].AId == v.AttributeId {
					b = true
					pa[i].Options = append(pa[i].Options, p.Options[0])
					continue
				}
			}
		}

		if !b || (attr == nil || len(attr) <= 0) {
			p.AId = v.AttributeId
			p.Name = v.AttributeName
			pa = append(pa, p)
		}
	}

	chPa <- &pa
	chErr <- nil
	close(chPa)
	return
}

//获取商品sku信息x
func getProductSKUDetailInfo(db *gorm.DB, productId uint64, imgServiceUrl string, chList chan *[]ProductSKUDetailModel, chErr chan error) {

	var (
		list []ProductSKUDetailModel
		err  error
		//AttributeId uint64 `json:"attribute_id"`
		//AttributeOptionId uint64 `json:"attribute_option_id"`
	)

	sql := `SELECT sku.id AS sku_id, sku.name AS sku_name, sku.image_json AS sku_image, sku.price AS price
	, sku.stock AS stock, sku.status AS status, sku.sort AS sort,sku.attribute_info,sku.is_single_attribute
FROM product_sku sku
	JOIN products p ON sku.product_id = p.id 
	where p.id =? and sku.deleted_at is null and p.deleted_at is null `

	//rows, err := db.Raw(sql, productId).Rows()
	err = db.Raw(sql, productId).Scan(&list).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		chList <- &list
		chErr <- err
		return
	}
	//for rows.Next() {
	//	var (
	//		v                 *ProductSKUDetailModel
	//		attributeId       uint64
	//		attributeOptionId uint64
	//		b                 bool
	//	)
	//
	//	v = &ProductSKUDetailModel{}
	//	//err=db.ScanRows(rows, &v)
	//	err = rows.Scan(&v.SkuId, &v.SkuName, &v.SkuImage, &v.Price, &v.Stock, &v.Status, &v.Sort, &v.AttributeGroupId, &attributeId, &attributeOptionId)
	//
	//	if err != nil {
	//		chList <- &list
	//		chErr <- err
	//		return
	//	}
	//
	//	b = false
	//	for i := 0; i < len(list); i++ {
	//		if v.SkuId == list[i].SkuId {
	//			v = &list[i]
	//			b = true
	//			continue
	//		}
	//	}
	//
	//	//err=rows.Scan(&attributeId, &attributeOptionId)
	//	//
	//	//if err!=nil {
	//	//	return err,nil
	//	//}
	//
	//	if v.PropPath == nil {
	//		v.PropPath = map[uint64]uint64{}
	//	}
	//	v.PropPath[attributeId] = attributeOptionId
	//
	//	if !b {
	//		list = append(list, *v)
	//	}
	//}

	//TODO:图片后面优化
	for i := 0; i < len(list); i++ {
		if list[i].SkuImage.Url != "" {
			list[i].SkuImage.Url = imgServiceUrl + list[i].SkuImage.Url
		}
	}

	chList <- &list
	chErr <- nil
	return
}

//获取商品图片信息
func getProductResourceInfo(db *gorm.DB, productId uint64, imgServiceUrl string, chList chan *[]ProductResourceModel, chErr chan error) {
	var (
		list []ProductResourceModel
		err  error
	)

	sql2 := `SELECT r.guid, concat(?,r.url) as url
FROM products p
	INNER JOIN product_resources pr ON p.id = pr.product_id
	INNER JOIN resources r ON r.id = pr.resource_id
where p.id =? AND p.deleted_at is null and pr.deleted_at is null and r.deleted_at is null
ORDER BY pr.cover DESC, pr.position 
`
	err = db.Raw(sql2, imgServiceUrl, productId).Find(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		chList <- &list
		chErr <- err
		close(chList)
		return
	}
	chList <- &list
	chErr <- nil
	close(chList)
	return
}

//获取商品富文本信息
func getProductContentInfo(db *gorm.DB, productId uint64, chContent chan *model.ProductContent, chErr chan error) {
	var (
		c   model.ProductContent
		err error
	)

	err = db.Where("product_id=?", productId).First(&c).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		chContent <- &c
		chErr <- err
		close(chContent)
		return
	}

	chContent <- &c
	chErr <- nil
	close(chContent)
	return
}

//获取商品所有的详细信息

//商品Tags 处理
func getProductTags(db *gorm.DB, productId uint64, chList chan *[]ProductTagsModel, chErr chan error) {

	var (
		m   []ProductTagsModel
		err error
	)

	//db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)

	sql := `select pt.id as id ,t.name as name from product_tags pt INNER JOIN tags t on pt.tag_id=t.id
		WHERE pt.deleted_at IS  NULL and t.deleted_at is NULL and pt.product_id = ? `
	err = db.Raw(sql, productId).Scan(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		chList <- &m
		chErr <- err
		close(chList)
		return
	}

	chList <- &m
	chErr <- nil
	close(chList)
	return
}

//商品交付方式获取
func getProductDeliveryTypes(db *gorm.DB, productId uint64, chList chan *[]ProductDeliveryTypeModel, chErr chan error) {
	var (
		m   []ProductDeliveryTypeModel
		err error
	)
	//db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)

	sql := `select pdt.delivery_type_id FROM product_delivery_type pdt 
		where pdt.deleted_at is NULL and product_id=?
		`
	err = db.Raw(sql, productId).Scan(&m).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		chList <- &m
		chErr <- err
		close(chList)
		return
	}

	if len(m) <= 0 {
		m = append(m, ProductDeliveryTypeModel{
			DeliveryTypeId: 5, //默认方式
			Name:           "线下自提",
		})
		m = append(m, ProductDeliveryTypeModel{
			DeliveryTypeId: 10, //默认方式
			Name:           "送货上门",
		})
	}

	chList <- &m
	chErr <- nil
	close(chList)
	return
}

//商品商户信息
func getProductMerInfo(db *gorm.DB, merId uint64, chList chan *merModel.Merchant, chErr chan error) {
	var (
		m   merModel.Merchant
		err error
	)
	err = db.Where("id=?", merId).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		chList <- &m
		chErr <- err
		close(chList)
		return
	}

	if err == gorm.ErrRecordNotFound {
		m = merModel.Merchant{}
	}

	chList <- &m
	chErr <- nil
	close(chList)
	return
}

//商品是否收藏
func exitProductCollected(db *gorm.DB, productId uint64, userId uint64, chExit chan bool, chErr chan error) {
	var (
		err   error
		count int
	)

	if userId <= 0 {
		chExit <- false
		chErr <- err
		close(chExit)
		return
	}

	err = db.Model(&model.ProductCollection{}).Where("user_id = ? and product_id = ? and deleted_at is null",
		userId, productId).Count(&count).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		chExit <- false
		chErr <- err
		close(chExit)
		return
	}

	chExit <- count > 0
	chErr <- nil
	close(chExit)
	return
}
