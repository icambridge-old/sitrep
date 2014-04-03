package framework

import (
	"fmt"
	"net/http"
	"html/template"
)

type Controller struct {

	request http.Request

	response http.Response

}

func (c *Controller) displayView(view string, data map[string]interface{}) {
	c.response.Header().Set("Content-Type", "text/html; charset=utf-8")
	t :=  template.Must(template.New("*").ParseFiles("templates/base.html", fmt.Sprintf("templates/%s.html", view)))
	data["Section"] = template
	t.ExecuteTemplate(w, "base", data)

}
