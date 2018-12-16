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
	fmt.Printf("\nHello, serving Runway REST API on port :%v", conf.Http.Port)

	router.HandleFunc("/health", helper.IsDBHealthy).Methods("GET")
	router.HandleFunc("/alive", helper.IsDBHealthy).Methods("GET")

	router.HandleFunc("/runways", dao.GetRunways).Methods("GET")
	router.HandleFunc("/runways/airport_ident/{airport_ident}", dao.GetRunwaysByAirportIdent).Methods("GET")
	router.HandleFunc("/runways/{id}", dao.GetRunway).Methods("GET")
	router.HandleFunc("/runways/{id}", dao.CreateRunway).Methods("POST")
	router.HandleFunc("/runways/{id}", dao.DeleteRunway).Methods("DELETE")
	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
