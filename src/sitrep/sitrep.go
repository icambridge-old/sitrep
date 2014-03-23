package sitrep

import (
	"github.com/icambridge/gobucket"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	bitbucket *gobucket.Client
	memClient  *memcache.Client
)

func init() {
	bitbucket = gobucket.NewClient("", "")
	memClient = memcache.New("127.0.0.1:11211")
}
