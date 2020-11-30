package tool

import "testing"

func TestSha256(t *testing.T) {
	str := "this is message"
	rs := Sha256(str)
	t.Log(rs)
}
