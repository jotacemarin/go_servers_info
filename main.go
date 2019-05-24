package main

import (
	"log"
	"net/http"

	"./config"
	"./db"
	"./routes"

	"github.com/go-chi/chi"
)

func main() {
	configuration, errConfig := config.LoadConfig()
	if errConfig != nil {
		log.Panicf("error: %s", errConfig.Error())
	}
	r := routes.Router()
	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route)
		return nil
	}
	if err := chi.Walk(r, walkFunc); err != nil {
		log.Panicf("error: %s", err.Error())
	}
	db.CreateTable(db.Db)
	log.Fatal(http.ListenAndServe(configuration.Port, r))
}
