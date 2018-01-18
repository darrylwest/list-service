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
func (hnd *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	hnd.StatusHandler(w, r)
}

// StatusHandler returns the service status
func (hnd *Handlers) StatusHandler(w http.ResponseWriter, r *http.Request) {

	blob := GetStatusAsJSON(hnd.cfg)

	log.Info(blob)
	fmt.Fprintf(w, "%s\n\r", blob)
}

// DBBackupHandler creates a backup of the current database
func (hnd *Handlers) DBBackupHandler(w http.ResponseWriter, r *http.Request) {
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
func (hnd *Handlers) GetLogLevel(w http.ResponseWriter, r *http.Request) {
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

func (hnd *Handlers) writeJSONBlob(w http.ResponseWriter, wrapper map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	blob, err := json.Marshal(wrapper)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), 501)
		return
	}

	log.Info("blob: %s", blob)
	fmt.Fprintf(w, "%s\n\r", blob)
}

func (hnd Handlers) writeErrorResponse(w http.ResponseWriter, str string) {
	log.Warn(str)
	http.Error(w, str, 501)
}
