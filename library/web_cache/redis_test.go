package web_cache

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
	"time"
)

type UserName struct {
	Name string `json:"name"`
}

func TestSetNX(t *testing.T) {
	//opts := &RedisOpts{
	//	Host: "127.0.0.1:6379",
	//}
	//redis := NewRedis(opts)
	//var err error
	//timeoutDuration := 30 * time.Second

	//err = redis.SetNX("key6", "value9", timeoutDuration)
	//
	//if err != nil {
	//	t.Errorf("delete Error , err=%v", err)
	//}
	//
	//if !redis.IsExist("key3") {
	//	t.Error("IsExist Error")
	//}
	//
	//val := redis.Get("key3")
	//
	//fmt.Printf(val.(string))
}

func TestRedis(t *testing.T) {
	opts := &RedisOpts{
		Host: "127.0.0.1:6379",
	}
	redis := NewRedis(opts)
	var err error

	timeoutDuration := 1000 * time.Second

	var a = UserName{}
	a.Name = "test"

	var aa interface{}

	aa = a

	//body, err := json.Marshal(a)

	if err = redis.Set("username", aa, timeoutDuration); err != nil {
		t.Error("set Error", err)
	}

	if !redis.IsExist("username") {
		t.Error("IsExist Error")
	}

	var num float64 = 1.2345
	//ty := reflect.TypeOf(num)
	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))

	data := redis.GetBytes("username")

	if err = json.Unmarshal(data, &a); err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)

	//if name.Name != "silenceper" {
	//	t.Error("get Error")
	//}

	err = redis.Delete("username")
	if err != nil {
		t.Errorf("delete Error , err=%v", err)
	}
}
