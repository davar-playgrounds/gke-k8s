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

func GetCountries(w http.ResponseWriter, r *http.Request) {
	result := make([]models.Country, 0, 10)

	if err := persistence.GetCollection().Find(nil).All(&result); err != nil {
		helper.PrintErrorMessage(w, 500, "Could not process request")
		panic(err)
	}

	data, err := json.Marshal(&result)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Fatal(err)
	}

	helper.PrintMessage(w, 200, data)
}

func GetCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result := models.Country{}

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Fatal(err)
	}

	if err := persistence.GetCollection().Find(bson.M{"id": i}).One(&result); err != nil {
		helper.PrintErrorMessage(w, 404,"Entry not found")
		log.Fatal(err)
	}

	data, err := json.Marshal(&result)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Fatal(err)
	}

	helper.PrintMessage(w, 200, data)
}

func CreateCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country := models.Country{}

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Fatal(err)
	}

	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input body")
		log.Fatal(err)
	}

	country.ID = int(i)

	if err := persistence.GetCollection().Insert(&country); err != nil {
		helper.PrintErrorMessage(w, 400, "Failed to save data")
		log.Fatal(err)
	}

	data, err := json.Marshal(&country)

	if err != nil {
		helper.PrintErrorMessage(w, 500,"Could not process response")
		log.Fatal(err)
	}

	helper.PrintMessage(w, 200, data)
}

func DeleteCountry(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, err := strconv.ParseInt(params["id"], 10, 0)

	if err != nil {
		helper.PrintErrorMessage(w, 400, "Invalid input")
		log.Fatal(err)
	}

	if err := persistence.GetCollection().Remove(bson.M{"id": i}); err != nil {
		helper.PrintErrorMessage(w, 404,"Entry not found")
		log.Fatal(err)
	}

	helper.PrintMessage(w, 200, []byte("{}"))
}
