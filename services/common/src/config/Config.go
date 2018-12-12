package config

import (
	"github.com/tkanos/gonfig"
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
	Airports  int `env:"SERVICE_AIRPORTS"`
	Countries int `env:"SERVICE_COUNTRIES"`
	Runways   int `env:"SERVICE_RUNWAYS"`
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

func loadConfiguration(path string, configuration interface{}) {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		err := gonfig.GetConf(path, &configuration)

		if err != nil {
			panic(err)
		}
	}
}

func createMongoConfig() *mongo {
	mongo := mongo{}

	loadConfiguration("resources/persistence.json", &mongo)

	return &mongo
}

func createHTTPConfig() *http {
	http := http{}

	loadConfiguration("resources/http.json", &http)

	return &http
}

func createServicesConfig() *services {
	services := services{}

	loadConfiguration("resources/services.json", &services)

	return &services
}
