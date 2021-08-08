package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ququgou-shop/api/merProgram/cache"
	"github.com/ququgou-shop/api/merProgram/db"
	jwt "github.com/ququgou-shop/api/merProgram/middleware"
	"github.com/ququgou-shop/modules/merchant"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/service/merchantService"

	"net/http"
)

func GetUserInfo(c *gin.Context) (error, *model.User) {
	v, _ := c.Get("claims")

	if v == nil {
		return errors.New("获取用户信息失败"), nil
	}
	va := v.(*jwt.CustomClaims)

	var modelUser model.User
	//通过缓存获取
	if bol := cache.WebCache.GetData(cache.User+va.Guid, &modelUser); !bol {
		db.MysqlConn().Where("guid =? ", va.Guid).First(&modelUser)
		//存入缓存
		go cache.WebCache.SetData(cache.User+va.Guid, modelUser, 0)
	}

	return nil, &modelUser
}

func getConnDB() *gorm.DB  {
	return db.MysqlConn()
}

var merchantModule=merchant.ModuleMerchant{
	DB:getConnDB(),
}

//获取当前用户基本信息
func GetCustomUserInfo(c *gin.Context) (error, *CustomUserInfo) {
	var (
		data      CustomUserInfo
		modelUser model.User
	)

	v, _ := c.Get("claims")

	if v == nil {
		return errors.New("获取用户信息失败"), nil
	}
	va := v.(*jwt.CustomClaims)

	//通过缓存获取
	if bol := cache.WebCache.GetData(cache.UserInfo+va.Guid, &data); bol {
		return nil, &data
	}

	// 获取 用户信息
	err := db.MysqlConn().Where("guid = ? and type=?", va.Guid, 3).Find(&modelUser).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("获取用户信息失败"), nil
		}
		return err, nil
	}

	if modelUser.ID == 0 {
		//获取用户信息失败
		return errors.New("获取用户信息失败"), nil
	}


	//获取商户信息
	err, merData := merchantService.GetMerchantForUser(getConnDB(),modelUser.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("获取商户信息失败"), nil
		}
		return err, nil
	}

	if merData.MerId == 0 {
		//获取商户信息失败
		return errors.New("获取商户信息失败"), nil
	}

	data = CustomUserInfo{
		UserInfo: &modelUser,
		MerInfo:  merData,
	}

	//存入缓存
	go cache.WebCache.SetData(cache.UserInfo+va.Guid, data, 0)

	return nil, &data
}

type CustomUserInfo struct {
	UserInfo *model.User
	MerInfo  *merchantService.GetMerchantForUserModel
}

//
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
