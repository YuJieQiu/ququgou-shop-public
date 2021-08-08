package common

import (
	"strconv"
	"time"
)

//生成交易流水号
func CreateTradeNo(userGuid string) string {
	// x - xxxxxx - xxxxx -xxx -xxx

	// x 业务编号 1 下单
	// xxxxxx  年月日 6
	// xxxxx 时间戳 最后 10位数
	// xxxx  用户Guid 最后 4位
	// 总共 21 位

	no := "3"

	t := time.Now()

	nt := string(t.Format("060102")) // xxxxxx  年月日
	timeUnixNano := int(t.UnixNano())
	tunS := strconv.Itoa(timeUnixNano)[6:16]

	userGuidS := userGuid[len(userGuid)-4 : len(userGuid)]

	no += nt + tunS + userGuidS

	return no
}
