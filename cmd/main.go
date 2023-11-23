package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	handler "github.com/tverytinov/test-task_calculate-factorial/app/api"
	"github.com/tverytinov/test-task_calculate-factorial/app/middleware"
)

var port string = "8989"

func main() {
	router := httprouter.New()
	router.POST("/calculate", middleware.Middleware(handler.Calculate))

	log.Fatal(http.ListenAndServe(":"+port, router))
}
