package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matheusandrade23/go-bid/internal/validator"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, payload T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		return fmt.Errorf("failed to encode data json %w", err)
	}

	return nil
}

func DecodeValidJson[T validator.Validator](r *http.Request) (T, map[string]string, error) {
	var payload T

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return payload, nil, fmt.Errorf("decode json %w", err)
	}

	if problems := payload.Valid(r.Context()); len(problems) > 0 {
		return payload, problems, fmt.Errorf("invalid %T: %d problems", payload, len(problems))
	}

	return payload, nil, nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json failed: %w", err)
	}
	return data, nil
}