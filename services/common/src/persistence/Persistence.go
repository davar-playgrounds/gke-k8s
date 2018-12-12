package persistence

import (
	"github.com/mhaddon/gke-k8s/services/common/src/config"
	"fmt"
	"gopkg.in/mgo.v2"
	"log"
	"sync"
	"time"
)

var instance *mgo.Session
var once sync.Once

func GetInstance() *mgo.Session {
	once.Do(func() {
		instance = connect()
	})
	return instance
}

func GetDatabase() *mgo.Database {
	conf := config.GetInstance()

	return GetInstance().DB(conf.Mongo.DB)
}

func GetCollection() *mgo.Collection {
	conf := config.GetInstance()

	return GetDatabase().C(conf.Mongo.DB)
}

func connect() *mgo.Session {
	conf := config.GetInstance()

	dialInfo := &mgo.DialInfo{
		Username: conf.Mongo.Username,
		Password: conf.Mongo.Password,
		Source: "admin",
		Addrs: []string{fmt.Sprintf("%s:%v", conf.Mongo.Domain, conf.Mongo.Port)},
		Timeout: 60 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}

	return session
}
