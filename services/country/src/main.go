package main

import (
	"./config"
	"./models"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	router := mux.NewRouter()
	fmt.Printf("\nHello, serving on :%v", conf.Http.Port)
	router.HandleFunc("/country", models.GetCountries).Methods("GET")
	router.HandleFunc("/country/{id}", models.GetCountry).Methods("GET")
	router.HandleFunc("/country/{id}", models.CreateCountry).Methods("POST")
	router.HandleFunc("/country/{id}", models.DeleteCountry).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
