package gobucket

import (
  "fmt"
  "log"
  "bytes"
  "encoding/json"
)

type RepositoryService struct {
    Client *Client
}

func (rs *RepositoryService) Get(user string, repoName string) *Repository {
  
   url := fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s", user, repoName)
   resp := rs.Client.Request.get(url)
  
	var repo Repository
	body := &bytes.Buffer{}
	_, err := body.ReadFrom(resp.Body)

	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
  
	dec := json.NewDecoder(body)
	err = dec.Decode(&repo)
	
	// TODO move
  repo.Owner = user
  repo.Slug = repoName
  repo.Request = rs.Client.Request
	return &repo
}

type Repository struct {
	Website     string `json:"website"`
	Fork        bool   `json:"fork"`
	Name        string `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`	 
	Private     bool   `json:"is_private"`
	Request     *Request
}


func (r *Repository) GetPullRequestForBranch(branch string) *PullRequest {
  prs := r.GetPullRequests()
  for _, pr := range prs {
		if pr.Source.Branch.Name == branch {
		  pr.Repo = r
	    return &pr
		}	
	}
  
  return nil
}

func (r *Repository) GetPullRequests() []PullRequest {
  
    log.Println(r.Owner)
    log.Println(r.Slug)
	url := fmt.Sprintf("https://bitbucket.org/api/2.0/repositories/%s/%s/pullrequests/", r.Owner, r.Slug)
	resp := r.Request.get(url)

	var pullRequests PullRequestList
	body := &bytes.Buffer{}
	_, err := body.ReadFrom(resp.Body)

	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
  
	dec := json.NewDecoder(body)
	err = dec.Decode(&pullRequests)

	return pullRequests.Values
}
