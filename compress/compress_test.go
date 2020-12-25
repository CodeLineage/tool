package tool

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestGzipData(t *testing.T) {
	data := "this is a mess"
	rs := GzipData(data)
	t.Log(len(data))
	t.Log(len(rs))
}

func TestPtrType(t *testing.T) {
	// 字符串
	data := "this is a message"
	ptr := &data
	t.Log(reflect.ValueOf(ptr).Kind(), reflect.ValueOf(ptr).Pointer(), ptr)

	// []byte
	dataByte := []byte("this is a message")
	ptrType := &dataByte
	t.Log(reflect.ValueOf(ptrType).Kind(), ptrType)
	_ = *(*string)(unsafe.Pointer(&dataByte))
}
