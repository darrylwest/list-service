//
// handlers - methods to handle requests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-11-27 08:35:20
//

package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/go-zoo/bone"
)

var httpClientRequestTimeout = time.Duration(10 * time.Second)

// Handlers the handlers struct for configuration
type Handlers struct {
	cfg *Config
}

// NewHandlers create the new handlers object
func NewHandlers(cfg *Config) *Handlers {
	hnd := Handlers{}
	hnd.cfg = cfg

	return &hnd
}

// HomeHandler read templates; compile and return the home page
func (hnd Handlers) HomeHandler() http.HandlerFunc {
	log.Info("read and serve the index page...")

	type IndexPage struct {
		PageTitle string
		Logo      string
		Version   string
	}

	index := IndexPage{
		PageTitle: "list service",
		Logo:      appLogo(),
		Version:   Version(),
	}

	text := hnd.ReadIndexTemplate()

	t, err := template.New("index").Parse(text)
	if err != nil {
		log.Error("template parse error: %s", err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var data bytes.Buffer
		err = t.Execute(&data, index)
		if err != nil {
			log.Error("error executing template: %s", err)
		}

		log.Info("show the home page: %d bytes", data.Len())
		if n, err := data.WriteTo(w); err != nil {
			log.Error("home page bytes written: %d, error: %s", n, err)
		}
	}
}

// ReadIndexTemplate reads the index file and returns the text string
func (hnd Handlers) ReadIndexTemplate() string {
	box := hnd.cfg.Box

	page, err := box.MustString("/index.html")
	if err != nil {
		log.Error("error reading index : %v", err)
		panic(err)
	}

	return page
}

// QueryHandler - queries and returns list items
func (hnd Handlers) QueryHandler(w http.ResponseWriter, r *http.Request) {
	params := make(map[string]interface{})

	// TODO : validate and execute the query

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["items"] = params

	hnd.writeJSONBlob(w, wrapper)
}

// FindByIDHandler - queries and returns list items
func (hnd Handlers) FindByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := bone.GetValue(r, "id")

	// TODO : query by id

	wrapper := hnd.CreateResponseWrapper("ok")
	wrapper["item"] = id // should be item

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

	// item, err := ListItemFromJSON(data)

	// TODO : fetch and compare to version...

	wrapper := hnd.CreateResponseWrapper("ok")
	// wrapper["item"] = item

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

	// item, err := NewListItemFromJSON(data)

	// item, err = item.Save(hnd.db)

	wrapper := hnd.CreateResponseWrapper("ok")
	// wrapper["item"] = item

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
		log.Warn("attempt to set log level to invalid value: %d, ignored...", level)
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
