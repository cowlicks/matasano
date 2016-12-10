package hmac

import (
	"time"
	"encoding/hex"
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestEmpty(t *testing.T) {
	o := HmacSha1([]byte(""), []byte(""))
	e := "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d"
	if !(hex.EncodeToString(o) == e) {
		t.Fail()
	}
}

func Test(t *testing.T) {
	o := HmacSha1([]byte("key"), []byte("The quick brown fox jumps over the lazy dog"))
	e := "de7c9b85b8b78aa6bc8a7a36f70a90701c9db4d9"
	if !(hex.EncodeToString(o) == e) {
		t.Fail()
	}
}

func TestHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	res, _ := http.Get(ts.URL)
	if res.StatusCode != http.StatusBadRequest {
		t.Fail()
	}

	sig := "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d"
	fullurl := ts.URL + "/?file=&signature=" + sig
	res, _ = http.Get(fullurl)
	if res.StatusCode != http.StatusOK {
		t.Fail()
	}
}

// this is slow af
func TestTimingAttack(t *testing.T) {
	out := ""
	for i := 0; i < 20; i++ {
		times := make([]int, 256)
		for j := 0; j < 256; j++ {
			guess := hex.EncodeToString([]byte{byte(j)})
			fullguess :=  out + guess
			sigbytes, _ := hex.DecodeString(fullguess)
			before := time.Now()
			badEq(HmacSha1(KEY, []byte("")), sigbytes)
			after := time.Now()
			times[j] = int(after.Sub(before))
		}
		out = out + hex.EncodeToString([]byte{byte(ArgMax(times))})
		P(out)
	}
	sig := "fbdb1d1b18aa6c08324b7d64b71fb76370690e1d"
	P(sig)
}
