package compress

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestFlateData(t *testing.T) {
	data := "this is a message 这是一个消息"
	t.Log(len(data))
	// 压缩数据
	rs, err := FlateEncode([]byte(data), -1)
	// rs, err := FlateEncodeV(data)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(len(rs))

	// 解压数据
	res, err := FlateDecode(rs)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(len(res), string(res))
}

func TestGzipEncode(t *testing.T) {
	data := "this is a message"
	rs := GzipEncode([]byte(data))
	t.Log(len(data))
	t.Log(len(rs))
	if res, err := GzipDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}

}

func TestGzipLevelEncode(t *testing.T) {
	data := `this is a message；239o8jv9接口路径而；a9324jva9023v速看的浪费加埃里克顺道减肥alkjfaaoi挨饿的等级
	kj34j9fqu案例二京东方jfaaldjf阿菊发阿施蒂利克埃里克对接放埃里克地方就埃里克记得发完全二问沙发阿斯蒂芬阿斯蒂芬发
	阿斯利康对接放曲儿GV打发va啊的说法昂防守打法(⊙o⊙)…5帅哥商定发阿斯蒂芬gas大飞哥阿达发送到阿斯顿发达案发生的发生发
	阿萨德，放开就埃里克森2342访问发起个分ads法 aefja3q904kajflk埃里克染发剂阿里09i90jojlkjaf阿里副科级阿拉山口的`
	rs, err := GzipLevelEncode([]byte(data), 9)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("原数据长度:", len(data))
	t.Log("压缩数据长度", len(rs))
	if res, err := GzipDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}
}

func TestZlibEncode(t *testing.T) {
	data := "this is a message"
	rs := ZlibEncode([]byte(data))
	t.Log(len(data))
	t.Log(len(rs))
	if res, err := ZlibDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}
}

func TestZlibLevelEncode(t *testing.T) {
	data := `this is a message；239o8jv9接口路径而；a9324jva9023v速看的浪费加埃里克顺道减肥alkjfaaoi挨饿的等级
	kj34j9fqu案例二京东方jfaaldjf阿菊发阿施蒂利克埃里克对接放埃里克地方就埃里克记得发完全二问沙发阿斯蒂芬阿斯蒂芬发
	阿斯利康对接放曲儿GV打发va啊的说法昂防守打法(⊙o⊙)…5帅哥商定发阿斯蒂芬gas大飞哥阿达发送到阿斯顿发达案发生的发生发
	阿萨德，放开就埃里克森2342访问发起个分ads法 aefja3q904kajflk埃里克染发剂阿里09i90jojlkjaf阿里副科级阿拉山口的`
	rs, err := ZlibLevelEncode([]byte(data), 9)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("原数据长度:", len(data))
	t.Log("压缩数据长度", len(rs))
	if res, err := ZlibDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}
}

func TestBrotliEncode(t *testing.T) {
	data := "this is a message"
	rs := BrotliEncode([]byte(data))
	t.Log(len(data))
	t.Log(len(rs))
	if res, err := BrotliDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}
}

func TestBrotliLevelEncode(t *testing.T) {
	data := `{
		"code": 53000,
		"msg": "error1d233244",
		"body": {}
	}`
	rs := BrotliLevelEncode([]byte(data), 6)
	t.Log("原数据长度:", len(data))
	t.Log("压缩数据长度", len(rs))
	if res, err := BrotliDncode(rs); err == nil {
		t.Log(string(res))
	} else {
		t.Error(err.Error())
	}
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
