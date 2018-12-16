package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"log"
	"net/http"
)

func main() {
	conf := config.GetInstance()

	router := mux.NewRouter()
	fmt.Printf("\nHello, serving Frontend on port :%v", conf.Http.Port)

	router.HandleFunc("/health", helper.AlwaysHealthy).Methods("GET")
	router.HandleFunc("/alive", helper.AlwaysHealthy).Methods("GET")

	helper.ServeFile(router, "/", "resources/static/index.html", "text/html")
	helper.ServeFile(router, "/app.js", "resources/static/app.js", "application/javascript")
	helper.ServeFile(router, "/app.css", "resources/static/app.css", "text/css")
	helper.ServeFile(router, "/reset.css", "resources/static/reset.css", "text/css")

	helper.RouteTraffic(router, conf.Services.Countries, "/countries/search/{query}")

	helper.RouteTraffic(router, conf.Services.Runways, "/runways/airport_ident/{query}")

	helper.RouteTraffic(router, conf.Services.RunwaysCountry, "/runways-country/country_code/{country_code}")
	helper.RouteTraffic(router, conf.Services.RunwaysCountry, "/runways-country/country_code/{country_code}/search/{query}")

	helper.RouteTraffic(router, conf.Services.Airports, "/airports/country_code/{country_code}/search/{query}")

	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
