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

	"github.com/darrylwest/go-unique/unique"
)

const (
    // ListStatusOpen tags an item as open
	ListStatusOpen     = "open"
    // ListStatusClosed tags an item as closed
	ListStatusClosed   = "closed"
    // ListStatusArchived tags an item as archived
	ListStatusArchived = "archived"
)

// ListItem the list item structure
type ListItem struct {
	ID          string                 `json:"id"`
	DateCreated time.Time              `json:"dateCreated"`
	LastUpdated time.Time              `json:"lastUpdated"`
	Version     int                    `json:"version"`
	Title       string                 `json:"title"`
	Category    string                 `json:"category"`
	Attributes  map[string]interface{} `json:"attributes"`
	Status      string                 `json:"status"`
}

// Save save the current model; update version and last updated date
func (list ListItem) Save(db DataAccessObject) (*ListItem, error) {
    list.Version++
    list.LastUpdated = time.Now()

    blob, err := json.Marshal(list)
    if err != nil {
        return nil, err
    }

    if err = db.Put(list.ID, blob); err != nil {
        return nil, err
    }

    return &list, nil
}

// QueryListItems query and return a list of items
func QueryListItems(db DataAccessObject, params map[string]interface{}) ([]*ListItem, error) {
    var items []*ListItem
    blob, err := db.Query(params)
    if err != nil {
        return items, err
    }

    items = make([]*ListItem, 0, len(blob))
    for _, v := range blob {
        item, err := ParseListItemFromJSON(v)
        if err != nil {
            log.Warn("error parsing item: %s", err)
        }

        items = append(items, item)
    }

    return items, nil
}

// ToJSON convert the struct to a json blob
func (list ListItem) ToJSON() ([]byte, error) {
	blob, err := json.Marshal(list)

	return blob, err
}

// ParseListItemFromJSON parse the blob and return a list item
func ParseListItemFromJSON(blob []byte) (*ListItem, error) {
    var hash map[string]interface{}
    err := json.Unmarshal(blob, &hash)
    if err != nil {
        return nil, err
    }

    return ListItemFromJSON(hash)
}

// NewListItemFromJSON create a new list item from the partial json blob
func NewListItemFromJSON(raw interface{}) (*ListItem, error) {

	hash, ok := raw.(map[string]interface{})
    if !ok {
		return nil, fmt.Errorf("could not convert raw interface to hash map")
    }

	list := new(ListItem)

	if list.Title, ok = hash["title"].(string); !ok {
		return nil, fmt.Errorf("could not convert to list model: missing title")
	}

	list.Category, ok = hash["category"].(string)

	if list.Status, ok = hash["status"].(string); !ok {
		list.Status = ListStatusOpen
	}

    list.ID = unique.CreateULID()
    list.DateCreated = time.Now()
    list.LastUpdated = time.Now()
    list.Version = 0

    return list, nil
}

// ListItemFromJSON convert the hash interface to a list item struct
func ListItemFromJSON(raw interface{}) (*ListItem, error) {
	var err error
	hash, ok := raw.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("could not convert raw interface to hash map")
	}

	list := new(ListItem)
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

	list.Category, ok = hash["category"].(string)

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
