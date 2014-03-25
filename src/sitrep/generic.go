package sitrep

import (
  "log"
  "net/http"
  "html/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
 	log.Println("Called /")
	jobsJson := string(getJenkinsJobs().Value)

	displayPage(w, "index", map[string]interface{}{"jobs": template.JS(jobsJson)})
}

func About(w http.ResponseWriter, r *http.Request) {
  log.Println("Called /about")
  displayPage(w, "about", map[string]interface{}{})
}
