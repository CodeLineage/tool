package tool

import(
	"time"
)

func TimeUnix() int64{
	return time.Now().Unix()
}