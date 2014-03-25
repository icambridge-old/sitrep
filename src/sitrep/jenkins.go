package sitrep

import (
	"log"
  	"encoding/json"
  	"net/http"
	"fmt"
  	"github.com/icambridge/genkins"
	"github.com/bradfitz/gomemcache/memcache"
)

func getJenkinsJobs() *memcache.Item {
	keyStr := "jenkins.jobs.all"

	item, err := memClient.Get(keyStr)

	if err != nil {
		log.Println(err)
	}

	if item == nil {
		jobs, err := jenkins.Jobs.GetAll()

		if err != nil {
			log.Printf("Tried to get jobs for jenkins but got %v", err)
		}

		json, err := json.Marshal(jobs.Jobs)

		if err != nil {
			log.Printf("Tried to turn jobs into json but got %v", err)
		}

		item = &memcache.Item{Key: keyStr, Value: json, Expiration: 300}
		memClient.Set(item)
	}

	return item
}



func JenkinsHook(w http.ResponseWriter, r *http.Request) {
  // move logic to genkins
  log.Println("==== START OF JENKINS ====")
  p := make([]byte, r.ContentLength)
  _, err := r.Body.Read(p)
  if err != nil {
        log.Println("Unable to unmarshall the JSON request", err);
    return
  }
  var job genkins.Hook
    err1 := json.Unmarshal(p, &job)
    if err1 == nil {
        log.Println(job)
    } else {
        log.Println("Unable to unmarshall the JSON request", err1);
        return

    }


  log.Println("==== END OF JENKINS ====")
}

