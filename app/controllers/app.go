package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/icambridge/genkins"
	"html/template"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	jenkins := genkins.NewClient(
		revel.Config.StringDefault("jenkins.hostname", ""),
		revel.Config.StringDefault("jenkins.username", ""),
		revel.Config.StringDefault("jenkins.token", ""),
	)
	jobs, err := jenkins.Jobs.GetAll()

	if err != nil {
		revel.TRACE.Printf("/ Jenkins jobs - %v", err)
	}

	jsonBytes, err := json.Marshal(jobs.Jobs)

	if err != nil {
		revel.TRACE.Printf("/ Jenkins jobs json - %v", err)
	}
	json := template.JS(string(jsonBytes))
	return c.Render(json)
}
