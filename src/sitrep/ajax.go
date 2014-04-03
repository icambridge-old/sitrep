package sitrep

import (
	"fmt"
	"github.com/icambridge/gobucket"
	"log"
)

type PullRequestList struct {
	Values []PullRequest `json:"values"`
}

type PullRequest struct {
	Title       string     `json:"title"`
	Author      string     `json:"author"`
	Source      string     `json:"source"`
	Destination string     `json:"destination"`
	Id          int        `json:"id"`
	Approvals   []Approval `json:"approvals"`
	Url         string     `json:"url"`
}

type Approval struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

type Branch struct {
	Name    string `json:"name"`
	BuildId int    `json:"build_id"`
	Status  string `json:"status"`
}

func gobucketToSitRepSinglePr(pr *gobucket.PullRequest) PullRequest {
	fullPr, err := bitbucket.PullRequests.GetById(pr.GetOwner(), pr.GetRepoName(), pr.Id)

	if err != nil {
		log.Println(err)
	}
	rawApprovals := fullPr.GetApprovals()

	approvals := gobucketToSitRepApproval(rawApprovals)

	urlStr := fmt.Sprintf("http://bitbucket.org/%s/pull-request/%d", pr.Destination.Repository.FullName, pr.Id)
	return PullRequest{
		Title:       pr.Title,
		Author:      pr.Author.DisplayName,
		Source:      pr.Source.Branch.Name,
		Destination: pr.Destination.Branch.Name,
		Id:          pr.Id,
		Approvals:   approvals,
		Url:         urlStr,
	}
}

func gobucketToSitRepMultiPrs(prs []*gobucket.PullRequest) []PullRequest {

	output := []PullRequest{}

	for _, pr := range prs {
		newPr := gobucketToSitRepSinglePr(pr)
		output = append(output, newPr)
	}

	return output
}

func gobucketToSitRepApproval(approvals []gobucket.User) []Approval {
	output := []Approval{}
	for _, approval := range approvals {

		a := Approval{
			Avatar: approval.Links.Avatar.Href,
			Name:   approval.DisplayName,
		}
		output = append(output, a)
	}
	return output
}

func gobucketToSitRepBranch(applicationName string, branch *gobucket.Branch) Branch {

	build, err := buildModel.GetByApplicationNameAndBranch(applicationName, branch.Branch)

	outputBranch := Branch{Name: branch.Branch}

	if err == nil {
		outputBranch.Status = build.Status
		outputBranch.BuildId = build.BuildId
	} else {
		outputBranch.Status = "Unknown"
	}

	return outputBranch
}

func gobucketToSitRepBranches(applicationName string, branches map[string]*gobucket.Branch) []Branch {

	output := []Branch{}

	for branchName, branch := range branches {
		newBranch := gobucketToSitRepBranch(applicationName, branch)
		newBranch.Name = branchName
		output = append(output, newBranch)
	}

	return output
}

type RepoInfo struct {
	Name         string        `json:"name"`
	PullRequests []PullRequest `json:"pulls"`
	Branches     []Branch      `json:"branches"`
}
