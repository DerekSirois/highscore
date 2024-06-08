package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type JsonResponse struct {
	Message string `json:",omitempty"`
	Token   string `json:",omitempty"`
	Data    any    `json:",omitempty"`
}

func Handler(next func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := next(w, r); err != nil {
			res := JsonResponse{
				Message: err.Error(),
			}
			WriteJson(w, http.StatusBadRequest, res)
		}
	}
}

func WriteJson(w http.ResponseWriter, statusCode int, data JsonResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("error while encoding the data: %v\n", err)
	}
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")

	if tokenAuth == "" {
		return ""
	}

	return tokenAuth[7:] // remove bearer
}
