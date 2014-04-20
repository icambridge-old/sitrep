package hooks

import (
	"fmt"
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

	fmt.Println(bitbucketOwner)
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
