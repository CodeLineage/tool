package tool

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GoroutineID
// 获取协程ID
func GoroutineID() int {
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Sprintf("panic recover:panic info:%v", err))
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
