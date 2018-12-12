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
	router.HandleFunc("/runways", dao.GetRunways).Methods("GET")
	router.HandleFunc("/runways/{id}", dao.GetRunway).Methods("GET")
	router.HandleFunc("/runways/{id}", dao.CreateRunway).Methods("POST")
	router.HandleFunc("/runways/{id}", dao.DeleteRunway).Methods("DELETE")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
