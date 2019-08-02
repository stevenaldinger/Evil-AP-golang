package httpserver

import (
	"net/http"
)

func Serve() {
  http.Handle("/", http.FileServer(http.Dir("./static")))

  http.ListenAndServe(":80", nil)
}
