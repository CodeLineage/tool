package tool

import (
	"crypto/sha256"
	"testing"
)

func TestSha256(t *testing.T) {
	str := "this is message"
	rs := Sha256(str)
	t.Log(rs)
}

func TestHmacHash(t *testing.T) {
	data := `{"timestamp":"20201223134000","data":{"keyword":"玄幻","pageSize":"2","page":"1","requestOrgin":"00"}}`
	key := `Mp0yEDqJYguA/uqjRXQVGojMtBx0nU1i4tGDDVrfRujt3lN8y40hIw1vvS6BbtBwcNfJE2wOASg/FHBmdcP1IPrrUo2Hg=uxiA1hDwUqUg87jizuxf5cAopzXnqyTSyoco/MbNVnuAcC0k3jXuD5crYt+YGICgHFVPekmvzCmSl/`

	rs := HmacHash(sha256.New, key, data)
	t.Log(rs)
}
