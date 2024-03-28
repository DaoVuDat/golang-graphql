package sqlStore

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	. "github.com/go-jet/jet/v2/sqlite"
	"log"
)

func ListCompany(db *sql.DB) []*model.Company {
	// Query
	queryString := SELECT(Company.AllColumns).FROM(Company)

	var companies []*model.Company
	err := queryString.Query(db, &companies)
	if err != nil {
		log.Fatal(err)
	}

	return companies
}

func FindCompanyById(db *sql.DB, id string) (*model.Company, error) {
	// Query
	queryString := SELECT(Company.AllColumns).
		FROM(Company).
		WHERE(Company.ID.EQ(String(id)))

	var company model.Company
	err := queryString.Query(db, &company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
