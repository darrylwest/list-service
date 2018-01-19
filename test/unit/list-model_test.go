//
// list model tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lister"
    "strings"
	"testing"
	"time"

	"github.com/darrylwest/go-unique/unique"

	. "github.com/franela/goblin"
)

func readListResult() (map[string]interface{}, error) {
	var data map[string]interface{}
	filename := "../fixtures/list.json"

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return data, nil
	}

	return data, nil
}

func createListModel(title string) lister.List {
	item := lister.List{}

	item.ID = unique.CreateULID()
	item.DateCreated = time.Now()
	item.LastUpdated = time.Now()
	item.Version = 1
	item.Title = title

	item.Status = lister.ListStatusOpen

	return item
}

func TestListModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("ListModel", func() {
		lister.CreateLogger()

		g.It("should create a list struct", func() {
			model := lister.List{}
			g.Assert(fmt.Sprintf("%T", model)).Equal("lister.List")
		})

		g.It("should serialize a list object to json", func() {
			model := createListModel("my subject")

			g.Assert(fmt.Sprintf("%T", model)).Equal("lister.List")

			blob, err := model.ToJSON()
			g.Assert(err).Equal(nil)
            g.Assert(fmt.Sprintf("%T", blob)).Equal("[]uint8")

            json := fmt.Sprintf("%s", blob)

            g.Assert(strings.Contains(json, "id")).IsTrue()
            g.Assert(strings.Contains(json, "dateCreated")).IsTrue()
            g.Assert(strings.Contains(json, "lastUpdated")).IsTrue()
            g.Assert(strings.Contains(json, "version")).IsTrue()
            g.Assert(strings.Contains(json, "title")).IsTrue()
            g.Assert(strings.Contains(json, "category")).IsTrue()
            g.Assert(strings.Contains(json, "status")).IsTrue()
		})

		g.It("should unmarshall a list of items from json", func() {
			data, err := readListResult()
			g.Assert(err).Equal(nil)

			rawItems, ok := data["items"].([]interface{})
			g.Assert(ok).IsTrue()

			for _, raw := range rawItems {
				list, err := lister.ListFromJSON(raw)

				g.Assert(err).Equal(nil)
				g.Assert(len(list.ID)).Equal(26)
				g.Assert(list.DateCreated.Year()).Equal(2018)
				g.Assert(list.LastUpdated.Year()).Equal(2018)
				g.Assert(list.LastUpdated.Minute()).Equal(list.DateCreated.Minute())
				g.Assert(list.Version).Equal(1)
				g.Assert(len(list.Title) > 5).IsTrue()
				g.Assert(list.Category).Equal("")

				g.Assert(list.Status).Equal(lister.ListStatusOpen)

				g.Assert(len(list.Attributes)).Equal(0)
			}

		})
	})
}
