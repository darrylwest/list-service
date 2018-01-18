//
// list model - a datamodel for list objects
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package lister

import (
    "encoding/json"
	"time"
)

type List struct {
	ID          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
	title       string
	category    string
	attributes  map[string]interface{}
	status      string
}

// ToJSON
func (list List) ToJSON() ([]byte, error) {
    blob, err := json.Marshal(list)
    return blob, err
}

