package config

import (
	"github.com/tkanos/gonfig"
	"sync"
)

type configuration struct {
	Mongo *mongo
	Http  *http
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

var instance *configuration
var once sync.Once

func GetInstance() *configuration {
	once.Do(func() {
		instance = &configuration{
			Mongo: createMongoConfig(),
			Http:  createHTTPConfig(),
		}
	})
	return instance
}

func createMongoConfig() *mongo {
	mongo := mongo{}
	err := gonfig.GetConf("resources/persistence.json", &mongo)

	if err != nil {
		panic(err)
	}

	return &mongo
}

func createHTTPConfig() *http {
	http := http{}
	err := gonfig.GetConf("resources/http.json", &http)

	if err != nil {
		panic(err)
	}

	return &http
}
