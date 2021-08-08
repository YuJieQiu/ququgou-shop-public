package ext_struct

/**
	将JsonImage 对象 转化未string, 数据库查询时,将string 转为json 对象
**/

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JsonImage struct {
	Guid string `json:"guid"`
	Path string `json:"path"`
	Url  string `json:"url"`
	Type int    `json:"type" default:"1"` //预留字段
}

//根据 string 中的 , 进行分割
type JsonImageArrayString []JsonImage

// 实现 X 的json序列化方法
//MarshalJSON : 自定义类型转换到 json
func (current JsonImageArrayString) MarshalJSON() ([]byte, error) {

	var ttt []JsonImage
	//arr:=[]JsonImage{}

	if len(current) > 0 {
		ttt = []JsonImage(current)

		//for i:=0;i<len(ttt) ;i++ {
		//
		//}
	}

	urlsJson, err := json.Marshal(ttt)

	if err != nil {
		fmt.Println(err.Error())
	}

	//s:=string(this)

	//_ = json.Unmarshal([]byte(s), &arr)

	//urlsJson, _ := json.Marshal(s)

	return []byte(urlsJson), nil
}

// UnmarshalJSON : UnmarshalJSON 自定义从json->自定义类型
func (current *JsonImageArrayString) UnmarshalJSON(data []byte) error {

	var r []JsonImage

	//s:=strings.Replace(string(data), "\"", "", -1)

	_ = json.Unmarshal(data, &r)

	*current = JsonImageArrayString(r)

	return nil
}

//数据库类型
func (current JsonImageArrayString) Value() (driver.Value, error) {

	//var gatewayInfo = make(JsonImage, this)
	var ttt []JsonImage

	if len(current) > 0 {
		ttt = []JsonImage(current)
	}

	urlsJson, err := json.Marshal(ttt)

	if err != nil {
		fmt.Println(err.Error())
	}

	s := string(urlsJson)

	return s, nil
}

//数据库类型
func (current *JsonImageArrayString) Scan(value interface{}) error {

	var r []JsonImage

	//driver.IsolationLevel()
	t := value.([]byte)

	_ = json.Unmarshal(t, &r)

	*current = JsonImageArrayString(r)

	return nil
}

///////////////////////////////////////////////////////////////////////////////
//将string转未 JsonImage
type JsonImageString JsonImage

// 实现 X 的json序列化方法
//MarshalJSON : 自定义类型转换到 json
func (current JsonImageString) MarshalJSON() ([]byte, error) {

	var ttt JsonImage

	ttt = JsonImage(current)

	urlsJson, err := json.Marshal(ttt)

	if err != nil {
		fmt.Println(err.Error())
	}

	return []byte(urlsJson), nil
}

// UnmarshalJSON : UnmarshalJSON 自定义从json->自定义类型
func (current *JsonImageString) UnmarshalJSON(data []byte) error {

	var r JsonImage

	_ = json.Unmarshal(data, &r)

	*current = JsonImageString(r)

	return nil
}

//数据库类型
func (current JsonImageString) Value() (driver.Value, error) {

	var ttt JsonImage

	ttt = JsonImage(current)

	urlsJson, err := json.Marshal(ttt)

	if err != nil {
		fmt.Println(err.Error())
	}

	s := string(urlsJson)

	return s, nil
}

//数据库类型
func (current *JsonImageString) Scan(value interface{}) error {

	var r JsonImage

	t := value.([]byte)

	_ = json.Unmarshal(t, &r)

	*current = JsonImageString(r)

	return nil
}
