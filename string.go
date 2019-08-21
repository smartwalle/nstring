package str4go

import "unsafe"

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

func StrToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToStr(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
