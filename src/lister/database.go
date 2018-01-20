//
// database - database operations
//
// @author darryl.west@ebay.com
// @created 2018-01-17 12:57:59
//

package lister

import (
    "fmt"
	"strings"
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
	Put(string, []byte) error
	Get(string) ([]byte, error)
	Query(map[string]interface{}) ([]map[string]interface{}, error)
	Remove(string) (map[string]interface{}, error)
	Backup() (string, error)
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

// Put update database with blob 
func (db Database) Put(key string, blob []byte)  error  {
    if len(key) < 26 {
        return fmt.Errorf("invalid key: %s", key)
    }

    log.Info("put %s", blob)

    err := boltdb.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket(listBucket)
        err := b.Put([]byte(key), blob)

        return err
    })

	return err
}

func (db Database) Get(key string) ([]byte, error) {
    var blob []byte

    err := boltdb.View(func(tx *bolt.Tx) error {
        b := tx.Bucket(listBucket)
        blob = b.Get([]byte(key))
        if len(blob) == 0 {
            return fmt.Errorf(DoesNotExistForID, "listitem", key)
        }

        return nil
    })

    log.Info("put %s", blob)
	return blob, err
}

func (db Database) Remove(key string) (map[string]interface{}, error) {
	model := make(map[string]interface{})
	return model, nil
}

func (db Database) Query(params map[string]interface{}) ([]map[string]interface{}, error) {
	list := make([]map[string]interface{}, 0)

	return list, nil
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

// Backup creates a backup file and backs up the current database; the backup filename is returned
func (db Database) Backup() (string, error) {
	bufile := strings.Replace(db.filename, ".db", "-backup.db", 1)
	log.Info("backup current db: %s to %s", db.filename, bufile)
	var err error

	boltdb.View(func(tx *bolt.Tx) error {
		err = tx.CopyFile(bufile, 0644)
		return err
	})

	return bufile, err
}
