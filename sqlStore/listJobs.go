package sqlStore

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	. "github.com/go-jet/jet/v2/sqlite"
	"log"
)

func ListJobs(db *sql.DB) []*model.Job {
	// Query
	queryString := SELECT(Job.AllColumns).FROM(Job)

	var jobs []*model.Job
	err := queryString.Query(db, &jobs)
	if err != nil {
		log.Fatal(err)
	}

	return jobs
}

func FindJobById(db *sql.DB, id string) (*model.Job, error) {
	// Query
	queryString := SELECT(Job.AllColumns).FROM(Job).WHERE(Job.ID.EQ(String(id)))

	var job model.Job
	err := queryString.Query(db, &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func FindJobByCompanyId(db *sql.DB, companyId string) []*model.Job {
	// Query
	queryString := SELECT(Job.AllColumns).
		FROM(Job).
		WHERE(Job.CompanyId.EQ(String(companyId)))

	var jobs []*model.Job
	err := queryString.Query(db, &jobs)
	if err != nil {
		log.Fatal(err)
	}

	return jobs
}
