package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/service/merchantService"
	"net/http"
)

//获取商户信息
func GetMerchantInfo(c *gin.Context) {
	var (
		req merchantService.GetMerchantInfoModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	//EEEEEEEE

	req.IsWhereUserId = true
	req.MerId = u.MerInfo.MerId
	req.MerCode = u.MerInfo.MerGuid

	err, data := merchantService.GetMerchantInfo(getConnDB(),&req, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//更新商户信息
func UpdateMerchantInfo(c *gin.Context) {
	var (
		req merchantService.MerInfoModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err)
		return
	}
	//EEEEEEEE

	req.Id = u.MerInfo.MerGuid

	err, data := merchantService.UpdateMerchantInfo(getConnDB(),&req, u.UserInfo)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

