package sitrep

import (
  "log"
  "encoding/json"
  "net/http"
  "github.com/icambridge/genkins"
)

func JenkinsHook(w http.ResponseWriter, r *http.Request) {
  // move logic to genkins
  log.Println("==== START OF JENKINS ====")
  p := make([]byte, r.ContentLength)    
  _, err := r.Body.Read(p)
  if err != nil {
        log.Println("Unable to unmarshall the JSON request", err);
    return
  }
  var job genkins.HookJob
    err1 := json.Unmarshal(p, &job)
    if err1 == nil {
        log.Println(job)
    } else {
        log.Println("Unable to unmarshall the JSON request", err1);
        return

    }
    
    
  log.Println("==== END OF JENKINS ====")
  displayPage(w, "about", map[string]interface{}{})
}



