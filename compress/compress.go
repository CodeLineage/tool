package tool

import (
	"compress/gzip"
	"strings"
)

// GzipData
// gzip压缩数据
func GzipData(data string) string {
	content := &strings.Builder{}
	gzipWrite := gzip.NewWriter(content)
	defer gzipWrite.Close()
	gzipWrite.Write([]byte(data))
	return content.String()
}
