package genkins

type JobView struct {
	Jobs []Job
}

type Job struct {
	Name  string
	Color string
	Url string
}

