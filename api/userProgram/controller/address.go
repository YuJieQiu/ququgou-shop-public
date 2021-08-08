package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/addressService"

	"net/http"
)

// CreateAddress 创建用户收获地址
// @Summary 创建用户收获地址
// @Description 创建用户收获地址
// @Accept json
// @Produce json
// @Param body body model.Address true "body参数"
// @Success 200 {object} Response{data=model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/create [post]
// @Security ApiKeyAuth
func CreateAddress(c *gin.Context) {

	var (
		req model.Address
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	req.UserId = u.ID

	err, data := addressService.CreateAddress(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// SaveAddress 保存用户收获地址
// @Summary 保存用户收获地址
// @Description 保存用户收获地址
// @Accept json
// @Produce json
// @Param body body model.Address true "body参数"
// @Success 200 {object} Response{data=model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/save [post]
func SaveAddress(c *gin.Context) {

	var (
		req model.Address
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	req.UserId = u.ID

	err, data := addressService.SaveAddress(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// DeleteAddress 删除用户收获地址
// @Summary 删除用户收获地址
// @Description 删除用户收获地址
// @Accept json
// @Produce json
// @Param body body addressService.DeleteAddressModel true "body参数"
// @Success 200 {object} Response "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/user/delete [post]
func DeleteAddress(c *gin.Context) {
	var (
		req addressService.DeleteAddressModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	req.UserId = u.ID

	err = addressService.DeleteAddress(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", nil)
	return
}

// GetAddress 获取用户收获地址信息
// @Summary 获取用户收获地址信息
// @Description 获取用户收获地址信息
// @Accept json
// @Produce json
// @Param body body addressService.GetAddressModel true "body参数"
// @Success 200 {object} Response{data=model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /user/person/login [post]
func GetAddress(c *gin.Context) {
	var (
		req addressService.GetAddressModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err)
		return
	}

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	req.UserId = u.ID

	err, data := addressService.GetAddress(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetAddressList 获取全部地址列表
// @Summary 获取全部地址列表
// @Description 获取全部地址列表
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/getList [get]
func GetAddressList(c *gin.Context) {

	err, data := addressService.GetAddressList(getConnDB())

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetUserAddressList 获取用户收获地址列表
// @Summary 获取用户收获地址列表
// @Description 获取用户收获地址列表
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=[]model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/user/getList [get]3
func GetUserAddressList(c *gin.Context) {

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	err, data := addressService.GetUserAddressList(getConnDB(), u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

// GetUserAddressFirst 获取用户收获地址(默认第一个)
// @Summary 获取用户收获地址(默认第一个)
// @Description 获取用户收获地址(默认第一个)
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=model.Address} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /address/user/first [post]
func GetUserAddressFirst(c *gin.Context) {

	//获取用户信息 后面会进行优化
	err, u := GetCustomUserInfo(c)
	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	err, data := addressService.GetUserAddressFirst(getConnDB(), u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err)
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}
