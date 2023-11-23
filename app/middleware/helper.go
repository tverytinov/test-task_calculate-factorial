package middleware

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/tverytinov/test-task_calculate-factorial/app/model"
)

var (
	negativeOrOutError = "Incorrect input"
)

type ResponseWriterAndMap func(http.ResponseWriter, map[string]int)

func parseInput(body io.Reader) (map[string]int, error) {
	data := make(map[string]int)

	if err := json.NewDecoder(body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func validateInput(data map[string]int) error {
	if _, ok := data["a"]; !ok || data["a"] < 0 {
		return errors.New(negativeOrOutError)
	}

	if _, ok := data["b"]; !ok || data["b"] < 0 {
		return errors.New(negativeOrOutError)
	}

	return nil
}

func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(model.Response{Error: err.Error()}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
