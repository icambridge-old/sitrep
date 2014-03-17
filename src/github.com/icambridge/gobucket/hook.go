package gobucket

import ( 
  "encoding/json"
)

func GetHookData(payload []byte) (*Hook, error) {

  var h Hook
  
  err := json.Unmarshal(payload, &h)
  
  return &h, err
}

type Hook struct {
  Repository Repository `json:"repository"`
  Truncated bool `json:"truncated"`
  Commits []Commit `json:"commits"`
  ConnonUrl string `json:"canon_url"`
  User string `json:"user"`
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

type Repository struct {
	Website     string `json:"website"`
	Fork        bool   `json:"fork"`
	Name        string `json:"name"`
	Scm         string `json:"scm"`
	Owner       string `json:"owner"`
	AbsoluteUrl string `json:"absolute_url"`
	Slug        string `json:"slug"`	 
	Private     bool   `json:"is_private"`
}
