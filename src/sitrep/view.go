package sitrep

import (
	"fmt"
	"html/template"
	"net/http"
)

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
