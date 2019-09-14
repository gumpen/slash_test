package gae

import (
	"net/http"

	"github.com/gumpen/slash_test/app"
)

func init() {
	router := app.AppInit(true)
	http.Handle("/", router)
}
