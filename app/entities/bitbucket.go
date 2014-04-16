package entities


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

type RepoInfo struct {
	Name         string        `json:"name"`
	PullRequests []PullRequest `json:"pulls"`
	Branches     []Branch      `json:"branches"`
}
