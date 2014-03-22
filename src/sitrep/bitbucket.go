package sitrep

import (
  "log"
  "net/http"
  "github.com/icambridge/gobucket"  
  "github.com/gorilla/mux"
)


func BitbucketHook(w http.ResponseWriter, r *http.Request) {
  log.Println("=== START OF BITBUCKET ===")
  r.ParseForm()
  log.Println(r.Form["payload"][0])
  payload := []byte(r.Form["payload"][0])
  
  h, err := gobucket.GetHookData(payload)
  
  if err != nil {
    log.Println(err)
    return
  }
  
  gobucket.ProccessHook(h)
  
  log.Println("=== END OF BITBUCKET ===")
  displayPage(w, "about", map[string]interface{}{})
}


func BitbucketListPullRequests(w http.ResponseWriter, r *http.Request) {
  
  params := mux.Vars(r)
  repo := params["repo"]
  c := gobucket.GetClient(bitbucketUsername, bitbucketPassword)
  pr := c.Repository.Get(bitbucketGroup, repo).GetPullRequests()

  log.Println(pr)
  
  displayPage(w, "pullrequests", map[string]interface{}{"prs": pr})
}

func init()  {
  
  gobucket.AddHook(Unapprove{})
  
}

type Unapprove struct {

}

func (u Unapprove) Exec(h *gobucket.Hook) {

  if len(h.Commits) == 0 {
    return
  }

  pr := h.Repository.GetPullRequestForBranch(h.Commits[0].Branch)
  
  if pr == nil {
    return
  }
  
  pr.Unapprove()
}

