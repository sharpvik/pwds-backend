package iputils

import (
	"testing"
	"net/http"
	"strings"
)



func TestReadIP(t *testing.T) {
	r, _ := http.NewRequest("GET", "localhost:8000", strings.NewReader(""))
	r.RemoteAddr = "127.0.0.1:40021"
	cleanIP := ReadIP(r)

	if cleanIP != r.RemoteAddr {
		t.Error("ReadIP function doesn't work!")
	}
}



func TestReadCleanIP(t *testing.T) {
	r, _ := http.NewRequest("GET", "localhost:8000", strings.NewReader(""))
	r.RemoteAddr = "127.0.0.1:40021"
	cleanIP := ReadCleanIP(r)

	if cleanIP != "127.0.0.1" {
		t.Error("ReadCleanIP function doesn't work!")
	}
}
