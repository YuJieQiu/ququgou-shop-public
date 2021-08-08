package cache

import (
	"github.com/ququgou-shop/api/merProgram/config"
	"github.com/ququgou-shop/library/web_cache"
	"time"
)

var WebCache web_cache.WebCacheModel

func init() {

	opts := &web_cache.RedisOpts{
		Host: config.Config.Cache.Redis.Host,
	}
	redis := web_cache.NewRedis(opts)
	WebCache.Cache = redis

	WebCache.DefaultTimeout = 3600 * time.Second //一小时
	WebCache.Enabled = config.Config.Cache.Enabled
}
