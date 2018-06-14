package respond

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
	return json.NewEncoder(w).Encode(v)
}

// Call should be used for all succesful requests
// to encode and return data.
// e.g respond.Call(w, r, http.StatusOk, products)
func Call(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Location")
	if data != nil {
		encodeBody(w, r, data)
	}
}

// Error should be used when returning custom
// error messages.
// e.g respond.Error(w, r, http.StatusNoContent, "no products found in db")
func Error(w http.ResponseWriter, r *http.Request, status int, args ...interface{}) {
	Call(w, r, status, map[string]interface{}{
		"error": map[string]interface{}{
			"message": fmt.Sprint(args...),
		},
	})
}

// HTTPErr can be used when returning default
// HTTP error text.
// e.g respond.HTTPErr(w, r, http.StatusNotFound)
func HTTPErr(w http.ResponseWriter, r *http.Request, status int) {
	Error(w, r, status, http.StatusText(status))
}
