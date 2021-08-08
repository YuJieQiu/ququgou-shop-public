package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/modules/label/model"
	"github.com/ququgou-shop/service/merchantService"
	"net/http"
)

//获取标签信息
func GetLabelList(c *gin.Context) {
	var (
		req merchantService.GetLabelListModel
	)

	if err := c.Bind(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	//err, u := GetCustomUserInfo(c)
	//if err != nil {
	//	JSON(c, http.StatusInternalServerError, "", err.Error())
	//	return
	//}
	//EEEEEEEE

	req.Type = 1

	err, data, _ := merchantService.GetLabelList(getConnDB(), &req, false)

	if err != nil {
		JSON(c, http.StatusBadRequest, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "", data)
	return
}

//创建标签
func CreateLabel(c *gin.Context) {
	var (
		req model.Label
	)

	if err := c.BindJSON(&req); err != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", err.Error())
		return
	}

	//获取用户信息 后面会进行优化
	//err, u := GetCustomUserInfo(c)
	//if err != nil {
	//	JSON(c, http.StatusInternalServerError, "", err.Error())
	//	return
	//}
	//EEEEEEEE
	req.Type = 1

	err, w := merchantService.GetLabel(getConnDB(), &merchantService.GetLabelModel{
		Type: req.Type,
		Text: req.Text,
	})
	if err == nil && w != nil && w.ID > 0 {
		JSON(c, http.StatusOK, "ok", w)
		return
	}

	err, data := merchantService.CreateLabel(getConnDB(), &req)

	if err != nil {
		JSON(c, http.StatusInternalServerError, "", err.Error())
		return
	}

	JSON(c, http.StatusOK, "ok", data)
	return
}
