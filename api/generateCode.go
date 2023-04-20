package api

import "net/http"

func generateCode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
