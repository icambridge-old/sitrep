package sitrep

import (
  "log"
  "net/http"
  "encoding/json"
	"github.com/gorilla/mux"
	"github.com/icambridge/gobucket"
  "github.com/bradfitz/gomemcache/memcache"
  "fmt"
	"strings"
	"strconv"
)


func BitbucketHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	payload := []byte(r.Form["payload"][0])

	h, err := gobucket.GetHookData(payload)


	if err != nil {
		log.Println(err)
	}

	hookProcessors.Process(h)

}


func BitbucketListPullRequests(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo := strings.ToLower(params["repo"])

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


func AjaxBitbucketRepo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo := strings.ToLower(params["repo"])

	keyStr := "bitbucket.repo." + repo

	item, err := memClient.Get(keyStr)

	if err != nil {
		log.Println(err)
	}


	bitbucketOwner, _ := cfg.String("bitbucket", "owner")

	if item == nil {
		prs, err := bitbucket.PullRequests.GetAll(bitbucketOwner, repo)
		rawBranches, err := bitbucket.Repositories.GetBranches(bitbucketOwner, repo)

		if err != nil {
			log.Printf("Tried to get pullrequests for %s but got %v", repo, err)
		}
		fmt.Sprintf("%v", prs)
		pullRequests := gobucketToSitRepMultiPrs(prs)
		branches := gobucketToSitRepBranches(repo, rawBranches)

		r := &RepoInfo{Name: repo, PullRequests: pullRequests, Branches: branches}
		json, err := json.Marshal(r)
		if err != nil {
			log.Println(err)
		}
		item = &memcache.Item{Key: keyStr, Value: json, Expiration: 300}
		memClient.Set(item)
	}

	fmt.Fprint(w, string(item.Value))
}

func AjaxBitbucketMerge(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	repo := strings.ToLower(params["repo"])
	id, err   := strconv.Atoi(params["id"])
	if err != nil {
		// handle error
		log.Println(err)
		fmt.Fprint(w, `{"status":"error"}`)
		return
	}
	bitbucketOwner, _ := cfg.String("bitbucket", "owner")
	err = bitbucket.PullRequests.Merge(bitbucketOwner, repo, id, "Merge from SitRep")

	if err != nil {
		// handle error
		log.Println(err)
		fmt.Fprint(w, `{"status":"error"}`)
		return
	}
	fmt.Fprint(w, `{"status":"success"}`)
}
type Unapprove struct {

}

func (u Unapprove) Exec(h *gobucket.Hook) {

	if len(h.Commits) < 1 {
		return
	}

	bitbucketOwner, _ := cfg.String("bitbucket", "owner")

	pr, err := bitbucket.PullRequests.GetBranch(bitbucketOwner, h.Repository.Slug, h.Commits[0].Branch)

	if err != nil {
		// handle error
		log.Println(err)
		return
	}

	if pr != nil {
		return
	}

	bitbucket.PullRequests.Unapprove(bitbucketOwner, h.Repository.Slug, pr.Id)
}
