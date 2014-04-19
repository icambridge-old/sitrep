package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"sitrep/app/converters"
	"sitrep/app/services"
	"sitrep/app/entities"
	"strings"
)

type Job struct {
	*revel.Controller
}

func (c Job) Info() revel.Result {
	jobName := c.Params.Get("jobName")

	revel.TRACE.Printf("%v", jobName)
	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	bitbucket := services.GetBitbucket()

	prs, err := bitbucket.PullRequests.GetAll(bitbucketOwner, jobName)

	if err != nil {
		revel.TRACE.Printf("%v", err)
	}
	convertedPrs := converters.GobucketToSitRepMultiPrs(prs)
	entity := entities.RepoInfo{PullRequests: convertedPrs}
	jsonData, err := json.Marshal(entity)
	//json := template.JS(string(jsonData))
	json := string(jsonData)
	c.Request.Format = "json"
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

func (c Job) Branches() revel.Result {
	jobName := c.Params.Get("jobName")
	jobName = strings.ToLower(jobName)
	revel.TRACE.Printf("%v", jobName)
	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	bitbucket := services.GetBitbucket()

	branches, err := bitbucket.Repositories.GetBranches(bitbucketOwner, jobName)

	revel.TRACE.Printf("%v", branches["error"])

	if err != nil {
		revel.TRACE.Printf("%v", err)
	}
	convertedPrs := converters.GobucketToSitRepBranches(branches)
	entity := entities.RepoInfo{Branches: convertedPrs}
	jsonData, err := json.Marshal(entity)
	//json := template.JS(string(jsonData))
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}
