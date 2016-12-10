/*
Implement and break HMAC-SHA1 with an artificial timing leak The psuedocode on
Wikipedia should be enough. HMAC is very easy.

Using the web framework of your choosing (Sinatra, web.py, whatever), write a
tiny application that has a URL that takes a "file" argument and a "signature"
argument, like so:

http://localhost:9000/test?file=foo&signature=46b4ec586117154dacd49d664e5d63fdc88efb51
Have the server generate an HMAC key, and then verify that the "signature" on
incoming requests is valid for "file", using the "==" operator to compare the
valid MAC for a file with the "signature" parameter (in other words, verify the
HMAC the way any normal programmer would verify it).

Write a function, call it "insecure_compare", that implements the == operation
by doing byte-at-a-time comparisons with early exit (ie, return false at the
first non-matching byte).

In the loop for "insecure_compare", add a 50ms sleep (sleep 50ms after each
byte).

Use your "insecure_compare" function to verify the HMACs on incoming requests,
and test that the whole contraption works. Return a 500 if the MAC is invalid,
and a 200 if it's OK.

Using the timing leak in this application, write a program that discovers the
valid MAC for any file.
*/

package hmac

import (
	"time"
    "net/http"
	"encoding/hex"
)

var KEY = []byte("")

func badEq(a, b []byte) bool {
	l := len(a)
	if len(a) > len(b) {
		l = len(b)
	}
	for i := 0; i < l; i++ {
		if a[i] != b[i] {
			return false
		}
		time.Sleep(2 * time.Millisecond)
	}
	return true
}

func handler(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	file, ok := vals["file"]

	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sig, ok := vals["signature"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	sigbytes, err := hex.DecodeString(sig[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res := badEq(HmacSha1(KEY, []byte(file[0])), sigbytes)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func ArgMax(times []int) int {
	maxi := 0
	maxv := 0
	for i, v := range times {
		if v > maxv {
			maxv = v
			maxi = i
		}
	}
	return maxi
}

type DummyWriter struct {}

func (d * DummyWriter) Header() http.Header {
	return nil
}
func (d * DummyWriter) Write(b []byte) (int, error) {
	return 0, nil
}
func (d * DummyWriter) WriteHeader(b int) {
	return
}


