package web_cache

import "time"

type Cache interface {
	Get(key string) interface{}
	GetBytes(key string) []byte
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
	FlushAll() error
	SetNX(key string, val interface{}, timeout time.Duration) (error, bool)
}
