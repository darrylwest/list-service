//
// handlers - methods to handle requests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-11-27 08:35:20
//

package lister

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"
)

var httpClientRequestTimeout = time.Duration(10 * time.Second)

// Handlers the handlers struct for configuration
type Handlers struct {
	cfg *Config
	db  DataAccessObject
}

// NewHandlers create the new handlers object
func NewHandlers(cfg *Config) *Handlers {
	hnd := Handlers{}
	hnd.cfg = cfg

	return &hnd
}

// HomeHandler return the home page
func (hnd Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
    // read and serve the index page...
	hnd.StatusHandler(w, r)
}

// QueryHandler - queries and returns list items
func (hnd Handlers) QueryHandler(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})
	items, err := QueryListItems(hnd.db, params)
	if err != nil {
		http.Error(w, "Query error", 400)
		return
	}

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["items"] = items

	hnd.writeJSONBlob(w, wrapper)
}

// FindByIDHandler - queries and returns list items
func (hnd Handlers) FindByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")

	item, err := GetListItem(hnd.db, id)
	if err != nil {
		http.Error(w, "Missing request body", 400)
		return
	}

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["item"] = item

	hnd.writeJSONBlob(w, wrapper)
}

// UpdateHandler - updates and existing list item
func (hnd Handlers) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Missing request body", 400)
		return
	}

	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Request body has errors", 400)
		return
	}

	item, err := ListItemFromJSON(data)

	// todo -- fetch and compare to version...

	item, err = item.Save(hnd.db)
	if err != nil {
		http.Error(w, "Request body has errors", 400)
		return
	}

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["item"] = item

	hnd.writeJSONBlob(w, wrapper)
}

// InsertHandler - inserts a new list item
func (hnd Handlers) InsertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Missing request body", 400)
		return
	}

	var data map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Error("decode error: %s", err)
		http.Error(w, "Request body has errors", 400)
		return
	}

	item, err := NewListItemFromJSON(data)
	if err != nil {
		log.Error("new item error: %s", err)
		http.Error(w, "Post data parse failed: "+err.Error(), 400)
		return
	}

	item, err = item.Save(hnd.db)
	if err != nil {
		log.Error("save item error: %s", err)
		http.Error(w, "Save data failed: "+err.Error(), 400)
		return
	}

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["item"] = item

	hnd.writeJSONBlob(w, wrapper)
}

// DeleteHandler - archives a list item
func (hnd Handlers) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	hnd.writeErrorResponse(w, "not implemented yet...")
}

// StatusHandler returns the service status
func (hnd *Handlers) StatusHandler(w http.ResponseWriter, r *http.Request) {
	blob := GetStatusAsJSON(hnd.cfg)

	log.Info(blob)
	fmt.Fprintf(w, "%s\n\r", blob)
}

// DBBackupHandler creates a backup of the current database
func (hnd Handlers) DBBackupHandler(w http.ResponseWriter, r *http.Request) {
	db := hnd.db

	if db == nil {
		log.Error("db backup failed, database not assigned")
		hnd.writeErrorResponse(w, "database not assigned")
		return
	}

	format := `{status":"%s","filename":"%s","errors":"%s"}`
	var blob string

	filename, err := db.Backup()
	if err != nil {
		log.Error("db backup failed, target: err: %s", err.Error())
		blob = fmt.Sprintf(format, "failed", filename, err.Error())
	} else {
		blob = fmt.Sprintf(format, "ok", filename, "zero")
	}

	fmt.Fprintf(w, "%s\n\r", blob)
}

// GetLogLevel returns the current log level 0..5
func (hnd Handlers) GetLogLevel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
}

// SetLogLevel sets the log level 1..4
func (hnd *Handlers) SetLogLevel(w http.ResponseWriter, r *http.Request) {
	value := bone.GetValue(r, "level")
	if value == "" {
		hnd.writeErrorResponse(w, "must supply a level between 0 and 5")
		return
	}

	level, err := strconv.Atoi(value)
	if err != nil {
		log.Warn("attempt to set log level to invalid value: %s, ignored...", level)
		hnd.writeErrorResponse(w, err.Error())
		return
	}

	if level < 0 {
		level = 0
	}

	if level > 5 {
		level = 5
	}

	log.SetLevel(level)

	fmt.Fprintf(w, "{\"%s\":\"%d\"}\n\r", "loglevel", log.GetLevel())
}

// CreateResponseWrapper cxreate a map
func (hnd Handlers) CreateResponseWrapper(status string) map[string]interface{} {
	wrapper := make(map[string]interface{})
	wrapper["status"] = status
	wrapper["version"] = "1.0"
	wrapper["ts"] = time.Now()

	return wrapper
}

func (hnd Handlers) writeJSONBlob(w http.ResponseWriter, wrapper map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	blob, err := json.Marshal(wrapper)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), 501)
		return
	}

	log.Debug("blob: %s", blob)
	fmt.Fprintf(w, "%s\n\r", blob)
}

func (hnd Handlers) writeErrorResponse(w http.ResponseWriter, str string) {
	log.Warn(str)
	http.Error(w, str, 501)
}
