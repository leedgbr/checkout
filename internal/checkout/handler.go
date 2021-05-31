package checkout

import (
	"encoding/json"
	"net/http"
)

type service interface {
	checkout(sku []string) (*result, error)
}

type request struct {
	SKU []string `json:"sku"`
}

type errorResponse struct {
	Message string `json:"message"`
}

var (
	badRequestError    = errorResponse{Message: "error parsing request"}
	systemErrorMessage = "system error"
)

func NewSimpleHandler() func(http.ResponseWriter, *http.Request) {
	return newHandler(newSimpleService())
}

func newHandler(s service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		marshalledRequest := &request{}
		if err := json.NewDecoder(r.Body).Decode(&marshalledRequest); err != nil {
			respondWithBadRequestError(w)
			return
		}
		result, err := s.checkout(marshalledRequest.SKU)
		if err != nil {
			respondWithBusinessError(w, err)
			return
		}
		encode(w, result)
	}
}

func respondWithBusinessError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	encode(w, errorResponse{Message: err.Error()})
}

func respondWithBadRequestError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	encode(w, badRequestError)
}

func encode(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, systemErrorMessage, 500)
	}
}
