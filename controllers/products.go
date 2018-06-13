package controllers

import (
	"log"
	"net/http"

	"github.com/williamhgough/pql-api/models"
	"github.com/williamhgough/pql-api/utils"
)

// Index Handles base products route
func Index(w http.ResponseWriter, r *http.Request) {
	products, err := models.GetProducts(utils.GetDB())
	if err != nil {
		utils.RespondErr(w, r, http.StatusNoContent, err)
	}

	log.Println("Successful request")
	utils.Respond(w, r, http.StatusOK, products)
}
