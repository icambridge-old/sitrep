package controller

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/bradfitz/gomemcache/memcache"
	"html/template"
)

type Home struct {
	Base
}

func (c Home) Index() string {
	glog.Info("Index page hit")

	jobsJson := string(c.getJenkinsJobs().Value)

	return c.Controller.RenderTemplate("index", map[string]interface{}{"jobs": template.JS(jobsJson)})
}

func (c Home) getJenkinsJobs() *memcache.Item {
	glog.Info("Fetching jenkins info")
	memClient  := c.Memcache
	keyStr := "jenkins.jobs.all"

	item, err := memClient.Get(keyStr)

	if err != nil {
		glog.Info(err)
	}
	jenkins := c.Jenkins
	if item == nil {
		jobs, err := jenkins.Jobs.GetAll()

		if err != nil {
			glog.Error("Tried to get jobs for jenkins but got %v", err)
		}

		json, err := json.Marshal(jobs.Jobs)

		glog.Infof("Found %s",  string(json))

		if err != nil {
			glog.Error("Tried to turn jobs into json but got %v", err)
		}

		item = &memcache.Item{Key: keyStr, Value: json, Expiration: 300}
		memClient.Set(item)
	}

	return item
}
