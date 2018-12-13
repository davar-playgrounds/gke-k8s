package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func serveFile(router *mux.Router, endpoint string, filePath string, mimeType string) {
	router.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeType)

		http.ServeFile(w, r, filePath)
	}).Methods("GET")
}

func routeTraffic(router *mux.Router, host string, endpoint string) {
	router.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s%s", host, r.RequestURI)

		fmt.Printf("\nUrl: %s\n", url)

		timeout := time.Duration(5) * time.Second
		transport := &http.Transport{
			ResponseHeaderTimeout: timeout,
			Dial: func(network, addr string) (net.Conn, error) {
				return net.DialTimeout(network, addr, timeout)
			},
			DisableKeepAlives: true,
		}
		client := &http.Client{
			Transport: transport,
		}

		resp, err := client.Get(url)
		if err != nil {
			helper.PrintErrorMessage(w, 500, "Routing error; service down")
			log.Print(err)
			return
		}

		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			helper.PrintErrorMessage(w, 500, "Error creating response")
			log.Print(err)
			return
		}

	}).Methods("GET")
}

func main() {
	conf := config.GetInstance()

	router := mux.NewRouter()
	fmt.Printf("\nHello, serving Frontend on port :%v", conf.Http.Port)

	serveFile(router, "/", "resources/static/index.html", "text/html")
	serveFile(router, "/app.js", "resources/static/app.js", "application/javascript")
	serveFile(router, "/app.css", "resources/static/app.css", "text/css")
	serveFile(router, "/reset.css", "resources/static/reset.css", "text/css")

	routeTraffic(router, conf.Services.Countries, "/countries")
	routeTraffic(router, conf.Services.Countries, "/countries/search/{query}")
	routeTraffic(router, conf.Services.Countries, "/countries/{id}")

	routeTraffic(router, conf.Services.Runways, "/runways")
	routeTraffic(router, conf.Services.Runways, "/runways/{id}")

	routeTraffic(router, conf.Services.Airports, "/airports")
	routeTraffic(router, conf.Services.Airports, "/airports/{id}")

	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
