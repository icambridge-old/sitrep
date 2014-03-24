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
	Source string `json:"source"`
	Destination string `json:"destination"`
	Id int `json:"id"`
}

func gobucketToSitRepSinglePr(pr *gobucket.PullRequest) PullRequest {
	return PullRequest{Name: pr.Title, Author: pr.Author.DisplayName, Source: pr.Source.Branch.Name, Destination: pr.Destination.Branch.Name, Id: pr.Id}
}

func gobucketToSitRepMultiPrs(prs []*gobucket.PullRequest) []PullRequest {

	output := []PullRequest{}

	for _, pr := range prs {
		newPr := gobucketToSitRepSinglePr(pr)
		output = append(output, newPr)
	}

	return output
}

type RepoInfo struct {
	Name string `json:"name"`
	PullRequests []PullRequest `json:"pulls"`
}
