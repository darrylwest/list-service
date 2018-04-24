//
// access
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package service

import (
	"database/sql"
	"fmt"
	"github.com/darrylwest/go-unique/unique"
	"time"
)

// DOI digital object identifier shared by all models
type DOI struct {
	ID          string
	DateCreated time.Time
	LastUpdated time.Time
	Version     uint64
}

// NewDOI create a new DOI with ID set to ULID, date created + last updated set to now (utc), and version set to zero
func NewDOI() DOI {
	now := time.Now().UTC()
	doi := DOI{
		ID:          unique.CreateULID(),
		DateCreated: now,
		LastUpdated: now,
	}

	return doi
}

func CreateDatabase(db *sql.DB) error {
	// stmt := "create dataase if not exists lists";

	return nil
}

// DAO data access object
type DAO struct {
	Table  string
	Select string
}

// create the basic schema template for DOI and custom columns
func (dao DAO) createSchemaStatement() string {
	return `create table if not exists %s (
        %s,
        %s
    )`
}

func (dao DAO) createDOIColumns() string {
	stmt := "ID string primary key,\n\tDateCreated timestamp not null,\n\tLastUpdated timestamp not null,\n\tVersion int not null"

	return stmt
}

func (dao DAO) createSelect() string {
	return fmt.Sprintf("select * from %s", dao.Table)
}

// CreateQuery create a query string for the current dao
func (dao DAO) CreateQuery(clause string) string {
	return fmt.Sprintf("%s where %s", dao.Select, clause)
}

// CreateQuerySort create a query string with clause and sort/order by
func (dao DAO) CreateQuerySort(clause, sortby string) string {
	return fmt.Sprintf("%s order by %s", dao.CreateQuery(clause), sortby)
}
