package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/icambridge/genkins"
	"sitrep/app/services"
	"sitrep/app/models"
)

type Build struct {
	GorpController
}

func (c Build) Report() revel.Result {
	job, err := genkins.GetHook(c.Request.Request)

	if err != nil {
		fmt.Println(err)
	}

	if job.Build.Phase != "FINISHED" {
		return c.Render()
	}

	jenkins := services.GetJenkins()
	info, err := jenkins.Builds.GetInfo(&job.Build)
	if err != nil {
		fmt.Println(err)
	}

	branchName := info.GetBranchName()

	b := models.Build{
		BuildId:         job.Build.Number,
		ApplicationName: job.Name,
		Status:          job.Build.Status,
		Phase:           job.Build.Phase,
		Branch:          branchName,
	}
	err = c.Txn.Insert(&b)

	if err != nil {
		fmt.Println(err)
	}
	bitbucket := services.GetBitbucket()

	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	pr, err := bitbucket.PullRequests.GetBranch(bitbucketOwner, b.ApplicationName, b.Branch)

	if err != nil {
		fmt.Println(err)
	}

	if pr == nil {
		return c.Render()
	}

	if b.Status == "SUCCESS" {
		err = bitbucket.PullRequests.Approve(bitbucketOwner, b.ApplicationName, pr.Id)
		if err != nil {
			fmt.Println(err)
		}
	} else if b.Status == "FAILURE" {
		err = bitbucket.PullRequests.Unapprove(bitbucketOwner, b.ApplicationName, pr.Id)
		if err != nil {
			fmt.Println(err)
		}
	}

	c.Request.Format = "json"
	return c.Render()
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


func (c Build) Start() revel.Result {

	jobName := c.Params.Get("jobName")
	branch := c.Params.Get("branchName")
	fmt.Println(branch)
	fmt.Println(jobName)
	jenkins := services.GetJenkins()

	p := map[string]string{
		"branchName": branch,
	}

	jenkins.Builds.TriggerWithParameters(jobName, p)
	c.Request.Format = "json"
	return c.Render()
}
