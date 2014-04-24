package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/icambridge/genkins"
	"sitrep/app/services"
	"sitrep/app/models"
	"time"
)

type Build struct {
	GorpController
}

func (c Build) Report() revel.Result {
	job, err := genkins.GetHook(c.Request.Request)

	if err != nil {
		revel.ERROR.Println(err)
	}

	if job.Build.Phase != "FINISHED" {
		return c.Render()
	}

	jenkins := services.GetJenkins()
	info, err := jenkins.Builds.GetInfo(&job.Build)
	if err != nil {
		revel.ERROR.Println(err)
	}
	t := time.Now()
	branchName := info.GetBranchName()

	b := models.Build{
		BuildId:         job.Build.Number,
		ApplicationName: job.Name,
		Status:          job.Build.Status,
		Phase:           job.Build.Phase,
		Branch:          branchName,
		DoneAt:          t.Format("2006-01-02 15:04:05"),
	}
	err = c.Txn.Insert(&b)

	if err != nil {
		revel.ERROR.Println(err)
	}
	bitbucket := services.GetBitbucket()

	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	pr, err := bitbucket.PullRequests.GetBranch(bitbucketOwner, b.ApplicationName, b.Branch)

	if err != nil {
		revel.ERROR.Println(err)
	}

	if pr == nil {
		return c.Render()
	}

	if b.Status == "SUCCESS" {
		err = bitbucket.PullRequests.Approve(bitbucketOwner, b.ApplicationName, pr.Id)
		if err != nil {
			revel.ERROR.Println(err)
		}
	} else if b.Status == "FAILURE" {
		err = bitbucket.PullRequests.Unapprove(bitbucketOwner, b.ApplicationName, pr.Id)
		if err != nil {
			revel.ERROR.Println(err)
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

	jsonData, err := json.Marshal(builds)
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}


func (c Build) Start() revel.Result {

	jobName := c.Params.Get("jobName")
	branch := c.Params.Get("branchName")

	jenkins := services.GetJenkins()

	p := map[string]string{
		"branchName": branch,
	}

	jenkins.Builds.TriggerWithParameters(jobName, p)
	c.Request.Format = "json"
	return c.Render()
}

func (c Build) Branch() revel.Result {

	branch := c.Params.Get("branchName")
	var builds []models.Build
	_, err := c.Txn.Select(&builds, `SELECT * FROM builds WHERE branch = ? GROUP BY application_name ORDER BY id DESC LIMIT 0,10`, branch)

	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(builds)
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}
