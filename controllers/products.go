package controllers

import (
	"net/http"

	"github.com/williamhgough/go-postgres-api/models"
	"github.com/williamhgough/go-postgres-api/utils"
)

// Index Handles base products route
func Index(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts(utils.GetDB())
	if err != nil {
		utils.RespondErr(w, r, http.StatusNoContent, err)
	}

	utils.Respond(w, r, http.StatusOK, products)
}
