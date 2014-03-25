package main

import (

  "github.com/gorilla/mux"
  "log"
  "net/http"
  "sitrep"

)

func main() {


  rtr := mux.NewRouter()
  rtr.HandleFunc("/", sitrep.Index).Methods("GET")
	rtr.HandleFunc("/ajax/bitbucket/pullrequests/{repo}", sitrep.AjaxBitbucketRepo).Methods("GET")
	rtr.HandleFunc("/jenkins", sitrep.JenkinsHook).Methods("POST")
  rtr.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":47624", nil)
}


