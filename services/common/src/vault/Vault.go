package vault

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var instance *cache.Cache
var once sync.Once

func GetInstance() *cache.Cache {
	once.Do(func() {
		instance = cache.New(5*time.Minute, 10*time.Minute)
	})
	return instance
}
