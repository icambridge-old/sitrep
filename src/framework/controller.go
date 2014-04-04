package framework

import (
	"fmt"
	"net/http"
	"html/template"
)

type Controller struct {

	request *http.Request

	response http.ResponseWriter

}
