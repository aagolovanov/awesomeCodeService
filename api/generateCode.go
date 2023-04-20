package api

import (
	"encoding/json"
	"github.com/aagolovanov/awesomeCodeService/domain"
	"net/http"
)

func generateCode(w http.ResponseWriter, r *http.Request) {

	request := domain.RequestGenerate{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// TODO write error
		return
	}

	w.WriteHeader(http.StatusOK)
}
