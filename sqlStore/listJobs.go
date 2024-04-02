package sqlStore

import (
	"fmt"
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	"github.com/dgryski/trifles/uuid"
	. "github.com/go-jet/jet/v2/sqlite"
	"log"
	"time"
)

func (store *Store) ListJobs() []*model.Job {
	// Query
	queryString := SELECT(Job.AllColumns).FROM(Job)

	var jobs []*model.Job
	err := queryString.Query(store.db, &jobs)
	fmt.Println("=== ListJobs")

	if err != nil {
		log.Fatal(err)
	}

	return jobs
}

func (store *Store) FindJobById(id string) (*model.Job, error) {
	// Query
	queryString := SELECT(Job.AllColumns).FROM(Job).WHERE(Job.ID.EQ(String(id)))

	var job model.Job
	err := queryString.Query(store.db, &job)
	fmt.Println("=== FindJobById")

	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (store *Store) FindJobByCompanyId(companyId string) []*model.Job {
	// Query
	queryString := SELECT(Job.AllColumns).
		FROM(Job).
		WHERE(Job.CompanyId.EQ(String(companyId)))

	var jobs []*model.Job
	err := queryString.Query(store.db, &jobs)
	fmt.Println("=== FindJobByCompanyId")

	if err != nil {
		log.Fatal(err)
	}

	return jobs
}

func (store *Store) CreateJob(companyId, title string, description *string) (*model.Job, error) {
	job := model.Job{
		ID:          uuid.UUIDv4(),
		CompanyId:   companyId,
		Title:       title,
		Description: description,
		CreatedAt:   time.Now().String(),
	}

	// Query
	queryString := Job.INSERT(Job.AllColumns).MODEL(job).RETURNING(Job.AllColumns)

	var insertedJob model.Job
	err := queryString.Query(store.db, &insertedJob)
	if err != nil {
		return nil, err
	}

	return &insertedJob, nil

}

func (store *Store) DeleteJob(id string) (*model.Job, error) {
	// Query
	queryString := Job.DELETE().WHERE(Job.ID.EQ(String(id))).RETURNING(Job.AllColumns)

	var deletedJob model.Job
	err := queryString.Query(store.db, &deletedJob)
	if err != nil {
		return nil, err
	}

	return &deletedJob, nil
}

func (store *Store) UpdateJob(id string, title, description *string) (*model.Job, error) {
	var fieldToUpdate ColumnList
	var modelToUpdate model.Job
	if title != nil {
		fieldToUpdate = append(fieldToUpdate, Job.Title)
		modelToUpdate.Title = *title
	}

	if description != nil {
		fieldToUpdate = append(fieldToUpdate, Job.Description)
		modelToUpdate.Description = description
	}
	// Query
	queryString := Job.UPDATE(fieldToUpdate).
		WHERE(Job.ID.EQ(String(id))).
		MODEL(modelToUpdate).RETURNING(Job.AllColumns)

	var updatedJob model.Job
	err := queryString.Query(store.db, &updatedJob)
	if err != nil {
		return nil, err
	}

	return &updatedJob, nil

}
