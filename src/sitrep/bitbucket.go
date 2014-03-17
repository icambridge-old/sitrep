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
  log.Println(h.User)
  
  log.Println("=== END OF BITBUCKET ===")
  displayPage(w, "about", map[string]interface{}{})
}

