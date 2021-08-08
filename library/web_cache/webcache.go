package web_cache

import (
	"encoding/json"
	"fmt"
	"time"
)

type WebCacheModel struct {
	Cache          Cache
	DefaultTimeout time.Duration //默认过期时间
	//Key             string
	Enabled bool
}

//func (c *WebCacheModel) SetKey(key string) *WebCacheModel {
//
//	c.Key = key
//
//	return c
//}

func (c *WebCacheModel) GetData(key string, data interface{}) bool {

	if !c.Enabled {
		return false
	}

	if c.Cache.IsExist(key) {

		bytes := c.Cache.GetBytes(key)

		if bytes == nil {
			return false
		}

		if err := json.Unmarshal(bytes, data); err != nil {
			fmt.Printf("get cache Error ,key=%v err=%v", key, err.Error())
			return false
		}

		return true
	}
	return false
}

func (c *WebCacheModel) SetData(key string, data interface{}, timeout time.Duration) {
	var (
		err          error
		timeDuration time.Duration
	)
	if timeout > 0 {
		timeDuration = timeout
	} else {
		timeDuration = c.DefaultTimeout
	}

	if !c.Enabled {
		return
	}

	err = c.Cache.Set(key, &data, timeDuration)

	if err != nil {
		fmt.Printf("set cache Error ,key=%v err=%v", key, err.Error())
		return
	}

	return
}

func (c *WebCacheModel) Del(key string) {

	if !c.Enabled {
		return
	}

	err := c.Cache.Delete(key)

	if err != nil {
		fmt.Printf("set cache Error ,key=%v err=%v", key, err.Error())
		return
	}

	return
}

func (c *WebCacheModel) SetNX(key string, val interface{}, timeout time.Duration) bool {

	var (
		err          error
		b            bool
		timeDuration time.Duration
	)
	if timeout > 0 {
		timeDuration = timeout
	} else {
		timeDuration = c.DefaultTimeout
	}

	if !c.Enabled {
		fmt.Printf("Cache not Enabled")
		return false
	}

	err, b = c.Cache.SetNX(key, &val, timeDuration)

	if err != nil {
		fmt.Printf("set cache Error ,key=%v err=%v", key, err.Error())
		return false
	}

	return b
}
