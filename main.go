package main

import (
	"net/http"

	"github.com/gumpen/slack_why_bot/app"
)

func main() {
	router := app.AppInit(true)
	http.ListenAndServe(":8080", router)
}
