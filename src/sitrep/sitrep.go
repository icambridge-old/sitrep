package sitrep

import (
	"github.com/icambridge/gobucket"
	"github.com/icambridge/genkins"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/robfig/config"
)

var (
	bitbucket *gobucket.Client
	memClient *memcache.Client
	cfg    *config.Config
	jenkins   *genkins.Client
)

func init() {
	cfg, _ = config.ReadDefault("config.cfg")

	bitbucketUser, _ := cfg.String("bitbucket", "username")
	bitbucketPass, _ := cfg.String("bitbucket", "password")

	jenkinsHost, _   := cfg.String("jenkins", "hostname")
	jenkinsToken,_   := cfg.String("jenkins", "token")

	jenkins   = genkins.NewClient(jenkinsHost, jenkinsToken)
	bitbucket = gobucket.NewClient(bitbucketUser, bitbucketPass)
	memClient = memcache.New("127.0.0.1:11211")
}
