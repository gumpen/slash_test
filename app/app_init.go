package app

import (
	"net/http"

	"github.com/favclip/ucon"
)

func AppInit(gae bool) *ucon.ServeMux {
	ucon.Orthodox()
	ucon.HandleFunc("POST", "/command", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	ucon.DefaultMux.Prepare()
	return ucon.DefaultMux
}
