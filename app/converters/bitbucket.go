package converters

import (
	"github.com/icambridge/gobucket"
	"sitrep/app/entities"
	"fmt"
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
