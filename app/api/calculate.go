package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tverytinov/test-task_calculate-factorial/app/model"
)

func Calculate(w http.ResponseWriter, m map[string]int) {
	chA := make(chan int)
	chB := make(chan int)

	go factorialCount(m["a"], chA)
	go factorialCount(m["b"], chB)

	input := model.Input{
		A: <-chA,
		B: <-chB,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
