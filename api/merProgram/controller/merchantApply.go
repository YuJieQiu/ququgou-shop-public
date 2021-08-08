package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/service/merchantService"
)

func CreateMerApplyInfo(c *gin.Context) {
	var (
		req merchantService.CreateMerApplyInfoModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	//EEEEEEEE

	req.UserId = u.ID

	err = merchantService.CreateMerApplyInfo(getConnDB(), &req, u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}

func GetMerApplyInfo(c *gin.Context) {

	//获取用户信息 后面会进行优化
	err, u := GetUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}
	//EEEEEEEE

	err, data := merchantService.GetMerApplyInfo(getConnDB(), u, config.Config.ImgService.QiniuUrl)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

type AutoMerApplyVerifiedModel struct {
	Id uint64 `json:"id" form:"id"`
}

func AutoMerApplyVerified(c *gin.Context) {
	var (
		req AutoMerApplyVerifiedModel
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	err := merchantService.MerApplyVerified(getConnDB(), req.Id)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", "")
	return
}
