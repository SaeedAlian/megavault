package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func ParseJSONFromRequest(r *http.Request, payload any) error {
	body := r.Body

	if body == nil {
		return fmt.Errorf("Request body is not found")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSONInResponse(
	w http.ResponseWriter,
	status int,
	payload any,
	headers *map[string]string,
) error {
	if headers != nil {
		for k, v := range *headers {
			w.Header().Add(k, v)
		}
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

func WriteErrorInResponse(w http.ResponseWriter, status int, message string) error {
	return WriteJSONInResponse(w, status, map[string]string{"message": message}, nil)
}
