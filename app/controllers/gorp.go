package controllers

import (
	"database/sql"
	"fmt"
	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
)

var (
	Dbm *gorp.DbMap
)

func InitDB() {


	mysqlUsername := revel.Config.StringDefault("mysql.username", "")
	mysqlPassword := revel.Config.StringDefault("mysql.password", "")
	mysqlDatabase := revel.Config.StringDefault("mysql.database", "")
	mysqlDsn := fmt.Sprintf("%s:%s@/%s", mysqlUsername, mysqlPassword, mysqlDatabase)
	db, _ := sql.Open("mysql", mysqlDsn)
	Dbm = &gorp.DbMap{Db:db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
