//
// access
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package data

import (
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

// DAO data access object
type DAO struct {
    Table string
}

// CreateQuery create a query string for the current dao
func (dao DAO) CreateQuery(clause string) string {
    stmt := ""
    return stmt
}
