package sitrep

import (
//  "github.com/icambridge/genkins"
  "log"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  log.Println("Called /")

	displayPage(w, "index", map[string]interface{}{})
}

func About(w http.ResponseWriter, r *http.Request) {
  log.Println("Called /about")
  displayPage(w, "about", map[string]interface{}{})
}
