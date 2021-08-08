package shopCartService

import (
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/library/utils"
	merchantModel "github.com/ququgou-shop/modules/merchant/model"
	productModel "github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/modules/product/productEnum"
	"github.com/ququgou-shop/modules/shopcart/model"
)

//添加购物车产品
func AddShopCartItem(db *gorm.DB, q *AddShopCartItemModel) (error, *model.ShopCart) {
	var (
		err error
		sc  model.ShopCart
	)

	if err := db.Where("user_id=? and product_sku_id=?", q.UserId, q.ProductSkuId).First(&sc).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	if sc.ID > 0 {
		//更新数量
		sc.Number += q.Number

		err = db.Save(sc).Error
		if err != nil {
			return err, nil
		}

		return nil, &sc
	}

	sc = model.ShopCart{
		ProductNo:      q.ProductNo,
		ProductSkuId:   q.ProductSkuId,
		UserId:         q.UserId,
		MerId:          q.MerId,
		Number:         q.Number,
		JoinPrice:      q.Price,
		JoinTotalPrice: q.TotalPrice,
	}

	//TODO: 库存校验
	err = db.Create(&sc).Error
	if err != nil {
		return err, nil
	}

	return nil, &sc
}

//获取用户购物车商品
func GetUserShopCartProductInfo(db *gorm.DB, q *GetUserShopCartProductInfoModel, imgServiceUrl string) (error, *[]UserShopCartMerModel) {
	var (
		err             error
		errChan         chan error
		scart           []model.ShopCart
		skuListChan     chan *[]productModel.ProductSKU
		skuList         *[]productModel.ProductSKU
		productListChan chan *[]productModel.Product
		productList     *[]productModel.Product
		merListChan     chan *[]merchantModel.Merchant
		merList         *[]merchantModel.Merchant
		skuisd          []uint64
		pguids          []string
		merids          []uint64
		list            []UserShopCartMerModel
	)

	skuListChan = make(chan *[]productModel.ProductSKU)
	productListChan = make(chan *[]productModel.Product)
	merListChan = make(chan *[]merchantModel.Merchant)
	errChan = make(chan error)

	defer close(skuListChan)
	defer close(productListChan)
	defer close(merListChan)

	if err := db.Where("user_id=?", q.UserId).Find(&scart).Error; err != nil && err != gorm.ErrRecordNotFound {
		return err, nil
	}

	for _, i := range scart {
		skuisd = append(skuisd, i.ProductSkuId)
		pguids = append(pguids, i.ProductNo)
		merids = append(merids, i.MerId)
	}

	pguids = utils.RemoveRepeatedElementForString(pguids)
	merids = utils.RemoveRepeatedElemenForInt(merids)

	go func() {
		var s []productModel.ProductSKU
		var e error
		if e = db.Where("id in (?) ", skuisd).Find(&s).Error; e != nil && e != gorm.ErrRecordNotFound {
			skuListChan <- &s
			errChan <- e
			return
		}
		skuListChan <- &s
		errChan <- e
	}()

	go func() {
		var s []productModel.Product
		var e error
		if e = db.Where("guid in (?) ", pguids).Find(&s).Error; e != nil && e != gorm.ErrRecordNotFound {
			productListChan <- &s
			errChan <- e
			return
		}
		productListChan <- &s
		errChan <- e
	}()

	go func() {
		var s []merchantModel.Merchant
		var e error
		if e = db.Where("id in (?) ", merids).Find(&s).Error; e != nil && e != gorm.ErrRecordNotFound {
			merListChan <- &s
			errChan <- e
			return
		}
		merListChan <- &s
		errChan <- e
	}()

	skuList = <-skuListChan
	err = <-errChan
	if err != nil {
		return err, nil
	}

	productList = <-productListChan
	err = <-errChan
	if err != nil {
		return err, nil
	}

	merList = <-merListChan
	err = <-errChan
	if err != nil {
		return err, nil
	}

	p_first := func(no string, l *[]productModel.Product) *productModel.Product {
		for _, i := range *l {
			if i.Guid == no {
				return &i
			}
		}
		return &productModel.Product{}
	}

	sku_first := func(id uint64, l *[]productModel.ProductSKU) *productModel.ProductSKU {
		for _, i := range *l {
			if i.ID == id {
				return &i
			}
		}
		return &productModel.ProductSKU{}
	}

	mer_first := func(id uint64, l *[]merchantModel.Merchant) *merchantModel.Merchant {
		for _, i := range *l {
			if i.ID == id {
				return &i
			}
		}
		return &merchantModel.Merchant{}
	}

	for _, i := range scart {
		p := p_first(i.ProductNo, productList)
		s := sku_first(i.ProductSkuId, skuList)
		m := mer_first(i.MerId, merList)
		img := ext_struct.JsonImage{}
		if s.ResourceId <= 0 {
			img = p.ImageJsonModel.ImageJson[0]
		} else {
			img = ext_struct.JsonImage(s.ImageJsonSingleModel.ImageJson)
		}

		img.Url = imgServiceUrl + img.Url

		usp := UserShopCartProductModel{
			CartID:         i.ID,
			ProductSkuId:   i.ProductSkuId,
			ProductNo:      i.ProductNo,
			Number:         i.Number,
			JoinPrice:      i.JoinPrice,
			JoinTotalPrice: i.JoinTotalPrice,
			JoinTime:       i.CreatedAt,
			Name:           p.Name,
			Description:    p.Description,
			OriginalPrice:  s.OriginalPrice,
			Price:          s.Price,
			Stock:          s.Stock,
			AttributeInfo:  s.AttributeInfo,
			ProductStatus:  p.Status,
			SkuStatus:      s.Status,
			Img:            img,
		}

		//if len(list) == 0 {
		//	list = append(list, UserShopCartMerModel{
		//		MerId:   i.MerId,
		//		MerName: m.Name,
		//		MerCode: m.Guid,
		//	})
		//if productEnum.ProductStatus(usp.ProductStatus) != productEnum.ProductStatusPutaway {
		//	list[0].InvalidProducts = []UserShopCartProductModel{usp}
		//} else {
		//	list[0].Products = []UserShopCartProductModel{usp}
		//}

		b := true
		for j := 0; j < len(list); j++ {
			if list[j].MerId == i.MerId {
				b = false
				break
			}
		}

		if b {
			list = append(list, UserShopCartMerModel{
				MerId:   i.MerId,
				MerName: m.Name,
				MerCode: m.Guid,
			})
		}

		for j := 0; j < len(list); j++ {
			if list[j].MerId == i.MerId {
				if productEnum.ProductStatus(usp.ProductStatus) != productEnum.ProductStatusPutaway || usp.Stock < 1 {
					list[j].InvalidProducts = append(list[j].InvalidProducts, usp)
				} else {
					list[j].Products = append(list[j].Products, usp)
				}
				break
			}

		}

	}

	return err, &list
}

//删除购物车产品
func DeleteShopCart(db *gorm.DB, q *DeleteShopCartModel) error {
	var (
		err error
	)

	err = db.Delete(model.ShopCart{}, "id in (?) and user_id=?", q.Ids, q.UserId).Error
	if err != nil {
		return err
	}

	return nil
}
