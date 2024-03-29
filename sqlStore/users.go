package sqlStore

import (
	"database/sql"
	"github.com/DaoVuDat/graphql/.gen/model"
	. "github.com/DaoVuDat/graphql/.gen/table"
	. "github.com/go-jet/jet/v2/sqlite"
)

func FindUserByEmail(db *sql.DB, email string) (*model.User, error) {
	// Query
	queryString := SELECT(User.AllColumns).FROM(User).WHERE(User.Email.EQ(String(email)))

	var user model.User
	err := queryString.Query(db, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
