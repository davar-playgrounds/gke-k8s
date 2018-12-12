package helper

import (
	"encoding/json"
	"log"
	"net/http"
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