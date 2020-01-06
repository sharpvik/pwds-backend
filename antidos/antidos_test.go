package antidos

import (
	"testing"
	"time"
	"net/http"
	"strings"
)




func TestNotice(t *testing.T) {
	r, _ := http.NewRequest("POST", "localhost:8000", strings.NewReader(""))
	r.RemoteAddr = "127.0.0.1"

	ads := New(10, 10 * time.Second)
	count := 0

	for i := 0; i < 15; i++ {
		if ads.Notice(r) {
			count++
		}
	}

	time.Sleep(15 * time.Second)

	for i := 0; i < 15; i++ {
		if ads.Notice(r) {
			count++
		}
	}

	if count != 20 {
		t.Error("Function `Notice` works incorrectly.")
	}
}