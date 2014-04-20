package services

import (
	"github.com/icambridge/gobucket"
	"github.com/revel/revel"
)

func GetBitbucket() *gobucket.Client {
	bitbucket := gobucket.NewClient(
		revel.Config.StringDefault("bitbucket.username", ""),
		revel.Config.StringDefault("bitbucket.password", ""),
	)

	return bitbucket
}

func GetHookProcessor() *gobucket.HookObserver {
	hookProcessor := &gobucket.HookObserver{}
	return hookProcessor
}
