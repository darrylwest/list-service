//
// database tests
//
// @author darryl.west@ebay.com
// @created 2017-12-25 12:57:59
//

package unit

import (
	"fmt"
	"lister"
	"os"
	"testing"
	"time"

	"github.com/boltdb/bolt"

	. "github.com/franela/goblin"
)

const (
	refdb  = "../fixtures/list-service-ref.db"
	testdb = "../data/list-service-test.db"
)

func createTestConfig() *lister.Config {
	cfg := lister.NewDefaultConfig()
	cfg.DbFilename = testdb

	return cfg
}

func copyRefDatabase(copy string) error {
	var db *bolt.DB
	var err error

	db, err = bolt.Open(refdb, 0600, &bolt.Options{ReadOnly: true, Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		err = tx.CopyFile(copy, 600)
		return err
	})

	return err
}

func TestDatabase(t *testing.T) {
	g := Goblin(t)

	if err := copyRefDatabase(testdb); err != nil {
		fmt.Printf("error copying ref database to %s", testdb)
	}

	g.Describe("Databse", func() {
		log := lister.CreateLogger()
		log.SetLevel(4)

		g.It("should create a database struct", func() {
			cfg := createTestConfig()
			db, err := lister.NewDatabase(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", db)).Equal("lister.Database")
		})

		g.It("should open the database", func() {
			cfg := createTestConfig()
			db, err := lister.NewDatabase(cfg)

			err = db.Open()
			defer db.Close()
			g.Assert(err).Equal(nil)
		})

		g.It("should make a backup of a known database", func() {
			cfg := createTestConfig()
			db, err := lister.NewDatabase(cfg)
			g.Assert(err).Equal(nil)
			defer db.Close()

			db.Open()

			bufile := "../data/lister-ref-backup.db"
			err = db.Backup()
			g.Assert(err).Equal(nil)

			// check to see that the database exists
			info, err := os.Stat(bufile)
			g.Assert(err).Equal(nil)
			// fmt.Println(info)
			g.Assert(info.Name()).Equal("lister-ref-backup.db")
			g.Assert(info.Size() > 10000).IsTrue()
		})
	})
}
