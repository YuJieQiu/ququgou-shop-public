package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ququgou-shop/api/merProgram/config"
	jwt "github.com/ququgou-shop/api/merProgram/middleware"
	"github.com/ququgou-shop/library/utils"
	"github.com/ququgou-shop/modules/user"
	"github.com/ququgou-shop/modules/user/model"
	"github.com/ququgou-shop/modules/wechat"

	"time"

	"net/http"

	jwtgo "github.com/dgrijalva/jwt-go"
)

//微信登录
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


var userModule =user.ModuleUser{DB:getConnDB()}

//小程序登
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

	fmt.Println(loginReq)

	data, err := wechat.Login(config.Config.WX.AppID, config.Config.WX.AppSecret, loginReq.Code)

	if err != nil {
		JSON(c, http.StatusBadRequest, "wx res error", err.Error())
		return
	}

	q := &user.GetSingleUserModel{
		WechatOpenId: data.OpenID,
		Type:         3,
	}

	err, modelUser = userModule.GetSingleUser(q)

	if modelUser.ID == 0 {
		modelUser.UserName = loginReq.NickName
		modelUser.Gender = loginReq.Gender
		modelUser.AvatarUrl = loginReq.AvatarURL
		modelUser.WechatOpenId = data.OpenID
		modelUser.Guid = utils.CreateUUID()
		modelUser.Status = 1
		modelUser.Type = 3 //商户小程序

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

	//set cache TODO: 有问题，第一次登里的时候缓存无 则授权不成功
	//utils.Caches.Set(utils.Token+u.Guid, t, 5*time.Minute)

	//data.LoginResponse
	JSON(c, http.StatusOK, "", t)
	return
}
