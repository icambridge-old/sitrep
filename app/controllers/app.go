package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/icambridge/gobucket"
	"sitrep/app/services"
	"sitrep/app/converters"
	"sitrep/app/entities"
	"sitrep/app/models"
	"database/sql"
	"strings"
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

func (c App) All() revel.Result {

	pullRequests := []*gobucket.PullRequest{}
	bitbucket := services.GetBitbucket()
	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")

	repos, err := bitbucket.Repositories.GetAll(bitbucketOwner)

	if err != nil {
		panic(err)
	}

	for _, repo := range repos {
		parts := strings.Split(repo.FullName, "/")
		newRequests, err := bitbucket.PullRequests.GetAll(bitbucketOwner, parts[1])
		if err != nil {
			panic(err)
		}
		pullRequests = append(pullRequests, newRequests...)
	}
	convertedPrs := converters.GobucketToSitRepMultiPrs(pullRequests)

	entity := entities.RepoInfo{PullRequests: convertedPrs}

	jsonData, _ := json.Marshal(entity)
	json := string(jsonData)
	c.Request.Format = "json"
	return c.Render(json)
}


func (c App) getBuildInfoForPullRequests(jobName string, pullRequests []entities.PullRequest) []entities.PullRequest {

	for key, pullRequest := range pullRequests {
		pullRequests[key].LastBuild =  c.getBranchBuild(jobName, pullRequest.Source)
	}

	return pullRequests
}

func (c App) getBuildInfoForBranches(jobName string, branches []entities.Branch) []entities.Branch {

	for key, branch := range branches {
		branches[key].LastBuild = c.getBranchBuild(jobName, branch.Name)
	}

	return branches
}

func (c App) getBranchBuild(jobName string, branchName string) models.Build {
	var build models.Build
	err := c.Txn.SelectOne(&build, `SELECT * FROM builds WHERE application_name = ? AND branch = ? ORDER BY id DESC LIMIT 0,1`, jobName, branchName)

	if err != nil && err == sql.ErrNoRows {
		build.Status = "None"
	}

	if err != nil && err != sql.ErrNoRows {
		revel.ERROR.Println("Branches - %v", err)
	}

	return build
}
