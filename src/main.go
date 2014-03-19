package main

import (
  "github.com/gorilla/mux"
  "github.com/icambridge/gobucket"
  "log"
  "net/http"
  "sitrep"
)

func main() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/", sitrep.Index).Methods("GET")
  rtr.HandleFunc("/pullrequests/{repo}", sitrep.BitbucketListPullRequests).Methods("GET")
  rtr.HandleFunc("/about", sitrep.About).Methods("GET")
  rtr.HandleFunc("/bitbucket", sitrep.BitbucketHook).Methods("POST")
  rtr.HandleFunc("/jenkins", sitrep.JenkinsHook).Methods("POST")

  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":47624", nil)
}

func init() {
  gobucket.SetRequest(bitbucketUsername, bitbucketPassword)
}
