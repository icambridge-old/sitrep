package controllers

import (
	"github.com/revel/revel"
	"sitrep/app/services"
	"time"
	"fmt"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {

	jenkins := services.GetJenkins()

	jobsValue, err := jenkins.Jobs.GetAll()
	if err != nil {
		revel.TRACE.Printf("/ Jenkins jobs - %v", err)
	}
	jobs := jobsValue.Jobs

	return c.Render(jobs)
}
