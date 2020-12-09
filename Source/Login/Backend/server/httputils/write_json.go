package httputils

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, value interface{}, statusCode int) error {
	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err = w.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
