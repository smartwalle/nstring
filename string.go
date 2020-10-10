package str4go

import (
	"reflect"
	"unsafe"
)

// SubStr 获取子串
func SubStr(s string, start int, length int) string {
	if length == 0 {
		return ""
	}

	if start < 0 {
		start = 0
	}
	if length < -1 {
		length = -1
	}

	var sRunes = []rune(s)
	var rLen = len(sRunes)

	if start >= rLen {
		return ""
	}

	if length < 0 {
		return string(sRunes[start:])
	}

	var end = start + length

	if end > rLen {
		end = rLen
	}
	return string(sRunes[start:end])
}

// StringToBytes 将 string 转换成 byte 数组
// 注意：不要修改返回的 byte 数组的值
func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// BytesToString 将 byte 数组转换成 string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
