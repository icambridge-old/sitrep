package services

import (
	"github.com/icambridge/genkins"
	"github.com/revel/revel"
)

func GetJenkins() *genkins.Client {
	jenkins := genkins.NewClient(
		revel.Config.StringDefault("jenkins.hostname", ""),
		revel.Config.StringDefault("jenkins.username", ""),
		revel.Config.StringDefault("jenkins.token", ""),
	)

	return jenkins
}
