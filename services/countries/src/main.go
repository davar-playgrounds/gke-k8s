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
	fmt.Printf("\nHello, serving Country REST API on port :%v", conf.Http.Port)

	router.HandleFunc("/health", helper.IsDBHealthy).Methods("GET")
	router.HandleFunc("/alive", helper.IsDBHealthy).Methods("GET")

	router.HandleFunc("/countries", dao.GetCountries).Methods("GET")
	router.HandleFunc("/countries/search/{query}", dao.SearchCountries).Methods("GET")
	router.HandleFunc("/countries/{id}", dao.GetCountry).Methods("GET")
	router.HandleFunc("/countries/{id}", dao.CreateCountry).Methods("POST")
	router.HandleFunc("/countries/{id}", dao.DeleteCountry).Methods("DELETE")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
