package join

import (
    "io"
    "fmt"
    "time"
    "net/http"
    "encoding/json"

    "github.com/sharpvik/pwds-backend/antidos"
)



var antiDoS = antidos.New(100, 24 * time.Hour)



// Index is a handler function for any request to `/join`.
func Index(w http.ResponseWriter, r *http.Request) {
    if !antiDoS.Notice(r) {
        io.WriteString(w, "You've reached your join limit for today.")
        return
    }

    user := NewUser()

    defer r.Body.Close()
    err := json.NewDecoder(r.Body).Decode(user)

    if err != nil {
        io.WriteString(
            w,
            fmt.Sprintf("Wrong user data: %s", err),
        )
        return
    }

    user.Store()

    io.WriteString(
        w,
        fmt.Sprintf("Added user: %s:%s", user.Username, user.Masterhash),
    )
}
