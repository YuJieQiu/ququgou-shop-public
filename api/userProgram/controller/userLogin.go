package controller

import (
	"fmt"

	"github.com/ququgou-shop/library/utils"

	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/userProgram/config"
	jwt "github.com/ququgou-shop/api/userProgram/middleware"
	"github.com/ququgou-shop/modules/user"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/modules/wechat"

	"time"

	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	AvatarURL     string `form:"avatarUrl" json:"avatarUrl"`
	City          string `form:"city" json:"city"`
	Code          string `form:"code" json:"code"`
	Country       string `form:"country" json:"country"`
	Gender        byte   `form:"gender" json:"gender"`
	Language      string `form:"language" json:"language"`
	NickName      string `form:"nickName" json:"nickName"`
	Province      string `form:"province" json:"province"`
	EncryptedData string `form:"encryptedData" json:"encryptedData"`
	Signature     string `form:"signature" json:"signature"`
	Iv            string `form:"iv" json:"iv"`
	RawData       string `form:"rawData" json:"rawData"`
}

var userModule = user.ModuleUser{DB: getConnDB()}

// WeChatLogin 微信登陆
// @Summary 微信登陆
// @Description 微信登陆
// @Accept json
// @Produce json
// @Param body body LoginRequest true "body参数"
// @Success 200 {object} Response{data=model.User} "ok" "返回信息"
// @Failure 400 {object} Response "错误"
// @Failure 401 {object} Response "错误"
// @Failure 500 {object} Response "错误"
// @Router /user/person/login [post]
func WeChatLogin(c *gin.Context) {
	var (
		loginReq  LoginRequest
		modelUser *model.User
		err       error
	)

	if c.BindJSON(&loginReq) != nil {
		//解析失败
		JSON(c, http.StatusBadRequest, "参数解析失败", nil)
		return
	}

	data, err := wechat.Login(config.Config.WX.AppID, config.Config.WX.AppSecret, loginReq.Code)

	if err != nil {
		JSON(c, http.StatusBadRequest, "wx res error", err)
		return
	}

	q := &user.GetSingleUserModel{
		WechatOpenId: data.OpenID,
	}

	err, modelUser = userModule.GetSingleUser(q)

	if modelUser.ID == 0 {
		modelUser.UserName = loginReq.NickName
		modelUser.Gender = loginReq.Gender
		modelUser.AvatarUrl = loginReq.AvatarURL
		modelUser.WechatOpenId = data.OpenID
		modelUser.Guid = utils.CreateUUID()
		modelUser.Status = 1
		modelUser.Type = 1 //商城小程序

		err, modelUser = userModule.CreateUser(modelUser)

		if err != nil {
			JSON(c, http.StatusBadRequest, "create user", err)
			return
		}
	} else {
		//update user info
		modelUser.UserName = loginReq.NickName
		modelUser.Gender = loginReq.Gender
		modelUser.AvatarUrl = loginReq.AvatarURL

		err, modelUser = userModule.EditUser(modelUser)
		if err != nil {
			println(err)
		}
	}

	u := jwt.CustomClaims{
		Guid:     modelUser.Guid,
		UserInfo: modelUser,
		StandardClaims: jwtgo.StandardClaims{
			ExpiresAt: int64(time.Now().Unix() + 43200), // 过期时间 /s
		},
	}

	t, err := jwt.NewJWT().CreateToken(u)

	if err != nil {
		JSON(c, http.StatusBadRequest, "token error", err)
		return
	}

	//get ip address S
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")
	ipAddress := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ipAddress == "" {
		ipAddress = strings.TrimSpace(c.Request.Header.Get("X-Real-Ip"))
		if ipAddress == "" {
			ipAddress, _, _ = net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
		}
	}
	//get ip address E

	logErr, _ := userModule.CreateUserLoginLog(&model.UserLoginLog{
		UserId: modelUser.ID,
		Ip:     ipAddress,
		Remark: "WechatLogin",
	})
	if logErr != nil {
		fmt.Println(logErr.Error())
	}

	//data.LoginResponse
	JSON(c, http.StatusOK, "", t)
	return
}
