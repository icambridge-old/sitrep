package sitrep

import (
  "log"
  "net/http"
  "github.com/icambridge/gobucket"
)

func BitbucketHook(w http.ResponseWriter, r *http.Request) {
  log.Println("=== START OF BITBUCKET ===")
  r.ParseForm()
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


func init()  {
  
  gobucket.AddHook(Unapprove{})
  
}

type Unapprove struct {

}

func (u Unapprove) Exec(h *gobucket.Hook) {
  pr := h.Repository.GetPullRequestForBranch(h.Commits[0].Branch)
  
  if pr == nil {
    return
  }
  
  pr.Unapprove()
}
