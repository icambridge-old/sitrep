package controller

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/icambridge/framework"
	"html/template"
)

type Home struct {
	framework.Controller
}

func (c Home) Index() string {
	glog.Info("Index page hit")

	jobsJson := string(c.getJenkinsJobs().Value)

	return c.Controller.RenderTemplate(w, "index", map[string]interface{}{"jobs": template.JS(jobsJson)})
}

func (c Home) getJenkinsJobs() *memcache.Item {

	var memClient memcache.Client = c.Container.Get("memcache")
	keyStr := "jenkins.jobs.all"

	item, err := memClient.Get(keyStr)

	if err != nil {
		glog.Info(err)
	}
	jenkins := c.Container.Get("jenkins")
	if item == nil {
		jobs, err := jenkins.Jobs.GetAll()

		if err != nil {
			glog.Error("Tried to get jobs for jenkins but got %v", err)
		}

		json, err := json.Marshal(jobs.Jobs)

		if err != nil {
			glog.Error("Tried to turn jobs into json but got %v", err)
		}

		item = &memcache.Item{Key: keyStr, Value: json, Expiration: 300}
		memClient.Set(item)
	}

	return item
}
