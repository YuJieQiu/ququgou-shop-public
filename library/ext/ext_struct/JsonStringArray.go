package ext_struct

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

//根据 string 中的 , 进行分割
type JsonStringArray string

// 实现 X 的json序列化方法
//MarshalJSON : 自定义类型转换到 json
func (arr JsonStringArray) MarshalJSON() ([]byte, error) {

	s := string(arr)

	ss := strings.Split(s, ",")

	urlsJson, _ := json.Marshal(ss)

	return []byte(urlsJson), nil
}

// UnmarshalJSON : UnmarshalJSON 自定义从json->自定义类型
func (arr *JsonStringArray) UnmarshalJSON(data []byte) error {

	s := strings.Replace(string(data), "\"", "", -1)

	*arr = JsonStringArray(s)

	return nil
}

//数据库类型
func (arr JsonStringArray) Value() (driver.Value, error) {
	s := string(arr)

	return s, nil
}

//数据库类型
func (arr *JsonStringArray) Scan(value interface{}) error {

	//driver.IsolationLevel()
	t := value.([]byte)

	*arr = JsonStringArray(string(t))

	return nil
}
