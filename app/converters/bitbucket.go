package converters

import (
	"github.com/icambridge/gobucket"
	"github.com/revel/revel"
	"sitrep/app/entities"
	"fmt"
	"time"
)

func GobucketToSitRepMultiPrs(prs []*gobucket.PullRequest) []entities.PullRequest {

	output := make([]entities.PullRequest,0)

	for _, pr := range prs {
		newPr := GobucketToSitRepSinglePr(pr)
		output = append(output, newPr)
	}

	return output
}

func GobucketToSitRepSinglePr(pr *gobucket.PullRequest) entities.PullRequest {
	urlStr := fmt.Sprintf("http://bitbucket.org/%s/pull-request/%d", pr.Destination.Repository.FullName, pr.Id)
	return entities.PullRequest{
		Title:       pr.Title,
		Author:      pr.Author.DisplayName,
		Source:      pr.Source.Branch.Name,
		Destination: pr.Destination.Branch.Name,
		Id:          pr.Id,
		Url:         urlStr,
	}
}

func gobucketToSitRepApproval(approvals []gobucket.User) []entities.Approval {
	output := []entities.Approval{}
	for _, approval := range approvals {

		a := entities.Approval{
			Avatar: approval.Links.Avatar.Href,
			Name:   approval.DisplayName,
		}
		output = append(output, a)
	}
	return output
}

func GobucketToSitRepBranches(branches map[string]*gobucket.Branch) []entities.Branch {

	output := []entities.Branch{}

	for branchName, branch := range branches {
		newBranch := GobucketToSitRepBranch(branch)
		newBranch.Name = branchName
		output = append(output, newBranch)
	}

	return output
}

func GobucketToSitRepBranch(branch *gobucket.Branch) entities.Branch {

	layout := "2006-01-02 15:04:05"

	outputBranch := entities.Branch{Name: branch.Branch}

	t, err := time.Parse(layout, branch.Timestamp)

	if err != nil {
		revel.TRACE.Println(err)
		revel.TRACE.Println(branch.Timestamp)
	}

	outputBranch.Timestamp = t.Unix()

	outputBranch.Status = "Unknown"


	return outputBranch
}
