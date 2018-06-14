package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/williamhgough/go-postgres-api/models"
	"github.com/williamhgough/go-postgres-api/utils"
)

// Index Handles listing and creating products
func Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		products, err := models.GetProducts(utils.GetDB())
		if err != nil {
			utils.RespondErr(w, r, http.StatusNoContent, err)
		}

		utils.Respond(w, r, http.StatusOK, products)
	case "POST":
		utils.RespondErr(w, r, http.StatusInternalServerError, "not implemented yet")
	default:
		utils.RespondErr(w, r, http.StatusBadRequest, "invalid request method")
	}
}

// Product handles retrieving, updating and deletion
// of a single product by ID.
func Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		product, err := models.GetProduct(vars["id"], utils.GetDB())
		if err != nil {
			utils.RespondErr(w, r, http.StatusNoContent, err)
		}
		utils.Respond(w, r, http.StatusOK, product)
	case "PUT":
		utils.RespondErr(w, r, http.StatusInternalServerError, "not implemented yet")
	case "DELETE":
		utils.RespondErr(w, r, http.StatusInternalServerError, "not implemented yet")
	default:
		utils.RespondErr(w, r, http.StatusBadRequest, "invalid request method")
	}
}
