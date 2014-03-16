package sitrep

import (
  "log"
  "encoding/json"
  "net/http"
)

func BitbucketHook(w http.ResponseWriter, r *http.Request) {
  log.Println("=== START OF BITBUCKET ===")
  r.ParseForm()
  log.Println(r.Form["payload"][0])
  var h Hook
  err := json.Unmarshal([]byte(r.Form["payload"][0]), &h)
  if err != nil {
    log.Println(err)
    return
  }
  log.Println(h.User)
  
  log.Println("=== END OF BITBUCKET ===")
  displayPage(w, "about", map[string]interface{}{})
}



type Hook struct {
  Repository BitbucketRepository `json:"repository"`
  Truncated bool `json:"truncated"`
  Commits []BitbucketCommit `json:"commits"`
  ConnonUrl string `json:"canon_url"`
  User string `json:"user"`
}

type BitbucketCommit struct {
  Node string `json:"node"`
  Files []BitbucketFile `json:"files"`
  RawAuthor string `json:"raw_author"`
  UtcTimestamp string `json:"utctimestamp"`
  Author string `json:"author"`
  Timestamp string `json:"timestamp"`
  RawNode string `json:"raw_node"`
  Parents []string `json:"parents"`
  Branch string `json:"branch"`
  Message string `json:"message"`
  Size int `json:"size"`
}

type BitbucketFile struct {
  Type string `json:"type"`
  File string `json:"file"`
}

type BitbucketRepository struct {
	Website     string `json:"website"`
	Fork        bool `json:"fork"`
	Name        string               `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`	 
	Private     bool `json:"is_private"`
}
