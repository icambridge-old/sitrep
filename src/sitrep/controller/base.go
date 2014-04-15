package controller

import (
	"github.com/icambridge/framework"
	"github.com/icambridge/genkins"
	"github.com/icambridge/gobucket"
	"github.com/robfig/config"
	"github.com/bradfitz/gomemcache/memcache"
	_ "github.com/go-sql-driver/mysql"
)

type Base struct {
	framework.Controller

	Jenkins *genkins.Client

	Bitbucket *gobucket.Client

	Memcache *memcache.Client

	Config *config.Config
}
