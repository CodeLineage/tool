### Golang常见工具库

#### time.go文件 时间操作

| 函数 | 说明 |
| --- | --- |
| TimeUnix() | 获取当前时间戳 |
| DateTime() | 获取当前时间 |
| FormatStringToTime() | 字符串转换成Time |

#### format.go文件 类型转换

| 函数 | 说明 |
| --- | --- |
| FormatInt(data interface{}) (int, error) | 接口类型转换到整型 |
| FormatInt64(data interface{}) (int64, error) | 接口类型转换到int64 |
| FormatFloat(data interface{}) (float64, error) | 接口类型转换到float64 |
| FormatString(data interface{}) (string, error) | 接口类型转换成string |
