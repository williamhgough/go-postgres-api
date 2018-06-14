package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func decodeBody(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}

func encodeBody(w http.ResponseWriter, r *http.Request, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// Respond should be used for all successful requests
// to encode and return data.
// e.g utils.Respond(w, r, http.StatusOk, products)
func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Location")
	if data != nil {
		encodeBody(w, r, data)
	}
}

// RespondErr should be used when returning custom
// error messages.
// e.g utils.RespondErr(w, r, http.StatusNoContent, "no products found in db")
func RespondErr(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Respond(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

// RespondHTTPErr can be used when returning default
// HTTP error text.
// e.g utils.RespondHTTPErr(w, r, http.StatusNotFound)
func RespondHTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	RespondErr(w, r, status, http.StatusText(status))
}
