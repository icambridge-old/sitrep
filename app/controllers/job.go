package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"sitrep/app/converters"
	"sitrep/app/entities"
	"sitrep/app/services"
	"time"
)

type Job struct {
	*revel.Controller
}

func (c Job) Info() revel.Result {
	jobName := c.Params.Get("jobName")

	revel.TRACE.Printf("%v", jobName)
	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	bitbucket := services.GetBitbucket()
	t0 := time.Now()
	prs, err := bitbucket.PullRequests.GetAll(bitbucketOwner, jobName)

	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	if err != nil {
		revel.TRACE.Printf("%v", err)
	}
	convertedPrs := converters.GobucketToSitRepMultiPrs(prs)

	repoInfo := entities.RepoInfo{PullRequests: convertedPrs}

	jsonData, err := json.Marshal(repoInfo)
	//	json := template.JS(string(jsonData))
	json := string(jsonData)

	return c.Render(json)
}

func (c Job) Build() revel.Result {

	jobName := c.Params.Get("jobName")
	branchName := c.Params.Get("branchName")
	jenkins := services.GetJenkins()

	// TODO add app config for parameters.
	parameters := map[string]string{
		"branchName": branchName,
	}
	jenkins.Builds.TriggerWithParameters(jobName, parameters)

	json := "{\"status\":\"Success\"}"
	return c.Render(json)
}
