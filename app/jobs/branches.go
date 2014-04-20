package jobs

import (
	"fmt"
	"strings"
	"github.com/revel/revel"
	"github.com/revel/revel/modules/jobs/app/jobs"
	"github.com/revel/revel/cache"
	"sitrep/app/services"
)

type Branches struct {

}

func (j Branches) Run() {

	jenkins := services.GetJenkins()

	jobsValue, err := jenkins.Jobs.GetAll()
	if err != nil {
		revel.TRACE.Printf("/ Jenkins jobs - %v", err)
	}
	jobList := jobsValue.Jobs

	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")
	for _, job := range jobList {
		jobName := strings.ToLower(job.Name)
		bitbucket := services.GetBitbucket()

		branches, _ := bitbucket.Repositories.GetBranches(bitbucketOwner, jobName)

		go cache.Set("branches_"+jobName, branches, cache.DEFAULT)
		output := fmt.Sprintf("Got branches for %s", jobName)
		revel.TRACE.Println(output)
	}
}

func init() {
	revel.OnAppStart(func() {
		jobs.Schedule("*/5 * * * * ?",  Branches{})
	})
}
