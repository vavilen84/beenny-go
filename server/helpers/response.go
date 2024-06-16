package helpers

import (
	"net/http"
)

func setCacheHeaders(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
}

func setContentTypeHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func WriteResponse(w http.ResponseWriter, b []byte, status int) {
	setCacheHeaders(w)
	setContentTypeHeader(w)
	w.WriteHeader(status)
	_, err := w.Write(b)
	if err != nil {
		LogError(err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
