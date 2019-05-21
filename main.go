package main

import (
	"log"
	"net/http"

	"./routes"

	"github.com/go-chi/chi"
)

func main() {
	r := routes.Router()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("error: %s", err.Error())
	}

	log.Fatal(http.ListenAndServe(":4500", r))
}
