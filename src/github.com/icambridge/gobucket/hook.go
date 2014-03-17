package gobucket

import ( 
  "encoding/json"
)

var callbacks = []Callback{}
var request = Request{}

func SetRequest(username string, password string) {
  request = Request{Username: username, Password: password}
}


func AddHook(c Callback) {
  callbacks = append(callbacks, c)
}

func ProccessHook(h *Hook) {
  for _,c := range callbacks {
    c.Exec(h)
  }
    
}

func GetHookData(payload []byte) (*Hook, error) {

  var h Hook
  
  err := json.Unmarshal(payload, &h)
  h.Repository.Request = &request
  return &h, err
}

type Callback interface {
	Exec(h *Hook)
}

type Hook struct {
  Repository Repository `json:"repository"`
  Truncated bool `json:"truncated"`
  Commits []Commit `json:"commits"`
  ConnonUrl string `json:"canon_url"`
  User string `json:"user"`
}


type PlaceInfo struct {
	Commit     CommitInfo          `json:"commit"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type CommitInfo struct {
	Hash       string              `json:"hash"`
	Links      SelfLinks           `json:"links"`
	Repository Repository `json:"repository"`
	Branch     Branch     `json:"branch"`
}

type Branch struct {
	Name string `json:"name"`
}

type Commit struct {
  Node string `json:"node"`
  Files []File `json:"files"`
  RawAuthor string `json:"raw_author"`
  UtcTimestamp string `json:"utctimestamp"`
  Author string `json:"author"`
  Timestamp string `json:"timestamp"`
  RawNode string `json:"raw_node"`
  Parents []string `json:"parents"`
  Branch string `json:"branch"`
  Message string `json:"message"`
  Size int `json:"size"`
}

type File struct {
  Type string `json:"type"`
  File string `json:"file"`
}

