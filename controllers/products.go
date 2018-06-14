package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williamhgough/go-postgres-api/models"
	"github.com/williamhgough/go-postgres-api/utils/cache"
	"github.com/williamhgough/go-postgres-api/utils/database"
	"github.com/williamhgough/go-postgres-api/utils/respond"
)

// Index Handles listing and creating products
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if data, err := cache.Fetch(r); err == nil {
			respond.Call(w, r, http.StatusOK, data)
		} else {
			p, err := models.GetProducts(database.GetDB())
			if err != nil {
				respond.Error(w, r, http.StatusNoContent, err)
			}

			err = cache.Store(r, p)
			if err != nil {
				log.Print("couldn't set product in cache: ", err)
			}

			respond.Call(w, r, http.StatusOK, p)
		}
	case "POST":
		respond.Error(w, r, http.StatusInternalServerError, "not implemented yet")
	default:
		respond.Error(w, r, http.StatusBadRequest, "invalid request method")
	}
}

// Product handles retrieving, updating and deletion
// of a single product by ID.
func Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if data, err := cache.Fetch(r); err == nil {
			respond.Call(w, r, http.StatusOK, data)
		} else {
			vars := mux.Vars(r)
			p, err := models.GetProduct(vars["id"], database.GetDB())
			if err != nil {
				respond.Error(w, r, http.StatusNoContent, err)
			}

			err = cache.Store(r, p)
			if err != nil {
				log.Print("couldn't set product in cache: ", err)
			}

			respond.Call(w, r, http.StatusOK, p)
		}
	case "PUT":
		respond.Error(w, r, http.StatusInternalServerError, "not implemented yet")
	case "DELETE":
		respond.Error(w, r, http.StatusInternalServerError, "not implemented yet")
	default:
		respond.Error(w, r, http.StatusBadRequest, "invalid request method")
	}
}
