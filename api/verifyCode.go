package api

import (
	"encoding/json"
	"github.com/aagolovanov/awesomeCodeService/domain"
	"net/http"
)

func (s *Server) verifyCode(w http.ResponseWriter, r *http.Request) {
	request := domain.RequestWithCode{}
	// FIXME : JSON валидация
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(apiError{
			Error: "bad request body provided",
		})
		_, _ = w.Write(body)
		return
	}

	verifMessage, err := s.domain.VerifyCode(&request)
	if err != nil {
		if err.Error() == "internal" {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else if err.Error() == "limit" {
			w.WriteHeader(http.StatusTooManyRequests) // maybe another statuscode?

			body, _ := json.Marshal(apiError{
				Error: "Verification attempts limit has been reached",
			})

			_, err = w.Write(body)
			//_, err := fmt.Fprint(w, body)
			if err != nil {
				s.logg.Printf("Error while writing response: %v\n", err)
			}

			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			body, _ := json.Marshal(apiError{
				Error: err.Error(),
			})
			_, err = w.Write(body)
			//_, err := fmt.Fprint(w, body)
			if err != nil {
				s.logg.Printf("Error while writing response: %v\n", err)
			}
			return
		}
	}

	body, _ := json.Marshal(verifMessage)
	_, err = w.Write(body)
	if err != nil {
		s.logg.Printf("Error while writing response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
