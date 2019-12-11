package util

import (
	"strconv"
	"testing"
)

func TestGetTimeNow(t *testing.T) {
	now := GetTimestampNow()
	strN := strconv.FormatInt(now, 16)
	println(strN)
}
