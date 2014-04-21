package hooks

import (
	"sitrep/app/services"
	"github.com/icambridge/gobucket"
	"github.com/revel/revel"
)

type Unapprove struct {
}

func (u Unapprove) Exec(h *gobucket.Hook) {
	if len(h.Commits) < 1 {
		return
	}
	bitbucket := services.GetBitbucket()

	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")

	pr, err := bitbucket.PullRequests.GetBranch(bitbucketOwner, h.Repository.Slug, h.Commits[0].Branch)

	if err != nil {
		// handle error
		revel.TRACE.Println(err)
		return
	}

	if pr == nil {
		return
	}

	err = bitbucket.PullRequests.Unapprove(bitbucketOwner, h.Repository.Slug, pr.Id)
}

type AutoBuild struct {

}

func (a AutoBuild) Exec(h *gobucket.Hook) {
	if len(h.Commits) < 1 {
		return
	}
	bitbucket := services.GetBitbucket()

	bitbucketOwner := revel.Config.StringDefault("bitbucket.owner", "")

	pr, err := bitbucket.PullRequests.GetBranch(bitbucketOwner, h.Repository.Slug, h.Commits[0].Branch)

	if err != nil {
		// handle error
		revel.TRACE.Println(err)
		return
	}

	if pr == nil {
		revel.TRACE.Println("Empty")
		return
	}

	bitbucketUsername := revel.Config.StringDefault("bitbucket.username", "")
	// This isn't my fault, the API won't return all the info when you list them.
	fullPr, err := bitbucket.PullRequests.GetById(bitbucketOwner, h.Repository.Slug, pr.Id)

	for _, reviewer := range fullPr.Reviewers {

		if reviewer.Username == bitbucketUsername {

			jenkins := services.GetJenkins()

			p := map[string]string{
				"branchName": h.Commits[0].Branch,
			}

			jenkins.Builds.TriggerWithParameters(h.Repository.Slug, p)
			return
		}
	}
}

type BranchBuild struct {

}


func (b BranchBuild) Exec(h *gobucket.Hook) {
	if len(h.Commits) < 1 {
		return
	}
	// Todo move to database settings
	if h.Commits[0].Branch != "develop" && h.Commits[0].Branch != "master" && h.Commits[0].Branch != "release" {
		return
	}

	jenkins := services.GetJenkins()

	p := map[string]string{
		"branchName": h.Commits[0].Branch,
	}

	jenkins.Builds.TriggerWithParameters(h.Repository.Slug, p)
}
