package genkins

type JobView struct {
	Jobs []Job
}

type Job struct {
	Name  string
	Color string
	Url string
}

type HookJob struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Build Build  `json:"build"`
}

type Build struct {
	Number  int    `json:"number"`
	Phase   string `json:"phase"`
	Status  string `json:"status"`
	Url     string `json:"url"`
	FullUrl string `json:"full_url"`
}
