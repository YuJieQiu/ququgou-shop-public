package wechat

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// LoginResponse 返回给用户的数据
type LoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"sessionKey"`
	// 用户在开放平台的唯一标识符
	// 只在满足一定条件的情况下返回
	UnionID string `json:"unionid"`
	BaseResponse
}

//小程序登
// @appID 小程序 appID
// @secret 小程序的 app secret
// @code 小程序登录时获取的 code
func Login(appID, secret, code string) (*LoginResponse, error) {

	if code == "" {
		return nil, errors.New("code can not be null")
	}

	if appID == "" {
		return nil, errors.New("appID can not be null")
	}

	if secret == "" {
		return nil, errors.New("secret can not be null")
	}

	api, err := code2url(appID, secret, code)

	if err != nil {
		return nil, err
	}

	res, err := http.Get(api)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("WeChatServerError = 微信服务器发生错")
	}

	var data LoginResponse

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Errcode != 0 {
		return nil, errors.New(data.Errmsg)
	}

	if data.OpenID == "" {
		return nil, errors.New("OpenID Empty")
	}

	return &data, nil
}

// 拼接 获取 session_key 的 URL
func code2url(appID, secret, code string) (string, error) {

	url, err := url.Parse(BaseURL + CodeAPI)
	if err != nil {
		return "", err
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("js_code", code)
	query.Set("grant_type", "authorization_code")

	url.RawQuery = query.Encode()

	return url.String(), nil
}
