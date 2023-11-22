package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	port               = "8989"
	negativeOrOutError = "Incorrect input"
)

type ResponseWriterAndMap func(http.ResponseWriter, map[string]int)

type Input struct {
	A int `json:"a"`
	B int `json:"b"`
}

type Response struct {
	Error string `json:"error"`
}

func main() {
	router := httprouter.New()
	router.POST("/calculate", middleware(calculate))

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func calculate(w http.ResponseWriter, m map[string]int) {
	chA := make(chan int)
	chB := make(chan int)

	go factorialCount(m["a"], chA)

	go factorialCount(m["b"], chB)

	input := Input{
		A: <-chA,
		B: <-chB,
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(input); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func factorialCount(n int, ch chan int) {
	result := 1

	if n == 0 || n == 1 {
		ch <- 1
		return
	}

	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
}

func middleware(next ResponseWriterAndMap) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data := make(map[string]int)

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := data["a"]; !ok || data["a"] < 0 {
			incorrectInput(w)
			return
		}

		if _, ok := data["b"]; !ok || data["b"] < 0 {
			incorrectInput(w)
			return
		}

		next(w, data)
	}
}

func incorrectInput(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	if err := json.NewEncoder(w).Encode(Response{Error: negativeOrOutError}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
