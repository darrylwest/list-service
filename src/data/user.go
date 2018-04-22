//
// user
//
// @author darryl.west@ebay.com
// @created 2018-04-22 12:50:14
//

package data

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
type UserDao struct{
    DAO
}

// Query returns a slice of user objects 
func (dao UserDao) Query(db *sql.DB, clause string) ([]User, error) {
	var users []User
	var err error

	stmt := fmt.Sprintf("%s where %s", dao.CreateQuery("users"), clause)

    fmt.Printf("stmt: %s\n", stmt)

	return users, err
}
