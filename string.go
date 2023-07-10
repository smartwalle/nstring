package nstring

import (
	"reflect"
	"unsafe"
)

// Sub 获取子串
func Sub(s string, start int, length int) string {
	if length == 0 {
		return ""
	}

	if start < 0 {
		start = 0
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

func Chunk(s string, size int) []string {
	var sSize = len(s)
	if size <= 0 || sSize == 0 {
		return nil
	}

	if size >= sSize {
		return []string{s}
	}

	var cSize = ((sSize - 1) / size) + 1
	var chunks = make([]string, 0, cSize)

	var start = 0
	var count = 0
	for i := range s {
		if count == size {
			chunks = append(chunks, s[start:i])
			count = 0
			start = i
		}
		count++
	}
	chunks = append(chunks, s[start:])

	return chunks
}

// ToBytes 将 string 转换成 byte 数组
//
// 注意：不要修改返回的 byte 数组的值
func ToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// FromBytes 将 byte 数组转换成 string
func FromBytes(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
