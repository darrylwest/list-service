//
// database tests
//
// @author darryl.west@ebay.com
// @created 2018-01-19 12:57:59
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

	db, err = bolt.Open(refdb, 0644, &bolt.Options{ReadOnly: true, Timeout: 1 * time.Second})

	if err != nil {
		return err
	}

	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		err = tx.CopyFile(copy, 0644)
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

        g.It("should insert a new list blob and successfully read it back", func() {
            hash := make(map[string]interface{})
            hash["title"] = "My Test Title"
            hash["category"] = "TopLevel"
            
            list, err := lister.NewListItemFromJSON(hash)
            list.Version++
            g.Assert(err).Equal(nil)
            g.Assert(len(list.ID)).Equal(26)

            blob, err := list.ToJSON()
            g.Assert(err).Equal(nil)

			cfg := createTestConfig()
			db, err := lister.NewDatabase(cfg)

			err = db.Open()
			defer db.Close()

            err = db.Put(list.ID, blob)
            g.Assert(err).Equal(nil)

            // now read back the blob
            blob, err = db.Get(list.ID)
            g.Assert(err).Equal(nil)

            item, err := lister.ParseListItemFromJSON(blob)
            g.Assert(err).Equal(nil)
            g.Assert(item.ID).Equal(list.ID)
        })

        g.It("should update and existing list blob")

        g.It("should remove an existing list blob")

        g.It("should query and return a list of items")

		g.It("should make a backup of a known database", func() {
			cfg := createTestConfig()
			db, err := lister.NewDatabase(cfg)
			g.Assert(err).Equal(nil)
			defer db.Close()

			db.Open()

			bufile, err := db.Backup()
			g.Assert(err).Equal(nil)

			// check to see that the database exists
			info, err := os.Stat(bufile)
			g.Assert(err).Equal(nil)
			log.Info("%v", info)

			g.Assert(info.Size() > 10000).IsTrue()
		})
	})
}
