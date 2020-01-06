package iputils

import (
    "net/http"
    "strings"
)



// ReadIP retrievs sender's IP address from http.Request object.
func ReadIP(r *http.Request) string {
	IPAddress := r.Header.Get("X-Real-Ip")
	
    if IPAddress == "" {
        IPAddress = r.Header.Get("X-Forwarded-For")
	}
	
    if IPAddress == "" {
        IPAddress = r.RemoteAddr
	}
	
    return IPAddress
}



// ReadCleanIP uses ReadIP function to retrieve sender's IP address, chops off
// the port and returns the "clean" IP.
//
// For example:
//
//     r, _ := http.NewRequest("GET", "localhost:8000", strings.NewReader(""))
//     r.RemoteAddr = "127.0.0.1:40021"
//     cleanIP := ReadCleanIP(r)
//     // Output: "127.0.0.1"
//     
//     if cleanIP != "127.0.0.1" {
//         panic("ReadCleanIP function doesn't work!")
//     }
//
func ReadCleanIP(r *http.Request) string {
    return strings.Split(ReadIP(r), ":")[0]
}
