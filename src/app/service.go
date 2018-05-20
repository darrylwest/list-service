//
// service - define the routes and start the service
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package app

import (
	"fmt"
	"net/http"

	// "github.com/go-zoo/bone"
)

// Service - the service struct
type Service struct {
	cfg      *Config
	handlers *Handlers
}

// NewService create a new service by passing in config
func NewService(cfg *Config) (*Service, error) {
	handlers := NewHandlers(cfg)
	svc := Service{cfg, handlers}

	return &svc, nil
}

// StartDatabase create the db, tables, load data
func (svc Service) StartDatabase() error {
	log.Info("start the database...")

	return nil
}

// Start start the admin listener and event loop
func (svc Service) Start() error {
	log.Info("start the hub service...")

	// start the listener
	if err := svc.startServer(); err != nil {
		return err
	}

	return nil
}

/*
// CreateRoutes creates an http router and attaches the handlers
func (svc Service) CreateRoutes() *http.ServerMux {
	log.Info("configure the router/handler routes...")

	hnd := svc.handlers

	mux := 

	router.GetFunc("/api/status", hnd.StatusHandler)
	router.GetFunc("/api/logger", hnd.GetLogLevel)
	router.PutFunc("/api/logger/:level", hnd.SetLogLevel)

	// ok, now the list API...
	router.GetFunc("/list/query", hnd.QueryHandler)
	router.GetFunc("/list/:id", hnd.FindByIDHandler)
	router.PostFunc("/list", hnd.InsertHandler)
	router.PutFunc("/list/:id", hnd.UpdateHandler)
	router.DeleteFunc("/list/:id", hnd.DeleteHandler)

	return router
}
*/

func (svc Service) startServer() error {

    hnd := svc.handlers
	cfg := svc.cfg

    http.Handle("/", http.FileServer(cfg.Box))

	http.HandleFunc("/api/status", hnd.StatusHandler)
	http.HandleFunc("/api/logger", hnd.GetLogLevel)

	host := fmt.Sprintf(":%d", cfg.Port)
	log.Info("start listening on port %s", host)

	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Error("http error: %s", err)
		return err
	}

	return nil
}
