package model

import (
	"database/sql"
)

type BuildModel struct {

	Db *sql.DB

}

func (m *BuildModel) Save(b *Build) error {

	_, err := m.Db.Exec("insert into `builds` (`build_id`, `status`, `application_name`, `phase`, `branch`) values (?, ?, ?, ?,?)", b.BuildId, b.Status, b.ApplicationName, b.Phase, b.Branch)

	return err
}

type Build struct {
	Id int
	BuildId int
	Status string
	ApplicationName string
	Phase string
	Branch string
}
