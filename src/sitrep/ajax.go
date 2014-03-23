package sitrep

import (
	"github.com/icambridge/gobucket"
)

type PullRequestList struct {
	Values []PullRequest `json:"values"`
}


type PullRequest struct {
	Name string `json:"name"`
	Author string `json:"author"`
}

func gobucketToSitRepSinglePr(pr *gobucket.PullRequest) PullRequest {
	return PullRequest{Name: pr.Title, Author: pr.Author.DisplayName}
}

func gobucketToSitRepMultiPrs(prs []*gobucket.PullRequest) []PullRequest {

	output := []PullRequest{}

	for _, pr := range prs {
		newPr := gobucketToSitRepSinglePr(pr)
		output = append(output, newPr)
	}

	return output
}
