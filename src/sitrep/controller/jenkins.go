package controller

import (
	"sitrep/model"
	"fmt"
)

type Jenkins struct {
	Base
	Model model.BuildModel

}

func (c Jenkins) Build(jobName string, branchName string) string {

	parameters := map[string]string{
		"branchName": branchName,
	}
	c.Jenkins.Builds.TriggerWithParameters(jobName, parameters)


	return "{\"status\":\"Success\"}"
}

func (c Jenkins) Hook() string {
	fmt.Print(c.Request)
	return ""
}
