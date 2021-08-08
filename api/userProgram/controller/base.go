package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/api/userProgram/cache"
	"github.com/ququgou-shop/api/userProgram/db"
	jwt "github.com/ququgou-shop/api/userProgram/middleware"
	"github.com/ququgou-shop/modules/user/model"
)

func GetCustomUserInfo(c *gin.Context) (error, *model.User) {

	var modelUser model.User

	v, _ := c.Get("claims")

	if v == nil {
		return errors.New("获取用户信息失败"), nil
	}

	va := v.(*jwt.CustomClaims)

	if va == nil {
		return errors.New("获取用户信息失败"), nil
	}

	if bol := cache.WebCache.GetData(cache.UserInfo+va.Guid, &modelUser); bol {
		return nil, &modelUser
	}

	//通过openid 获取到 用户信息
	db.MysqlConn().Where("guid = ? and type=?", va.Guid, 1).Find(&modelUser)

	if modelUser.ID == 0 {
		//获取用户信息失败
		return errors.New("获取用户信息失败"), nil
	}

	go cache.WebCache.SetData(cache.UserInfo+va.Guid, modelUser, 0)

	return nil, &modelUser
}

func getConnDB() *gorm.DB {
	return db.MysqlConn()
}

type errorCode struct {
	SUCCESS      int
	ERROR        int
	AuthErr      int
	RoleErr      int
	NotFound     int
	LoginError   int
	LoginTimeout int
	InActive     int
}

// ErrorCode 错误码
var ErrorCode = errorCode{
	SUCCESS:      0,
	ERROR:        1,
	AuthErr:      40029, // 认证失败，请重新登陆
	RoleErr:      1406,  // 权限不够
	NotFound:     404,
	LoginError:   1000, //用户名或密码错误
	LoginTimeout: 1001, //登录超时
	InActive:     1002, //未激活账号
}

type errorMessage struct {
	SUCCESS      string
	ERROR        string
	AuthErr      string
	RoleErr      string
	NotFound     string
	LoginError   string
	LoginTimeout string
	InActive     string
}

// ErrorMessage 错误码
var ErrorMessage = errorMessage{
	SUCCESS:      "成功",
	ERROR:        "错误",
	AuthErr:      "认证失败,请重新登陆",
	RoleErr:      "权限不够", // 权限不够
	NotFound:     "未找到",
	LoginError:   "1000", //用户名或密码错误
	LoginTimeout: "1001", //登录超时
	InActive:     "1002", //未激活账号
}

// Response : JSON Response Object
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Response : JSON Response Object
type ResponsePages struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    Page        `json:"page,omitempty"`
	Te      string      `json:"test,omitempty"`
}
type Page struct {
	Total int `json:"total"`
}

// JSONSuccess ...
func JSON(c *gin.Context, code int, message string, d interface{}) {

	c.JSON(http.StatusOK, &Response{
		Code:    code,
		Message: message,
		Data:    d,
	})

}

// JSONSuccess ...
func JSONSuccess(c gin.Context, d interface{}) {

	c.JSON(http.StatusOK, &Response{
		Code:    ErrorCode.SUCCESS,
		Message: "success",
		Data:    d,
	})

}

// JSONSuccess ...
func JSONPage(c *gin.Context, code int, message string, d interface{}, Total int) {

	c.JSON(http.StatusOK, &ResponsePages{
		Code:    code,
		Message: message,
		Data:    d,
		Page: Page{
			Total: Total,
		},
	})

}
