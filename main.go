package main

import (
    "io"
    "os"
    "net/http"
    "fmt"
    "log"

    "github.com/sharpvik/pwds-backend/apps/root"
    "github.com/sharpvik/pwds-backend/apps/join"
    "github.com/sharpvik/pwds-backend/apps/favicon"
    "github.com/sharpvik/pwds-backend/iputils"
    "github.com/sharpvik/pwds-backend/config"
)



var logr *log.Logger



var mux map[string]func(http.ResponseWriter, *http.Request)



type mainHandler struct {}

// ServeHTTP is a top-level function that uses mux to delegate to a proper view.
// This function exists so as to allow mainHandler to implement http.Handler
// interface.
func (*mainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ip := iputils.ReadIP(r)

	if fn, ok := mux[r.URL.String()]; ok {
        logr.Printf( "%s\n\tURL: %s\n\t=> OK", ip, r.URL.String() )
        fn(w, r)
        return
    }

    notOK := fmt.Sprintf( "%s\n\tURL: %s\n\t=> NOT OK", ip, r.URL.String() )

    logr.Print(notOK)
    io.WriteString(w, notOK)
}



func main() {
    logr = log.New(os.Stdout, "", log.Ltime)


	server := http.Server{
		Addr: config.Port,
		Handler: &mainHandler{},
	}


	mux = make( map[string]func(http.ResponseWriter, *http.Request) )
    mux["/"] = root.Index
    mux["/join"] = join.Index
    mux["/favicon.ico"] = favicon.Index


    logr.Printf("Serving at localhost%s", config.Port)
    server.ListenAndServe()
}
