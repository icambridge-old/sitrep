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

	t0 := time.Now()
	jobsValue, err := jenkins.Jobs.GetAll()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	if err != nil {
		revel.TRACE.Printf("/ Jenkins jobs - %v", err)
	}
	jobs := jobsValue.Jobs

	return c.Render(jobs)
}
