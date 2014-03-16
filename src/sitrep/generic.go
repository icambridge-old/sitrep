package sitrep

import (
  "github.com/icambridge/genkins"
  "log"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  log.Println("Called /")
  jobView, err := genkins.GetJobView()
 
  if err != nil {
    displayError(w)  
    return
  }
  
  displayPage(w, "index", map[string]interface{}{"Jobs" : jobView.Jobs})
}

func About(w http.ResponseWriter, r *http.Request) {
  log.Println("Called /about")
  displayPage(w, "about", map[string]interface{}{})
}
