package sitrep

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/icambridge/gobucket"
)

func BitbucketHook(w http.ResponseWriter, r *http.Request) {
  log.Println("=== START OF BITBUCKET ===")
  r.ParseForm()
  log.Println(r.Form["payload"][0])
  var h gobucket.Hook
  err := json.Unmarshal([]byte(r.Form["payload"][0]), &h)
  if err != nil {
    log.Println(err)
    return
  }
  log.Println(h.User)
  
  log.Println("=== END OF BITBUCKET ===")
  displayPage(w, "about", map[string]interface{}{})
}

