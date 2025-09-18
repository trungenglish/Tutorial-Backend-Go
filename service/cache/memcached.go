package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

var Client *memcache.Client

func InitCache(address string) {
	Client = memcache.New(address)
}
