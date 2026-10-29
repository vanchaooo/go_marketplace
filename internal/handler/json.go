package handler

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func decodeJSON(r *http.Request, destination any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(destination); err != nil {
		return err
	}

	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return errors.New("request should have only one JSON object")
	}

	return nil
}

func writeJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Set("Content-Type", "aplication/JSON")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	_ = writeJSON(w, status, map[string]string{"message": message})
}