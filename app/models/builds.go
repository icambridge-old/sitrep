package models


type Build struct {
	Id int                 `db:"id" json:"id"`
	BuildId int            `db:"build_id" json:"build_id"`
	Status string          `db:"status" json:"status"`
	ApplicationName string `db:"application_name" json:"job_name"`
	Phase string           `db:"phase" json:"phase"`
	Branch string          `db:"branch" json:"branch"`
	DoneAt string          `db:"done_at" json:"done_at"`
}
