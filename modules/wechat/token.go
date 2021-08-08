package wechat

import (
	"encoding/json"
	"errors"
	"github.com/ququgou-shop/library/web_cache"
	"net/http"
	"net/url"
	"time"
)

// 获取 access_token 成功返回数据
type response struct {
	BaseResponse
	AccessToken string        `json:"accessToken"`
	ExpireIn    time.Duration `json:"expireIn"`
}

// AccessToken 通过微信服务器获取 access_token
// 有效期 暂不返回
func GetAccessToken(appID, secret string, cache *web_cache.WebCacheModel, cacheKey string) (string, error) {
	var token string

	timeoutDuration := 72000 * time.Second //两小时

	url, err := url.Parse(BaseURL + TokenAPI)

	if err != nil {
		return "", err
	}

	//token 缓存
	if cache != nil && cacheKey != "" {
		if cache.GetData(cacheKey, token) {
			return token, nil
		}
	}

	query := url.Query()

	query.Set("appid", appID)
	query.Set("secret", secret)
	query.Set("grant_type", "client_credential")

	url.RawQuery = query.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", errors.New(WeChatServerError)
	}

	var data response

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.Errcode != 0 {
		return "", errors.New(data.Errmsg)
	}

	if cache != nil && cacheKey != "" && data.AccessToken != "" {
		//time.Second * data.ExpireIn,
		cache.SetData(cacheKey, data.AccessToken, timeoutDuration)
	}

	return data.AccessToken, nil
}
