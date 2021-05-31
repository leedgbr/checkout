package main

import (
	"log"
	"net/http"

	"company.com/checkout/internal/checkout"
)

func main() {
	app := checkout.App{}
	app.Init()
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
