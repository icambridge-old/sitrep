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

func (m *BuildModel) GetByApplicationNameAndBranch(appName string, branch string) (*Build, error) {

	b := Build{}
	r := m.Db.QueryRow("SELECT id, build_id, status, application_name, phase, branch FROM `builds` WHERE LOWER(`application_name`) = LOWER(?) AND LOWER(branch) = LOWER(?)  ORDER BY ID DESC LIMIT 0,1", appName, branch)

	err := m.populateBuild(r, &b)

	if err != nil {
		return nil, err
	}

	return &b, nil
}

func (m *BuildModel) populateBuild(r *sql.Row, b *Build) error {

	err := r.Scan(&b.Id, &b.BuildId, &b.Status, &b.ApplicationName, &b.Phase, &b.Branch)

	if err != nil {
		return err
	}

	return nil
}

type Build struct {
	Id int
	BuildId int
	Status string
	ApplicationName string
	Phase string
	Branch string
}
