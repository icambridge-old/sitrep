package main

import (
  "github.com/gorilla/mux"
  "github.com/icambridge/genkins"
  "log"
  "net/http"
  "html/template"
  "fmt"
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
  jobView, _ := genkins.GetJobView()
  /*
  numberOfJobs := len(jobView.Jobs)

	for i := 0; i < numberOfJobs; i++ {
		status := ""
		switch jobView.Jobs[i].Color {
		case "red":
			status = "Failure"
		case "blue", "green":
			status = "Sucess"
		default:
			status = "Unknown"
		}

		log.Println(fmt.Sprintf("Build - %s - %s\r\n", jobView.Jobs[i].Name, status))
	}
 */
  
  displayPage(w, "index", map[string]interface{}{"Jobs" : jobView.Jobs})
}

func about(w http.ResponseWriter, r *http.Request) {
  displayPage(w, "about", map[string]interface{}{})
}

func errorpage(w http.ResponseWriter, r *http.Request) {

  displayPage(w, "error", map[string]interface{}{})
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
