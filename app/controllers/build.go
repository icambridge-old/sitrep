package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/revel/revel"
	"sitrep/app/models"
)

type Build struct {
	GorpController
}

func (c Build) List() revel.Result {
	var builds []models.Build
	_, err := c.Txn.Select(&builds, `SELECT * FROM builds ORDER BY id DESC LIMIT 0,10`)

	if err != nil {
		panic(err)
	}

	fmt.Println(len(builds))
	jsonData, err := json.Marshal(builds)
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}
