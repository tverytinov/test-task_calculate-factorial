package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Middleware(next ResponseWriterAndMap) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		data, err := parseInput(r.Body)
		if err != nil {
			handleError(w, err)
			return
		}

		if err := validateInput(data); err != nil {
			handleError(w, err)
			return
		}

		next(w, data)
	}
}
