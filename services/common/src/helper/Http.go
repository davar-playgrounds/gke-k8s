package helper

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func PrintErrorMessage(w http.ResponseWriter, responseCode int, response string) {
	message := map[string]interface{}{ "err": response, "code": responseCode }
	encodedMessage, _ := json.Marshal(message)
	PrintMessage(w, responseCode, encodedMessage)
}

func PrintMessage(w http.ResponseWriter, responseCode int, response[] byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	if _, err := w.Write(response); err != nil {
		log.Panic(err)
	}
}

func ServeFile(router *mux.Router, endpoint string, filePath string, mimeType string) {
	router.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mimeType)

		http.ServeFile(w, r, filePath)
	}).Methods("GET")
}

func QueryEndpoint(url string) (*http.Response, error) {
	timeout := time.Duration(2) * time.Minute
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
		return nil, err
	}

	return resp, nil
}

func RouteTraffic(router *mux.Router, host string, endpoint string) {
	router.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		url := fmt.Sprintf("%s%s", host, r.RequestURI)

		resp, err := QueryEndpoint(url)

		if err != nil {
			PrintErrorMessage(w, 500, "Routing error; service down")
			log.Print(err)
		}

		defer resp.Body.Close()

		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))
		w.WriteHeader(resp.StatusCode)

		if _, err := io.Copy(w, resp.Body); err != nil {
			PrintErrorMessage(w, 500, "Error creating response")
			log.Print(err)
			return
		}
	}).Methods("GET")
}