package root

import (
	"io"
    "net/http"
)



// Index is a handler function for any request to root URL `/`.
func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "`/signup` or `/login` -- there's no other way in!")
}
