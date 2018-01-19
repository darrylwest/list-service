//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@ebay.com>
// @created 2018-01-17 17:56:46

package lister

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// Config the config structure
type Config struct {
	Port       int
	LogLevel   int
	DbFilename string
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Port = 8080
	cfg.LogLevel = 2
	cfg.DbFilename = "/data/list-service.db"

	return cfg
}

// ShowHelp dump out the use/command line options
func ShowHelp() {
	fmt.Printf("\n%s USE:\n\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Printf("\n%s Version %s\n", os.Args[0], Version())
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")
	level := flag.Int("loglevel", dflt.LogLevel, "set the server's log level 0..5 for trace..error, default info=2")
	port := flag.Int("port", dflt.Port, "set the server's listening port")
	dbfilename := flag.String("db-filename", dflt.DbFilename, "set the databse file")

	flag.Parse()

	if *vers == true {
		fmt.Printf("Version %s\n", Version())
		return nil
	}

	log.Info("%s Version: %s\n", path.Base(os.Args[0]), Version())

	cfg := Config{
		Port:       *port,
		LogLevel:   *level,
		DbFilename: *dbfilename,
	}

	log.SetLevel(cfg.LogLevel)

	return &cfg
}
