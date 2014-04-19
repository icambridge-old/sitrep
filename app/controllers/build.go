package controllers

import (
	"github.com/revel/revel"
)

type Build struct {
	*revel.Controller
}

func (c Build) List() revel.Result {

	c.Request.Format = "json"
	return c.Render()
}
