//
// user
//
// @author darryl.west@ebay.com
// @created 2018-04-22 12:50:14
//

package service

import (
	"database/sql"
	"fmt"
)

// User the user struct
type User struct {
	DOI
	Username string
	Email    string
	SMS      string
	Info     string
	Status   string
}

// NewUser creates a new user with the given DOI and username
func NewUser(doi DOI, username string) User {
	user := User{}
	user.ID = doi.ID
	user.DateCreated = doi.DateCreated
	user.LastUpdated = doi.LastUpdated
	user.Version = doi.Version

	user.Username = username

	return user
}

// UserDao data access object for user
type UserDao struct {
	DAO
}

// NewUserDao create and return a new user dao
func NewUserDao() UserDao {
	dao := UserDao{}
	dao.Table = "users"
	dao.Select = dao.createSelect()

	return dao
}

func (dao UserDao) createSchemaColumns() string {
    stmt := `Username string not null,
        Fullname string not null,
        SMS string not null,
        Email string not null,
        Info jsonb,
        Status string not null`

    return stmt
}

// CreateSchema creates the table schema
func (dao UserDao) CreateSchema() string {
    stmt := `create table if not exists %s (
        %s,
        %s
    )`

    return fmt.Sprintf(stmt, dao.Table, dao.createDOIColumns(), dao.createSchemaColumns())
}

// Query returns a slice of user objects
func (dao UserDao) Query(db *sql.DB, clause string) ([]User, error) {
	var users []User
	var err error

	stmt := fmt.Sprintf("%s where %s", dao.CreateQuery("users"), clause)

	log.Info("stmt: %s\n", stmt)

	return users, err
}
