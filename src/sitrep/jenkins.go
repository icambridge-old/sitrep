package sitrep

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/icambridge/genkins"
	"log"
	"net/http"
	"sitrep/model"
	"strings"
)

func JenkinsBuild(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	repo := strings.ToLower(params["repo"])
	branch := params["branch"]

	p := map[string]string{
		"branchName": branch,
	}

	jenkins.Builds.TriggerWithParameters(repo, p)

	fmt.Fprint(w, "{\"status\":\"Success\"}")
}

func JenkinsHook(w http.ResponseWriter, r *http.Request) {

	job, err := genkins.GetHook(r)

	if err != nil {
		log.Println(err)
	}

	info, err := jenkins.Builds.GetInfo(&job.Build)
	if err != nil {
		log.Println(err)
	}

	branchName := info.GetBranchName()

	b := &model.Build{
		BuildId:         job.Build.Number,
		ApplicationName: job.Name,
		Status:          job.Build.Status,
		Phase:           job.Build.Phase,
		Branch:          branchName,
	}

	if b.Phase != "FINISHED" {
		return
	}

	// TODO seperate out logic
	err = buildModel.Save(b)

	if err != nil {
		log.Println(err)
	}
	owner, _ := cfg.String("bitbucket", "owner")

	pr, err := bitbucket.PullRequests.GetBranch(owner, b.ApplicationName, b.Branch)

	if err != nil {
		log.Println(err)
	}

	if pr == nil {
		return
	}

	if b.Status == "SUCCESS" {
		err = bitbucket.PullRequests.Approve(owner, b.ApplicationName, pr.Id)
		if err != nil {
			log.Println(err)
		}
	} else if b.Status == "FAILURE" {
		err = bitbucket.PullRequests.Unapprove(owner, b.ApplicationName, pr.Id)
		if err != nil {
			log.Println(err)
		}
	}
}
