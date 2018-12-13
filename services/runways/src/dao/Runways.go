package dao

import (
	"github.com/mhaddon/gke-k8s/services/common/src/persistence"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"github.com/mhaddon/gke-k8s/services/common/src/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"strconv"
)

func GetRunways(w http.ResponseWriter, r *http.Request) {
	result := make([]models.Runway, 0, 10)

	if err := persistence.GetCollection().Find(nil).All(&result); err != nil {
		helper.PrintErrorMessage(w, 500, "Could not process request")
		log.Print(err)
		return
	}

	data, err := json.Marshal(&result)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Print(err)
		return
	}

	helper.PrintMessage(w, 200, data)
}

func GetRunway(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result := models.Runway{}

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Print(err)
		return
	}

	if err := persistence.GetCollection().Find(bson.M{"id": i}).One(&result); err != nil {
		helper.PrintErrorMessage(w, 404,"Entry not found")
		log.Print(err)
		return
	}

	data, err := json.Marshal(&result)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Print(err)
		return
	}

	helper.PrintMessage(w, 200, data)
}

func CreateRunway(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	runway := models.Runway{}

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Print(err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&runway); err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input body")
		log.Print(err)
		return
	}

	runway.ID = int(i)

	if err := persistence.GetCollection().Insert(&runway); err != nil {
		helper.PrintErrorMessage(w, 400, "Failed to save data")
		log.Print(err)
		return
	}

	data, err := json.Marshal(&runway)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Print(err)
		return
	}

	helper.PrintMessage(w, 200, data)
}

func DeleteRunway(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Print(err)
		return
	}

	if err := persistence.GetCollection().Remove(bson.M{"id": i}); err != nil {
		helper.PrintErrorMessage(w, 404,"Entry not found")
		log.Print(err)
		return
	}

	helper.PrintMessage(w, 200, []byte("{}"))
}
