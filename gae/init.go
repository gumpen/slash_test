package gae

import (
	"net/http"

	"github.com/gumpen/slack_why_bot/app"
)

func init() {
	router := app.AppInit(true)
	http.Handle("/", router)
}
