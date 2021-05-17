package tools

import (
	"math/rand"
	"reflect"
)

var alphas = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
}

// GetStr 获取定长字符串
func GetStr(len uint32) string {
	s := ""
	for i := uint32(0); i < len; i++ {
		s += alphas[rand.Intn(36)]
	}

	return s
}

// InSlice 判断item是否是slice的元素
func InSlice(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		return false
	}
	if reflect.TypeOf(slice).Elem() != reflect.TypeOf(item) {
		return false
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
