package sqlStore

import (
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	. "github.com/go-jet/jet/v2/sqlite"
)

func (store *Store) FindUserByEmail(email string) (*model.User, error) {
	// Query
	queryString := SELECT(User.AllColumns).FROM(User).WHERE(User.Email.EQ(String(email)))

	var user model.User
	err := queryString.Query(store.db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (store *Store) FindUserById(id string) (*model.User, error) {
	// Query
	queryString := SELECT(User.AllColumns).FROM(User).WHERE(User.ID.EQ(String(id)))

	var user model.User
	err := queryString.Query(store.db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
