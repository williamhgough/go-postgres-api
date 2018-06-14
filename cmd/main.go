package main

import (
	"log"
	"net/http"
	"time"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
	"github.com/williamhgough/go-postgres-api/controllers"
)

func init() {
	// Need to allow time for postgres to migrate and seed DB before running.
	time.Sleep(3 * time.Second)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", controllers.Index)

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)

	log.Println("Starting web server on http://localhost:8080")
	graceful.Run(":8080", 2*time.Second, n)
}
