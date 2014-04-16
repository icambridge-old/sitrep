package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"sitrep/app/services"
	"html/template"
	"time"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	jenkins := services.GetJenkins()

	t0 := time.Now()
	jobs, err := jenkins.Jobs.GetAll()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
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
