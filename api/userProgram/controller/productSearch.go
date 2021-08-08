package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	"github.com/ququgou-shop/api/userProgram/elasticClientBase"
	"github.com/ququgou-shop/library/ext/ext_struct"
	"github.com/ququgou-shop/modules/product/model"
	"github.com/ququgou-shop/service/elasticsearchService"
	"github.com/ququgou-shop/service/merchantService"
	"github.com/ququgou-shop/service/productService"
	"github.com/ququgou-shop/service/searchService"
)

// TestProductSync 商品同步
// @Summary 搜索产品同步
// @Description 描述
// @Accept json
// @Produce json
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /search/sync [get]
func TestProductSync(c *gin.Context) {
	//TODO: 商品同步 ES 、获取ES等
	var ps []model.Product

	err := getConnDB().Find(&ps).Error
	if err != nil {
		JSON(c, http.StatusBadRequest, "查询失败", err.Error())
	}

	var mers map[uint64]*merchantService.MerInfoModel
	mers = make(map[uint64]*merchantService.MerInfoModel)

	for _, i := range ps {
		var mer *merchantService.MerInfoModel
		if len(mers) > 0 {
			if _, ok := mers[i.MerId]; ok {
				mer = mers[i.MerId]
			}
		}

		if mer == nil {
			err, mer = merchantService.GetMerchantInfo(getConnDB(), &merchantService.GetMerchantInfoModel{
				MerId: i.MerId,
			}, config.Config.ImgService.QiniuUrl)

			if err != nil {
				JSON(c, http.StatusBadRequest, err.Error(), err.Error())
				return
			}
			mers[i.ID] = mer
		}

		err, res := productService.GetProductInfo(getConnDB(), &productService.GetProductInfoModel{
			ProductId: i.ID,
			MerId:     i.MerId,
		}, config.Config.ImgService.QiniuUrl)

		if err != nil {
			JSON(c, http.StatusBadRequest, err.Error(), err.Error())
			return
		}

		var tag []string

		if len(res.Tags) > 0 {
			for _, pc := range res.Tags {
				tag = append(tag, pc.Tag.Name)
			}
		}

		tagStr, _ := json.Marshal(tag)

		var productCategory []string

		if res.CategoryProduct != nil && len(res.CategoryProduct) > 0 {
			for _, pc := range res.CategoryProduct {
				productCategory = append(productCategory, pc.Category.Name)
			}
		}

		productCategoryStr, _ := json.Marshal(productCategory)

		d := elasticsearchService.ElasticProductModel{
			ID:            i.ID,
			Guid:          i.Guid,
			MerId:         i.ID,
			TypeId:        i.TypeId,
			BrandId:       i.BrandId,
			BrandName:     "", //TODO:暂时没有品牌
			Name:          i.Name,
			Status:        i.Status,
			Description:   i.Description,
			Keywords:      string(i.Keywords),
			Tags:          string(tagStr),
			OriginalPrice: res.OriginalPrice,
			MinPrice:      res.MinPrice,
			MaxPrice:      res.MaxPrice,
			CurrentPrice:  res.CurrentPrice,
			Sales:         res.Sales,
			ProductType:   res.ProductType,
			Integral:      res.Integral,
			Active:        res.Active,
			City:          mer.Address.City,
			Region:        mer.Address.Region,
			Town:          mer.Address.Town,
			Address:       mer.Address.Address,
			Location:      elasticsearchService.LocationModel{Lat: mer.Address.Latitude, Lon: mer.Address.Longitude},
			Priority:      0,
			IsSingle:      res.IsSingle,
			CategoryIds:   res.CategoryIds,
			CategoryInfo:  string(productCategoryStr),
			CreatedTime:   ext_struct.JsonTime(res.CreatedAt),
			UpdatedTime:   ext_struct.JsonTime(res.UpdatedAt),
		}

		//d.ImageJsonModel = res.ImageJsonModel
		//var skus []domain.ElasticProductSKUModel
		//for _, s := range res.SKU {
		//	arrStr, _ := json.Marshal(s.AttributeInfo)
		//	skus = append(skus, domain.ElasticProductSKUModel{
		//		ID:                s.ID,
		//		Guid:              s.Guid,
		//		Name:              s.Name,
		//		OriginalPrice:     s.OriginalPrice,
		//		Sales:             s.Sales,
		//		Status:            s.Status,
		//		Sort:              s.Sort,
		//		IsSingleAttribute: s.IsSingleAttribute,
		//		AttributeInfo:     string(arrStr),
		//	})
		//}

		//d.SkuInfo = skus

		err, _ = elasticsearchService.AddProduct(elasticClientBase.ElasticClient, &d)
		if err != nil {
			JSON(c, http.StatusBadRequest, err.Error(), err.Error())
			return
		}
	}
	//data := elasticsearchService.GetProduct(elasticClientBase.ElasticClient, req.Id)
	JSON(c, http.StatusOK, "", nil)
	return
}

type TestProductGetModel struct {
	Id string `json:"id" form:"id"`
}

