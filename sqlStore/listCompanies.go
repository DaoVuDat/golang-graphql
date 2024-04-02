package sqlStore

import (
	"context"
	"fmt"
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	. "github.com/go-jet/jet/v2/sqlite"
	"log"
)

func (store *Store) ListCompany() []*model.Company {
	// Query
	queryString := SELECT(Company.AllColumns).FROM(Company)

	var companies []*model.Company
	err := queryString.Query(store.db, &companies)
	fmt.Println("=== ListCompany")

	if err != nil {
		log.Fatal(err)
	}

	return companies
}

func (store *Store) FindCompanyById(id string) (*model.Company, error) {
	// Query
	queryString := SELECT(Company.AllColumns).
		FROM(Company).
		WHERE(Company.ID.EQ(String(id)))

	var company model.Company
	err := queryString.Query(store.db, &company)

	fmt.Println("=== FindCompanyById")

	if err != nil {
		return nil, err
	}

	return &company, nil
}

func (store *Store) FindCompaniesByIds(ctx context.Context, ids []string) ([]*model.Company, []error) {
	var errs []error
	idsExp := make([]Expression, len(ids))

	for i, v := range ids {
		idsExp[i] = String(v)
	}
	// Query
	queryString := SELECT(Company.AllColumns).
		FROM(Company).
		WHERE(Company.ID.IN(idsExp...)) // Must unpack here

	var companies []*model.Company
	err := queryString.Query(store.db, &companies)

	fmt.Println("=== FindCompaniesByIds")

	if err != nil {
		errs = append(errs, err)
		return nil, errs
	}

	return companies, errs
}

// GetCompanyLoader returns single user by id efficiently
func (store *Store) GetCompanyLoader(ctx context.Context, companyId string) (*model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyLoader.Load(ctx, companyId)
}

// GetCompaniesLoader returns many users by ids efficiently
func (store *Store) GetCompaniesLoader(ctx context.Context, companyIds []string) ([]*model.Company, error) {
	loaders := For(ctx)
	return loaders.CompanyLoader.LoadAll(ctx, companyIds)
}
