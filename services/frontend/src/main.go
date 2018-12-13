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

func serveIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	http.ServeFile(w, r, "resources/static/index.html")
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")

	http.ServeFile(w, r, "resources/static/app.css")
}

func serveJS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")

	http.ServeFile(w, r, "resources/static/app.js")
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
	router.HandleFunc("/", serveIndex).Methods("GET")
	router.HandleFunc("/app.js", serveJS).Methods("GET")
	router.HandleFunc("/app.css", serveCSS).Methods("GET")

	routeTraffic(router, conf.Services.Countries, "/countries")
	routeTraffic(router, conf.Services.Countries, "/countries/{id}")



	log.Panic(http.ListenAndServe(fmt.Sprintf(":%v", conf.Http.Port), router))
}
