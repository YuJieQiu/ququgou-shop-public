package web_cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

//Redis redis cache
type Redis struct {
	conn *redis.Client
}

//RedisOpts redis 连接属性
type RedisOpts struct {
	Host        string `yml:"host" json:"host"`
	Password    string `yml:"password" json:"password"`
	Database    int    `yml:"database" json:"database"`
	MaxIdle     int    `yml:"max_idle" json:"max_idle"`
	MaxActive   int    `yml:"max_active" json:"max_active"`
	IdleTimeout int32  `yml:"idle_timeout" json:"idle_timeout"` //second
}

//NewRedis 实例化
func NewRedis(opts *RedisOpts) *Redis {

	pool := redis.NewClient(&redis.Options{
		Addr:     opts.Host,
		Password: opts.Password, // no password set
		DB:       opts.Database, // use default DB

	})

	//pool := &redis.NewClient{
	//	MaxActive:   opts.MaxActive,
	//	MaxIdle:     opts.MaxIdle,
	//	IdleTimeout: time.Second * time.Duration(opts.IdleTimeout),
	//	Dial: func() (redis.Conn, error) {
	//		return redis.Dial("tcp", opts.Host,
	//			redis.DialDatabase(opts.Database),
	//			redis.DialPassword(opts.Password),
	//		)
	//	},
	//	TestOnBorrow: func(conn redis.Conn, t time.Time) error {
	//		if time.Since(t) < time.Minute {
	//			return nil
	//		}
	//		_, err := conn.Do("PING")
	//		return err
	//	},
	//}
	return &Redis{pool}
}

//SetConn 设置conn
func (r *Redis) SetConn(conn *redis.Client) {
	r.conn = conn
}

//Get 获取一个值
func (r *Redis) Get(key string) interface{} {
	conn := r.conn
	var data []byte
	var err error
	//if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
	//	return nil
	//}

	if data, err = conn.Get(key).Bytes(); err != nil {
		return nil
	}
	var reply interface{}

	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}

	return reply
}

//Get 获取一个值
func (r *Redis) GetBytes(key string) []byte {
	conn := r.conn
	//defer conn.Close()
	var err error
	var data []byte
	//var err error
	//if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
	//	return nil
	//}

	if data, err = conn.Get(key).Bytes(); err != nil {
		return nil
	}

	return data
}

//Set 设置一个值
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.conn
	//defer conn.Close()

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}

	err = conn.Set(key, data, timeout).Err()
	if err != nil {
		return
	}

	//_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)

	return
}

//IsExist 判断key是否存在
func (r *Redis) IsExist(key string) bool {
	conn := r.conn
	//defer conn.Close()

	i, err := conn.Exists(key).Result()
	if err != nil {
		return false
	}

	//a, _ := conn.Do("EXISTS", key)
	//i := a.(int64)
	if i > 0 {
		return true
	}
	return false
}

//Delete 删除
func (r *Redis) Delete(key string) error {
	conn := r.conn
	//defer conn.Close()

	//if _, err := conn.Do("DEL", key); err != nil {
	//	return err
	//}
	if err := conn.Del(key).Err(); err != nil {
		return err
	}

	return nil
}

//清除所有
func (r *Redis) FlushAll() error {
	//conn := r.conn.Get()
	//defer conn.Close()
	//
	//if _, err := conn.Do("FLUSHALL"); err != nil {
	//	return err
	//}

	return nil
}

//设置值 只有在 key 不存在时设置 key 的值 TODO:后面考虑 使用 https://github.com/go-redis/redis 开源库
func (r *Redis) SetNX(key string, val interface{}, timeout time.Duration) (error, bool) {
	conn := r.conn
	//defer conn.Close()

	var (
		data []byte
		err  error
		b    bool
	)
	if data, err = json.Marshal(val); err != nil {
		return err, b
	}

	b, err = conn.SetNX(key, data, timeout).Result()

	if err != nil {
		return err, b
	}

	return err, b
	//
	//if timeout == 0 {
	//	// Use old `SETNX` to support old Redis versions.
	//	t, err := conn.Do("SETNX", key, data)
	//	fmt.Println(t)
	//	fmt.Println(err)
	//
	//} else {
	//	if usePrecise(timeout) {
	//		_, err = conn.Do("SET", key, data, "PX", formatMs(timeout), "NX")
	//	} else {
	//		t, err := conn.Do("SET", key, data, "EX", formatSec(timeout), "NX")
	//		fmt.Println(t)
	//		fmt.Println(err)
	//	}
	//}

}

//判断秒或毫秒
func usePrecise(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

func formatMs(dur time.Duration) int64 {
	if dur > 0 && dur < time.Millisecond {
		fmt.Printf(
			"specified duration is %s, but minimal supported value is %s",
			dur, time.Millisecond,
		)
	}
	return int64(dur / time.Millisecond)
}

func formatSec(dur time.Duration) int64 {
	if dur > 0 && dur < time.Second {
		fmt.Printf(
			"specified duration is %s, but minimal supported value is %s",
			dur, time.Second,
		)
	}
	return int64(dur / time.Second)
}
