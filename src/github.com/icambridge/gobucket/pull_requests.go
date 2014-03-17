package gobucket

import (

  "fmt"
)

type PullRequestList struct {
	PageLen int `json:"pagelen"`
	Values  []PullRequest `json:"values"`
	Page    int `json:"page"`
	Size    int `json:"size"`
}

type PullRequest struct {
	Description       string           `json:"description"`
	Links             PullRequestLinks `json:"links"`
	Author            BitbucketAuthor  `json:"author"`
	CloseSourceBranch bool             `json:"close_source_branch"`
	Destination       PlaceInfo        `json:"destination"`
	Source            PlaceInfo        `json:"source"`
	Title             string           `json:"title"`
	State             string           `json:"state"`
	CreatedOn         string           `json:"created_on"`
	UpdatedOn         string           `json:"updated_on"`
	Id                int              `json:"id"`
	Repo              *Repository       
}

func (p PullRequest) Approve() {
  urlValue := fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests/%d/approve", p.Repo.Owner, p.Repo.Slug, p.Id)
	p.Repo.Request.post(urlValue)
}

func (p PullRequest) Unapprove() {
  urlValue := fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests/%d/approve", p.Repo.Owner, p.Repo.Slug, p.Id)

	p.Repo.Request.delete(urlValue)
}


type PullRequestLinks struct {
	Decline  BitbucketLink `json:"decline"`
	Commits  BitbucketLink `json:"commits"`
	Self     BitbucketLink `json:"self"`
	Comments BitbucketLink `json:"comments"`
	Merge    BitbucketLink `json:"merge"`
	Html     BitbucketLink `json:"html"`
	Activity BitbucketLink `json:"activity"`
	Diff     BitbucketLink `json:"diff"`
	Approve  BitbucketLink `json:"approve"`
}
