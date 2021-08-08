package wechat

import "net/url"

const (
	// BaseURL 微信请求基础URL
	BaseURL           = "https://api.weixin.qq.com"
	TokenAPI          = "/cgi-bin/token"
	CodeAPI           = "/sns/jscode2session"
	WeChatServerError = "微信服务器发生错误"
)

// BaseResponse 请求微信返回基础数据
type BaseResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// TokenAPI 获取带 token 参数 的 API 地址
func GetTokenAPI(api, token string) (string, error) {
	u, err := url.Parse(api)
	if err != nil {
		return "", err
	}
	query := u.Query()
	query.Set("access_token", token)
	u.RawQuery = query.Encode()

	return u.String(), nil
}
