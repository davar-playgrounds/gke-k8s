package main

import (
	"./joins"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	log.Printf("\nHello, serving Runways-Country Service on port :%v", conf.Http.Port)

	router := mux.NewRouter()

	router.HandleFunc("/health", helper.AlwaysHealthy).Methods("GET")
	router.HandleFunc("/alive", helper.AlwaysHealthy).Methods("GET")

	router.HandleFunc("/runways-country/country_code/{country_code}", joins.GetRunwaysByCountryCodeCached).Methods("GET")
	router.HandleFunc("/runways-country/country_code/{country_code}/search/{query}", joins.GetRunwaysByCountryCodeCached).Methods("GET")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
