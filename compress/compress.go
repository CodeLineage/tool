package compress

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"io/ioutil"
)

// FlateEncode
// deflate压缩数据
// @param data 待压缩的数据
// @param level 压缩等级
func FlateEncode(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	flateWrite, err := flate.NewWriter(&buffer, level)
	if err != nil {
		return nil, err
	}
	// Fixed bug 为什么不能使用defer调用，Close不仅是释放资源，还有flush
	// defer flateWrite.Close()
	flateWrite.Write(data)
	flateWrite.Close()
	return buffer.Bytes(), nil
}

// FlateDecode
// deflate解压数据
// @param data 待解压的数据
func FlateDecode(data []byte) ([]byte, error) {
	content := bytes.NewBuffer(data)
	flateReader := flate.NewReader(content)
	return ioutil.ReadAll(flateReader)
}

// GzipEncode
// gzip压缩数据
// @param data 待压缩的数据
func GzipEncode(data []byte) []byte {
	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)
	gzipWriter.Write(data)
	gzipWriter.Close()
	return buffer.Bytes()
}

// GzipLevelEncode
// gzip压缩数据
// @param data 待压缩的数据
// @param level 压缩等级
func GzipLevelEncode(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	gzipWrite, err := gzip.NewWriterLevel(&buffer, level)
	if err != nil {
		return nil, err
	}

	gzipWrite.Write(data)
	gzipWrite.Close()
	return buffer.Bytes(), nil
}

// GzipDncode
// 解压gzip压缩数据
func GzipDncode(data []byte) ([]byte, error) {
	content := bytes.NewReader([]byte(data))
	gzipReader, err := gzip.NewReader(content)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(gzipReader)
}

// ZlibEncode
// zlib压缩数据
// @param data 待压缩的数据
func ZlibEncode(data []byte) []byte {
	var buffer bytes.Buffer
	zlibWriter := zlib.NewWriter(&buffer)
	zlibWriter.Write(data)
	zlibWriter.Close()
	return buffer.Bytes()
}

// ZlibLevelEncode
// zlib压缩数据
// @param data 待压缩的数据
// @param level 压缩等级
func ZlibLevelEncode(data []byte, level int) ([]byte, error) {
	var buffer bytes.Buffer
	zlibWriter, err := zlib.NewWriterLevel(&buffer, level)
	if err != nil {
		return nil, err
	}

	zlibWriter.Write(data)
	zlibWriter.Close()
	return buffer.Bytes(), nil
}

// GzipDncode
// 解压gzlib压缩数据
func ZlibDncode(data []byte) ([]byte, error) {
	content := bytes.NewReader([]byte(data))
	zlibReader, err := zlib.NewReader(content)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(zlibReader)
}
