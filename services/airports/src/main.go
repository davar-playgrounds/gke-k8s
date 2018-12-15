package main

import (
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"./dao"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	router := mux.NewRouter()
	fmt.Printf("\nHello, serving Airport REST API on port :%v", conf.Http.Port)
	router.HandleFunc("/airports", dao.GetAirports).Methods("GET")
	router.HandleFunc("/airports/country_code/{code}", dao.GetAirportsByCountryCode).Methods("GET")
	router.HandleFunc("/airports/country_code/{code}/search/{query}", dao.GetAirportsByCountryCodeAndSearch).Methods("GET")
	router.HandleFunc("/airports/{id}", dao.GetAirport).Methods("GET")
	router.HandleFunc("/airports/{id}", dao.CreateAirport).Methods("POST")
	router.HandleFunc("/airports/{id}", dao.DeleteAirport).Methods("DELETE")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
