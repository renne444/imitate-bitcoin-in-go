package util

import "time"

//GetTimestampNow :获取当前时间戳
func GetTimestampNow() int64 {
	t := time.Now()
	return t.Unix()
}
