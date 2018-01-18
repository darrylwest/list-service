//
// list model - a datamodel for list objects
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package lister

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	ListStatusOpen     = "open"
	ListStatusClosed   = "closed"
	ListStatusArchived = "archived"
)

type List struct {
	ID          string                 `json:"id"`
	DateCreated time.Time              `json:"dateCreated"`
	LastUpdated time.Time              `json:"lastUpdated"`
	Version     int                    `json:"version"`
	Title       string                 `json:"title"`
	Category    string                 `json:"category"`
	Attributes  map[string]interface{} `json:"attributes"`
	Status      string                 `json:"status"`
}

// ToJSON
func (list List) ToJSON() ([]byte, error) {
	blob, err := json.Marshal(list)

	return blob, err
}

// ListFromJSON
func ListFromJSON(raw interface{}) (*List, error) {
	var err error
	hash, ok := raw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert raw interface to hash map")
	}

	list := new(List)
	if list.ID, ok = hash["id"].(string); !ok {
		return nil, fmt.Errorf("could not convert to list model: missing id")
	}

	if list.DateCreated, err = ParseDateFromData(hash["dateCreated"]); err != nil {
		return nil, err
	}

	if list.LastUpdated, err = ParseDateFromData(hash["lastUpdated"]); err != nil {
		return nil, err
	}

	if list.Version, err = ParseIntFromData(hash["version"]); err != nil {
		return nil, err
	}

	if list.Title, ok = hash["title"].(string); !ok {
		return nil, fmt.Errorf("could not convert to list model: missing title")
	}

	list.Category = hash["category"].(string)

	if list.Status, ok = hash["status"].(string); !ok {
		return nil, fmt.Errorf("could not convert to list model: missing status")
	}

	return list, nil
}

// ParseDateFromData parse the ISO8601/RFC3339 date by converting to string the parsing
func ParseDateFromData(data interface{}) (time.Time, error) {
	var t time.Time

	str, ok := data.(string)
	if !ok {
		return t, fmt.Errorf("could not get string value")
	}

	t, err := time.Parse(time.RFC3339, str)

	return t, err
}

// ParseIntFromData parse the int64 from raw data by converting the float64 to int64
func ParseIntFromData(data interface{}) (int, error) {
	f, ok := data.(float64)
	if !ok {
		return 0, fmt.Errorf("could not get float64 value")
	}

	return int(f), nil
}
