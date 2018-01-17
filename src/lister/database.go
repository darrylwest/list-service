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
	workerBucket = []byte("workers")
	queueBucket  = []byte("queue")
	jobBucket    = []byte("jobs")
)

// Database the primary database structure
type Database struct {
	filename string
}

var db = Database{}
var boltdb *bolt.DB

// NewDatabase creates and opens a new database connection
func NewDatabase(cfg *Config) (Database, error) {
	db.filename = cfg.DbFilename

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
		log.Info("create the buckets: %s, %s, %s", workerBucket, queueBucket, jobBucket)

		if _, err = tx.CreateBucketIfNotExists(workerBucket); err != nil {
			log.Error("error creating %s: %s", workerBucket, err)
			return err
		}

		if _, err = tx.CreateBucketIfNotExists(queueBucket); err != nil {
			log.Error("error creating %s: %s", queueBucket, err)
			return err
		}

		if _, err = tx.CreateBucketIfNotExists(jobBucket); err != nil {
			log.Error("error creating %s: %s", jobBucket, err)
			return err
		}

		return nil

	})

	return err
}

// Close close the active database
func (db Database) Close() {
	if boltdb != nil {
		log.Info("close active database...")
		boltdb.Close()
		boltdb = nil
	}
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
