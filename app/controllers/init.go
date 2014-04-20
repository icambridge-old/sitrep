package controllers

import (
	"github.com/revel/revel"
	"sitrep/app/hooks"
	"sitrep/app/services"
)

func init() {
	revel.OnAppStart(InitDB)
	revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)

	processor := services.GetHookProcessor()
	processor.Add(&hooks.Unapprove{})
}
