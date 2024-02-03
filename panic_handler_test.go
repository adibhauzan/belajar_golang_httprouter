package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testing"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, "Panic :", i)
	}
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Uppss")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
