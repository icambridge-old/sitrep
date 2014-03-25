package sitrep

import (
	"log"
	"encoding/json"
	"net/http"
	"github.com/icambridge/genkins"
	"github.com/bradfitz/gomemcache/memcache"
	"sitrep/model"
)

func getJenkinsJobs() *memcache.Item {
	keyStr := "jenkins.jobs.all"

	item, err := memClient.Get(keyStr)

	if err != nil {
		log.Println(err)
	}

	if item == nil {
		jobs, err := jenkins.Jobs.GetAll()

		if err != nil {
			log.Printf("Tried to get jobs for jenkins but got %v", err)
		}

		json, err := json.Marshal(jobs.Jobs)

		if err != nil {
			log.Printf("Tried to turn jobs into json but got %v", err)
		}

		item = &memcache.Item{Key: keyStr, Value: json, Expiration: 300}
		memClient.Set(item)
	}

	return item
}



func JenkinsHook(w http.ResponseWriter, r *http.Request) {

	job, err := genkins.GetHook(r)

	if err != nil {
		log.Println(err)
	}

	info, err := jenkins.Builds.GetInfo(job.Build)
	if err != nil {
		log.Println(err)
	}

	branchName := info.GetBranchName()

	b := &model.Build{
		BuildId: job.Build.Number,
		ApplicationName: job.Name,
		Status: job.Build.Status,
		Phase: job.Build.Phase,
		Branch: branchName,
	}

	// TODO seperate out logic
	err = buildModel.Save(b)

	if err != nil {
		log.Println(err)
	}

	owner, _ := cfg.String("bitbucket", owner)

	pr, err := bitbucket.PullRequests.GetBranch(owner, b.ApplicationName, b.Branch)

	if err != nil {
		log.Println(err)
	}

	err = bitbucket.PullRequests.Approve(owner, b.ApplicationName, pr.Id)

	if err != nil {
		log.Println(err)
	}
}
