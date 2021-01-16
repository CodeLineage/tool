package tool

import "testing"

func TestGoroutineID(t *testing.T) {
	rs := GoroutineID()
	t.Log(rs)
}
