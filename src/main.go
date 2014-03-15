package main

import (
  "github.com/gorilla/mux"
  "github.com/icambridge/genkins"
  "log"
  "net/http"
  "html/template"
  "fmt"
  "bytes"
)

func main() {
  rtr := mux.NewRouter()
  rtr.HandleFunc("/", index).Methods("GET")
  rtr.HandleFunc("/about", about).Methods("GET")
  rtr.HandleFunc("/error", errorpage).Methods("GET")

  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":3000", nil)
}


func index(w http.ResponseWriter, r *http.Request) {
  jobView, err := genkins.GetJobView()
 
  if err != nil {
    displayError(w)  
    return
  }
  
  displayPage(w, "index", map[string]interface{}{"Jobs" : jobView.Jobs})
}

func about(w http.ResponseWriter, r *http.Request) {
  displayPage(w, "about", map[string]interface{}{})
}

func errorpage(w http.ResponseWriter, r *http.Request) {

    displayError(w)  
}

func displayError(w http.ResponseWriter) {
     http.Error(w, http.StatusText(500), 500)
}

func displayPage(w http.ResponseWriter, template string, data map[string]interface{}) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  t := newTemplate("templates/base.html", fmt.Sprintf("templates/%s.html", template))
  data["Section"] = template
  t.ExecuteTemplate(w, "base", data)
}

func newTemplate(files ...string) *template.Template {
	return template.Must(template.New("*").ParseFiles(files...))
}
