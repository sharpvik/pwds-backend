package favicon

import (
	"net/http"
	"path"
	
    "github.com/sharpvik/pwds-backend/config"
)



// Index is a handler function for any request to `/favicon.ico`
func Index(w http.ResponseWriter, r *http.Request) {
	favicon := path.Join(
		config.RootFolder, config.GlobalStaticFolder, "img/favicon.png",
	)

    http.ServeFile(w, r, favicon)
}
