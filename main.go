package main

import (
	"net/http"

	"github.com/gumpen/slash_test/app"
	"github.com/joho/godotenv"
)

func main() {
	envLoad()
	router := app.AppInit(true)
	http.ListenAndServe(":8080", router)
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
