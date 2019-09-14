package main

import (
	"net/http"

	"github.com/gumpen/slash_test/app"
)

func main() {
	router := app.AppInit(true)
	http.ListenAndServe(":8080", router)
}
