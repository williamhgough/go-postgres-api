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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", controllers.Index)

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)

	log.Println("Starting web server on http://localhost:8080")
	graceful.Run(":8080", 2*time.Second, n)
}
