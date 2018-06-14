package main

import (
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
	"github.com/williamhgough/go-postgres-api/controllers"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/products", controllers.Index).
		Methods("GET", "POST")
	mux.HandleFunc("/api/products/{id:[0-9]+}", controllers.Product).
		Methods("GET", "PUT", "DELETE")

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)

	log.Println("Starting web server on http://localhost:8080")
	graceful.Run(":8080", 2*time.Second, n)
}
