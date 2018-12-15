package joins

import (
	"github.com/patrickmn/go-cache"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhaddon/gke-k8s/services/common/src/helper"
	"github.com/mhaddon/gke-k8s/services/common/src/vault"
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"github.com/mhaddon/gke-k8s/services/common/src/models"
	"io/ioutil"
	"log"
	"net/http"
)

func getRunwaysFromAirportRef(airport_ident string) (*[]models.Runway, error) {
	conf := config.GetInstance()

	runways := make([]models.Runway, 0, 10)

	resp, err := helper.QueryEndpoint(fmt.Sprintf("%s%s%s", conf.Services.Runways, "/runways/airport_ident/", airport_ident))
	if err != nil {
		return &runways, errors.New("could not query endpoint")
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return &runways, errors.New("could not read message")
	}

	if err = json.Unmarshal(jsonData, &runways); err != nil {
		return &runways, errors.New("could not parse message")
	}

	return &runways, nil
}

func getAirportsByCountryCode(country_code string, query string) (*[]models.Airport, error) {
	conf := config.GetInstance()

	airports := make([]models.Airport, 0, 10)

	var endpoint = fmt.Sprintf("%s%s%s", conf.Services.Airports, "/airports/country_code/", country_code)

	if len(query) > 0 {
		endpoint = fmt.Sprintf("%s%s%s%s%s", conf.Services.Airports, "/airports/country_code/", country_code, "/search/", query)
	}

	resp, err := helper.QueryEndpoint(endpoint)
	if err != nil {
		return &airports, errors.New("could not query endpoint")
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return &airports, errors.New("could not read message")
	}

	if err = json.Unmarshal(jsonData, &airports); err != nil {
		return &airports, errors.New("could not parse message")
	}

	return &airports, nil
}

func getRunwaysByCountryCode(country_code string, query string) (*[]models.Runway, error) {
	runways := make([]models.Runway, 0, 10)

	airports, err := getAirportsByCountryCode(country_code, query)
	if err != nil {
		return &runways, err
	}

	for _, element := range *airports {
		newRunways, err := getRunwaysFromAirportRef(element.Ident)

		if err != nil {
			return &runways, err
		}

		runways = append(runways, *newRunways...)
	}

	return &runways, nil
}

func GetRunwaysByCountryCodeCached(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	country_code := params["country_code"]
	query := ""

	if value, ok := params["query"]; ok {
		query = value
	}

	cache_name := fmt.Sprintf("GetRunwaysByCountryCode-%s-%s", country_code, query)
	cached_data, found := vault.GetInstance().Get(cache_name)

	if found {
		helper.PrintMessage(w, 200, []byte(cached_data.(string)))
		return
	}

	runways, err := getRunwaysByCountryCode(country_code, query)

	if err != nil {
		helper.PrintErrorMessage(w, 500, err.Error())
		log.Print(err)
		return
	}

	data, err := json.Marshal(&runways)

	if err != nil {
		helper.PrintErrorMessage(w, 500, "Could not process response")
		log.Print(err)
		return
	}

	vault.GetInstance().Set(cache_name, string(data), cache.DefaultExpiration)

	helper.PrintMessage(w, 200, data)
}