package main

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)


func main() {
	//define router
	r := mux.NewRouter()

	//api endpoints
	r.Handle("/login", http.HandlerFunc(Login)).Methods("POST")
	r.Handle("/usr/info", http.HandlerFunc(UsrInfo)).Methods("GET")
	r.Handle("/usr/password", http.HandlerFunc(UpdateUsrPassword)).Methods("POST")

	//define options
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	//start server
	log.Fatal(http.ListenAndServe(":8080", corsWrapper.Handler(r)))
}
