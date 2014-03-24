package sitrep

import (
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/bradfitz/gomemcache/memcache"
  "fmt"
)


func BitbucketHook(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  log.Println(r.Form["payload"][0])

  displayPage(w, "about", map[string]interface{}{})
}


func BitbucketListPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo := params["repo"]

	item, err := memClient.Get("bitbucket.pull_reqests")

	if err != nil {
		log.Println(err)
	}

	if item == nil {
		prs, err := bitbucket.PullRequests.GetAll("workstars", repo)

		if err != nil {
			log.Printf("Tried to get pullrequests for %s but got %v", repo, err)
		}
		pullRequests := gobucketToSitRepMultiPrs(prs)

		json, err := json.Marshal(pullRequests)
		if err != nil {
			log.Println(err)
		}
		item = &memcache.Item{Key: "bitbucket.pull_reqests", Value: json, Expiration: 300}
		memClient.Set(item)
	}

	fmt.Fprint(w, string(item.Value))
}
