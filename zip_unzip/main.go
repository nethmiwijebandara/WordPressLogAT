package main

import (
	"./controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/fileupload", controller.FileHandler).
		Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))

}