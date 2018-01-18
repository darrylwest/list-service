//
// database - database operations
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package lister

import (
	"time"

	"github.com/boltdb/bolt"
)

const (
	DoesNotExistForID = "%s does not exist for id %s"
)

var (
	listBucket = []byte("list")
	boltdb     *bolt.DB
)

type DataAccessObject interface {
	Open() error
	Close() bool
	BackupTo(string) error
	Put(string, map[string]interface{}) (map[string]interface{}, error)
	Get(string) (map[string]interface{}, error)
	Remove(string) (map[string]interface{}, error)
}

// Database the primary database structure
type Database struct {
	filename string
}

// NewDatabase creates and opens a new database connection
func NewDatabase(cfg *Config) (DataAccessObject, error) {
	db := Database{
		filename: cfg.DbFilename,
	}

	return db, nil
}

// Open open the database
func (db Database) Open() error {
	log.Info("open databse %s", db.filename)

	var err error
	boltdb, err = bolt.Open(db.filename, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Error("error opening database: %s %s", db.filename, err)
		return err
	}

	boltdb.Update(func(tx *bolt.Tx) error {
		log.Info("create the list bucket...")

		if _, err = tx.CreateBucketIfNotExists(listBucket); err != nil {
			log.Error("error creating %s: %s", listBucket, err)
			return err
		}

		return nil

	})

	return err
}

func (db Database) Put(key string, model map[string]interface{}) (map[string]interface{}, error) {
	return model, nil
}

func (db Database) Get(key string) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	return model, nil
}

func (db Database) Remove(key string) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	return model, nil
}

// Close close the active database
func (db Database) Close() bool {
	if boltdb != nil {
		log.Info("close active database...")
		boltdb.Close()
		boltdb = nil
		return true
	}

	return false
}

// BackupTo backs up the current database to the specified file
func (db Database) BackupTo(bufile string) error {
	log.Info("backup current db: %s to %s", db.filename, bufile)
	var err error

	boltdb.View(func(tx *bolt.Tx) error {
		err = tx.CopyFile(bufile, 0600)
		return err
	})

	return err
}
