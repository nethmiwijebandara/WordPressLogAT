package main

import (
	"fmt"
	"log"
	"net/http"

	"./controller"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/fileupload", controller.FileHandler).
		Methods("POST")
	fmt.Println("This works!")

	r.HandleFunc("/register", controller.RegisterHandler).
		Methods("POST", "OPTIONS")
	r.HandleFunc("/login", controller.LoginHandler).
		Methods("POST", "OPTIONS")
	r.HandleFunc("/profile", controller.ProfileHandler).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))

}