// TestProductGet ES搜索产品获取
// @Summary ES搜索产品获取
// @Description ES搜索产品获取
// @Accept json
// @Produce json
// @Param body body TestProductGetModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /search/get [get]
func TestProductGet(c *gin.Context) {
	//TODO: ES搜索产品获取
	var (
		req TestProductGetModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	err, data := elasticsearchService.GetProduct(elasticClientBase.ElasticClient, req.Id)

	if err != nil {
		JSON(c, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

type TestProductQueryModel struct {
	Name string `json:"name" form:"name"`
	Text string `json:"text" form:"text"`
}

// TestProductQuery  ES产品查询
// @Summary ES产品查询
// @Description ES产品查询
// @Accept json
// @Produce json
// @Param body body TestProductQueryModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /search/query [get]
func TestProductQuery(c *gin.Context) {

	var (
		req TestProductQueryModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	q := &elasticsearchService.ProductQueryModel{
		QueryFields: req.Name,
		QueryText:   req.Text,
		QueryType:   5,
		From:        0,
		Size:        100,
	}
	err, data := elasticsearchService.QueryProduct(elasticClientBase.ElasticClient, q)
	if err != nil {
		JSON(c, http.StatusBadRequest, err.Error(), err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// SearchProductList 获取搜索产品列表
// @Summary 简单描述
// @Description 描述
// @Accept json
// @Produce json
// @Param body body productService.GetSearchProductListModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /product/search [post]
func SearchProductList(c *gin.Context) {
	//TODO：SearchProductList 获取搜索产品列表
	var (
		req productService.GetSearchProductListModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	////获取用户信息 后面会进行优化
	//err, u := GetCustomUserInfo(c)
	//if err != nil {
	//	JSON(c, http.StatusInternalServerError, "", err.Error())
	//	return
	//}

	err, data, count := searchService.GetSearchProductList(getConnDB(), &req, nil, false, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	JSONPage(c, http.StatusOK, "", data, count)
	return
}

type GetProductSearchPageConfigDataModel struct {
	CategoryId uint64  `json:"categoryId" form:"categoryId"` //分类Id
	SearchKey  string  `json:"searchKey" form:"searchKey"`   //搜索key
	Lat        float64 `form:"lat" json:"lat"`               //维度
	Lon        float64 `form:"lon" json:"lon"`               //经度
}

type ProductSearchPageConfigDataModel struct {
	AttrArr []AttrArrModel `json:"attrArr"`
}

type AttrArrModel struct {
	Id                  uint64         `json:"id"`
	AttrSelectItemIndex int            `json:"attrSelectItemIndex"`
	Name                string         `json:"name"`
	Selected            bool           `json:"selected"`
	SelectedName        string         `json:"selectedName"`
	IsActive            bool           `json:"isActive"`
	IsItemActive        bool           `json:"isItemActive"`
	List                []AttrArrModel `json:"list"`
}

// GetProductSearchPageConfigData 获取搜索页面配置信息
// @Summary 简单描述
// @Description 描述
// @Accept json
// @Produce json
// @Param body body GetProductSearchPageConfigDataModel true "body参数"
// @Success 200 {object} Response{data=ProductSearchPageConfigDataModel} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router/product/list/config [get]
func GetProductSearchPageConfigData(c *gin.Context) {
	//TODO: GetProductSearchPageConfigData 获取搜索页面配置信息
	var (
		err error
		req GetProductSearchPageConfigDataModel
		res ProductSearchPageConfigDataModel
		arr []AttrArrModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	cq := productService.GetCategoryListModel{
		IsSystem: true,
	}

	err, cList, _ := productService.GetCategoryList(getConnDB(), &cq, false, config.Config.ImgService.QiniuUrl)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}

	arr = append(arr, AttrArrModel{
		Id:                  0,
		AttrSelectItemIndex: -1,
		Name:                "新品",
		SelectedName:        "新品",
		Selected:            false,
		IsActive:            false,
	})

	var cArr []AttrArrModel
	var AttrSelectItemIndex int = -1
	for index, i := range cList {
		attrArrItem := AttrArrModel{
			Id:           i.ID,
			Name:         i.Name,
			SelectedName: i.Name,
			Selected:     false,
			IsActive:     false,
		}

		if attrArrItem.Id == req.CategoryId {
			attrArrItem.IsActive = true
			attrArrItem.Selected = true
			AttrSelectItemIndex = index
		}

		cArr = append(cArr, attrArrItem)
		if len(i.Child) > 0 {
			for cIndex, c := range i.Child {
				attrArrItem := AttrArrModel{
					Id:           c.ID,
					Name:         c.Name,
					SelectedName: c.Name,
					Selected:     false,
					IsActive:     false,
				}
				if attrArrItem.Id == req.CategoryId {
					attrArrItem.IsActive = true
					attrArrItem.Selected = true
					AttrSelectItemIndex = cIndex
				}

				cArr = append(cArr, attrArrItem)
			}
		}
	}

	arr = append(arr, AttrArrModel{
		Id:                  0,
		AttrSelectItemIndex: AttrSelectItemIndex,
		Name:                "分类",
		SelectedName:        "分类",
		Selected:            req.CategoryId > 0,
		IsActive:            req.CategoryId > 0,
		List:                cArr,
	})

	res.AttrArr = arr

	JSON(c, http.StatusOK, "", res)
	return
}
