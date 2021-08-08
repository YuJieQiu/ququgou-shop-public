package controller

type SearchProductMaterialModel struct {
	Text string `json:"text" form:"text"`
	Url  string `json:"url" form:"url"`
	Type int    `json:"type" form:"type"`
}
//
////搜索商品素材
//func SearchProductMaterial(c *gin.Context) {
//	var (
//		req SearchProductMaterialModel
//	)
//
//	if err := c.Bind(&req); err != nil {
//		//解析失败
//		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
//		return
//	}
//
//	//获取用户信息 后面会进行优化
//	err, _ := GetCustomUserInfo(c)
//	if err != nil {
//		JSON(c, http.StatusInternalServerError, "", err.Error())
//		return
//	}
//	//EEEEEEEE
//
//	if req.Type == 0 {
//		data := M1688.SearchProduct(req.Text)
//		JSON(c, http.StatusOK, "", data)
//		return
//	} else if req.Type == 1 {
//		data := M1688.SearchProductDetail(req.Url)
//		JSON(c, http.StatusOK, "", data)
//		return
//	}
//
//	//if err != nil {
//	//	JSON(c, http.StatusBadRequest, "", err.Error())
//	//	return
//	//}
//	JSON(c, http.StatusOK, "", nil)
//	return
//}
