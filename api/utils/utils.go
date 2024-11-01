package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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

func CreateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")

	slug = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			return r
		}
		return -1
	}, slug)

	return slug
}

