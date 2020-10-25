package tool

import (
	"time"
)

// TimeUnix
// 获取当前时间戳
func TimeUnix() int64 {
	return time.Now().Unix()
}

// DateTime
// 获取当前时间
func DateTime(layout string) string {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.Now().Format(layout)
}

// FormatStringToTime
// 字符串转换成Time
func FormatStringToTime(layout, value string) (time.Time, error) {
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	return time.ParseInLocation(layout, value, time.Now().Location())
}
