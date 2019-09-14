package app

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/favclip/ucon"
)

func AppInit(gae bool) *ucon.ServeMux {
	ucon.Orthodox()
	ucon.HandleFunc("POST", "/command", func(w http.ResponseWriter, r *http.Request) {
		params, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Printf("%#v\n", string(params))
		w.Write([]byte("Hello World!"))
	})
	ucon.DefaultMux.Prepare()
	return ucon.DefaultMux
}
