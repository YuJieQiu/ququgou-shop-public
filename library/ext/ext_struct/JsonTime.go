package ext_struct

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type JsonTime time.Time

// 实现 JsonTime 的json序列化方法
//MarshalJSON : 自定义类型转换到 json
func (j JsonTime) MarshalJSON() ([]byte, error) {
	layout := "2006-01-02 15:04:05" //"2006-01-02T15:04:05.000Z"
	//time.Time(this).Format(layout)
	loc, _ := time.LoadLocation("Local")

	//dt, err := time.ParseInLocation(layout, time.Time(this), loc)
	//In(loc)

	ts := time.Time(j).In(loc)

	var stamp string

	if ts.IsZero() {
		stamp = fmt.Sprintf("\"%s\"", "")
		return []byte(stamp), nil
	} else {
		stamp = fmt.Sprintf("\"%s\"", ts.Format(layout))
		return []byte(stamp), nil
	}

}

// UnmarshalJSON : UnmarshalJSON 自定义从json->自定义类型
func (j *JsonTime) UnmarshalJSON(data []byte) error {
	layout := "2006-01-02 15:04:05"
	//layout := "2006-01-02T15:04:05.000Z"
	//2018-12-29T12:46:17.000Z
	//2018-12-25T16:29:01.000Z
	// loc, _ := time.LoadLocation("Local")

	str := strings.Replace(string(data), "\"", "", -1)
	//"2014-11-12T11:45:26.371Z"

	if str == "" {
		j = nil
		return nil
	}

	//millis, err :=time.ParseInLocation(layout,str,loc)
	millis, err := time.Parse(layout, str)

	if err != nil {
		return err
	}

	if millis.IsZero() {
		j = nil
	} else {
		*j = JsonTime(millis)
	}

	return nil
}

func (j JsonTime) ConverTime() time.Time {

	data := time.Time(j)

	layout := "2006-01-02T15:04:05.000Z"

	loc, _ := time.LoadLocation("Local")

	str := strings.Replace(data.Format(layout), "\"", "", -1)

	millis, _ := time.ParseInLocation(layout, str, loc)

	return millis
}

func (j JsonTime) Value() (driver.Value, error) {
	st := time.Time(j)

	if st.IsZero() {
		return nil, nil
	}

	return time.Time(j), nil
}

func (j *JsonTime) Scan(value interface{}) error {

	//driver.IsolationLevel()
	st := value.(time.Time)

	if st.IsZero() {
		j = nil
		return nil
	}

	*j = JsonTime(st)

	return nil
}
