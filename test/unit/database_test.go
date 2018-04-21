//
// database tests
//
// @author darryl.west@ebay.com
// @created 2018-01-19 12:57:59
//

package unit

import (
	"fmt"
	"os"
	"service"
	"testing"
	"time"

	"github.com/boltdb/bolt"
	"github.com/darrylwest/go-unique/unique"

	. "github.com/franela/goblin"
)

const (
	refdb  = "../fixtures/list-service-ref.db"
	testdb = "../data/list-service-test.db"
)

func createTestConfig() *service.Config {
	cfg := service.NewDefaultConfig()
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
		refid := unique.CreateULID()
		refTitle := "My Test Title"

		log := service.CreateLogger()
		log.SetLevel(4)

		g.It("should create a database struct", func() {
			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", db)).Equal("service.Database")
		})

		g.It("should open the database", func() {
			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)

			err = db.Open()
			defer db.Close()
			g.Assert(err).Equal(nil)
		})

		g.It("should insert a new list blob and successfully read it back", func() {
			hash := make(map[string]interface{})
			hash["title"] = refTitle
			hash["category"] = "TopLevel"

			list, err := service.NewListItemFromJSON(hash)
			list.Version++
			g.Assert(err).Equal(nil)
			g.Assert(len(list.ID)).Equal(26)

			// now assign the reference
			list.ID = refid

			blob, err := list.ToJSON()
			g.Assert(err).Equal(nil)

			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)

			db.Open()
			defer db.Close()

			err = db.Put(list.ID, blob)
			g.Assert(err).Equal(nil)

			// now read back the blob
			blob, err = db.Get(refid)
			g.Assert(err).Equal(nil)

			item, err := service.ParseListItemFromJSON(blob)
			g.Assert(err).Equal(nil)
			g.Assert(item.ID).Equal(list.ID)
			g.Assert(item.Title).Equal(refTitle)
		})

		g.It("should update and existing list blob", func() {
			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)
			g.Assert(err).Equal(nil)

			db.Open()
			defer db.Close()

			blob, err := db.Get(refid)
			g.Assert(err).Equal(nil)
			ref, err := service.ParseListItemFromJSON(blob)
			g.Assert(err).Equal(nil)
			g.Assert(ref.ID).Equal(refid)
			g.Assert(ref.Title).Equal(refTitle)

			ref.Title = "My alternate title"

			item, err := ref.Save(db)
			g.Assert(err).Equal(nil)

			g.Assert(item.ID).Equal(ref.ID)
			g.Assert(item.Title).Equal(ref.Title)
			g.Assert(item.LastUpdated.After(item.DateCreated)).IsTrue()
		})

		g.It("should query and return a list of items", func() {
			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)
			g.Assert(err).Equal(nil)

			db.Open()
			defer db.Close()

			params := make(map[string]interface{})
			items, err := db.Query(params)
			g.Assert(err).Equal(nil)

			// fmt.Printf("%s\n", items)
			g.Assert(len(items) > 0).IsTrue()

			for _, v := range items {
				item, err := service.ParseListItemFromJSON(v)
				g.Assert(err).Equal(nil)
				g.Assert(len(item.ID)).Equal(26)
			}
		})

		g.It("should remove an existing list blob")

		g.It("should make a backup of a known database", func() {
			cfg := createTestConfig()
			db, err := service.NewDatabase(cfg)
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
