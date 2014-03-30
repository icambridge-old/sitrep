package sitrep

import (
	"github.com/icambridge/gobucket"
	"github.com/icambridge/genkins"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/robfig/config"
    "database/sql"
	"sitrep/model"
	 _ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	bitbucket *gobucket.Client
	memClient *memcache.Client
	cfg       *config.Config
	jenkins   *genkins.Client
	db        *sql.DB
	buildModel model.BuildModel
	hookProcessors *gobucket.HookObserver
)

func init() {

	hookProcessors = &gobucket.HookObserver{}
	hookProcessors.Add(&Unapprove{})

	cfg, _ = config.ReadDefault("config.cfg")

	bitbucketUser, _ := cfg.String("bitbucket", "username")
	bitbucketPass, _ := cfg.String("bitbucket", "password")

	jenkinsUser, _ := cfg.String("jenkins", "username")
	jenkinsHost, _   := cfg.String("jenkins", "hostname")
	jenkinsToken,_   := cfg.String("jenkins", "token")


	mysqlUsername, _ := cfg.String("mysql", "username")
	mysqlPassword, _ := cfg.String("mysql", "password")
	mysqlDatabase, _ := cfg.String("mysql", "database")
	mysqlDsn := fmt.Sprintf("%s:%s@/%s", mysqlUsername, mysqlPassword, mysqlDatabase)

	jenkins   = genkins.NewClient(jenkinsHost, jenkinsUser, jenkinsToken)
	bitbucket = gobucket.NewClient(bitbucketUser, bitbucketPass)
	memClient = memcache.New("127.0.0.1:11211")

	db, _ = sql.Open("mysql", mysqlDsn)

	buildModel = model.BuildModel{Db: db}
}
