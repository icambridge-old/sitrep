package models

type Setting struct {
	Id int64 `db:"id"`
	Name string `db:"name"`
	Value string `db:"value"`
}


