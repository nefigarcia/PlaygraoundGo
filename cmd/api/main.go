package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8082

type application struct {
	Domain string
}
func main() {
	var app application 
	app.Domain= "nefi.com"
	log.Println("Starting app on port", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
