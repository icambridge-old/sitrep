package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"sitrep/app/converters"
	"sitrep/app/services"
	"sitrep/app/models"
	"database/sql"
	"sitrep/app/entities"
	"github.com/icambridge/gobucket"
	"strings"
)

type Job struct {
	GorpController
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

	convertedPrs = c.getBuildInfoForPullRequests(jobName, convertedPrs)
	entity := entities.RepoInfo{PullRequests: convertedPrs}
	jsonData, err := json.Marshal(entity)
	//json := template.JS(string(jsonData))
	json := string(jsonData)
	c.Request.Format = "json"


	return c.Render(json)
}

func (c Job) getBuildInfoForPullRequests(jobName string, pullRequests []entities.PullRequest) []entities.PullRequest {

	for key, pullRequest := range pullRequests {
		pullRequests[key].LastBuild =  c.getBranchBuild(jobName, pullRequest.Source)
	}

	return pullRequests
}

func (c Job) getBuildInfoForBranches(jobName string, branches []entities.Branch) []entities.Branch {

	for key, branch := range branches {
		branches[key].LastBuild = c.getBranchBuild(jobName, branch.Name)
	}

	return branches
}

func (c Job) getBranchBuild(jobName string, branchName string) models.Build {
	var build models.Build
	err := c.Txn.SelectOne(&build, `SELECT * FROM builds WHERE application_name = ? AND branch = ? ORDER BY id DESC`, jobName, branchName)

	if err != nil && err == sql.ErrNoRows {
		build.Status = "None"
	}

	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Branches - %v", err)
	}

	return build
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
	var branches gobucket.BranchList

	if err := cache.Get("branches_"+jobName, &branches); err != nil || branches == nil {
		fmt.Println("Hello world")
		bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
		bitbucket := services.GetBitbucket()

		b, err := bitbucket.Repositories.GetBranches(bitbucketOwner, jobName)

		revel.TRACE.Printf("%v", b)

		if err != nil {
			revel.TRACE.Printf("%v", err)
		}
		go cache.Set("branches_"+jobName, b, cache.DEFAULT)
		branches = b
	}
	convertedBranches := converters.GobucketToSitRepBranches(branches)

	convertedBranches = c.getBuildInfoForBranches(jobName, convertedBranches)
	entity := entities.RepoInfo{Branches: convertedBranches}
	jsonData, _ := json.Marshal(entity)
	//json := template.JS(string(jsonData))
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}
