package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"sitrep/app/services"
	"sitrep/app/hooks"
	"github.com/icambridge/gobucket"
)

type Bitbucket struct {
	GorpController
}

func (c Bitbucket) Report() revel.Result {

	payload := []byte(c.Params.Get("payload"))


	h, err := gobucket.GetHookData(payload)

	if err != nil {
		fmt.Println(err)
	}

	processor := services.GetHookProcessor()
	// todo move to a centralized location.
	processor.Add(&hooks.Unapprove{})
	processor.Add(&hooks.AutoBuild{})
	processor.Add(&hooks.BranchBuild{})
	processor.Process(h)
	return c.Render()
}
