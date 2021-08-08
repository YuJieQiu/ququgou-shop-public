package shop_ext_struct

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type (
	SkuAttributeValuesArrayModel []SkuAttributeValuesModel

	SkuAttributeValuesModel struct {
		Aid       uint64 `json:"aid"`
		AttName   string `json:"attName"`
		Vid       uint64 `json:"vid"`
		ValueName string `json:"valueName"`
	}
)

// 实现 X 的json序列化方法
//MarshalJSON : 自定义类型转换到 json
func (model SkuAttributeValuesArrayModel) MarshalJSON() ([]byte, error) {

	var ttt []SkuAttributeValuesModel

	if model != nil && len(model) > 0 {
		ttt = []SkuAttributeValuesModel(model)
	}

	urlsJson, err := json.Marshal(ttt)

	if err != nil {
		fmt.Println(err.Error())
	}

	return []byte(urlsJson), nil
}

// UnmarshalJSON : UnmarshalJSON 自定义从json->自定义类型
func (model *SkuAttributeValuesArrayModel) UnmarshalJSON(data []byte) error {

	var r []SkuAttributeValuesModel

	_ = json.Unmarshal(data, &r)

	*model = SkuAttributeValuesArrayModel(r)

	return nil
}

//数据库类型
func (model SkuAttributeValuesArrayModel) Value() (driver.Value, error) {

	var arr []SkuAttributeValuesModel

	if len(model) > 0 {
		arr = []SkuAttributeValuesModel(model)
	} else {
		return nil, nil
	}

	urlsJson, err := json.Marshal(arr)

	if err != nil {
		fmt.Println(err.Error())
	}

	s := string(urlsJson)

	return s, nil
}

//数据库类型
func (model *SkuAttributeValuesArrayModel) Scan(value interface{}) error {

	var data []SkuAttributeValuesModel

	//driver.IsolationLevel()
	arr := value.([]byte)

	if arr == nil || len(arr) <= 0 {
		return nil
	}
	_ = json.Unmarshal(arr, &data)

	*model = SkuAttributeValuesArrayModel(data)

	return nil
}
