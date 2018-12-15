package config

import (
	"github.com/tkanos/gonfig"
	"log"
	"os"
	"sync"
)

type configuration struct {
	Mongo    *mongo
	Http     *http
	Services *services
}

type mongo struct {
	Domain     string `env:"MONGO_DOMAIN"`
	Port       int    `env:"MONGO_PORT"`
	DB         string `env:"MONGO_DB"`
	Collection string `env:"MONGO_COLLECTION"`
	Username   string `env:"MONGO_USERNAME"`
	Password   string `env:"MONGO_PASSWORD"`
}

type http struct {
	Port int `env:"HTTP_PORT"`
}

type services struct {
	Airports       string `env:"SERVICE_AIRPORTS"`
	Countries      string `env:"SERVICE_COUNTRIES"`
	Runways        string `env:"SERVICE_RUNWAYS"`
	RunwaysCountry string `env:"SERVICE_RUNWAYSCOUNTRY"`
}

var instance *configuration
var once sync.Once

func GetInstance() *configuration {
	once.Do(func() {
		instance = &configuration{
			Mongo:    createMongoConfig(),
			Http:     createHTTPConfig(),
			Services: createServicesConfig(),
		}
	})
	return instance
}

func createMongoConfig() *mongo {
	mongo := mongo{}

	path := "resources/persistence.json"

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := gonfig.GetConf(path, &mongo)

		if err != nil {
			log.Print(err)
		}
	}

	return &mongo
}

func createHTTPConfig() *http {
	http := http{}

	path := "resources/http.json"

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := gonfig.GetConf(path, &http)

		if err != nil {
			log.Print(err)
		}
	}

	return &http
}

func createServicesConfig() *services {
	services := services{}

	path := "resources/services.json"

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := gonfig.GetConf(path, &services)

		if err != nil {
			log.Print(err)
		}
	}

	return &services
}
