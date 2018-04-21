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
	"service"
	"strings"
	"testing"
	"time"

	"github.com/darrylwest/go-unique/unique"

	. "github.com/franela/goblin"
)

func readListItemResults() (map[string]interface{}, error) {
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

func createListItemModel(title string) service.ListItem {
	item := service.ListItem{}

	item.ID = unique.CreateULID()
	item.DateCreated = time.Now()
	item.LastUpdated = time.Now()
	item.Version = 1
	item.Title = title

	item.Status = service.ListStatusOpen

	return item
}

func TestListModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("ListItemModel", func() {
		service.CreateLogger()

		g.It("should create a list struct", func() {
			model := service.ListItem{}
			g.Assert(fmt.Sprintf("%T", model)).Equal("service.ListItem")
		})

		g.It("should serialize a list object to json", func() {
			model := createListItemModel("my subject")

			g.Assert(fmt.Sprintf("%T", model)).Equal("service.ListItem")

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

		g.It("should unmarshall a partial list item from json and create a new list itme", func() {
			now := time.Now()

			hash := make(map[string]interface{})
			hash["title"] = "My Test Title"

			item, err := service.NewListItemFromJSON(hash)
			g.Assert(err).Equal(nil)

			g.Assert(len(item.ID)).Equal(26)
			g.Assert(item.DateCreated.Year()).Equal(now.Year())
			g.Assert(item.LastUpdated.Year()).Equal(now.Year())
			g.Assert(item.Version).Equal(0)
			g.Assert(item.Title).Equal(hash["title"].(string))
			g.Assert(item.Category).Equal("")
			g.Assert(item.Status).Equal("open")
		})

		g.It("should unmarshall a list of items from json", func() {
			data, err := readListItemResults()
			g.Assert(err).Equal(nil)

			rawItems, ok := data["items"].([]interface{})
			g.Assert(ok).IsTrue()

			for _, raw := range rawItems {
				list, err := service.ListItemFromJSON(raw)

				g.Assert(err).Equal(nil)
				g.Assert(len(list.ID)).Equal(26)
				g.Assert(list.DateCreated.Year()).Equal(2018)
				g.Assert(list.LastUpdated.Year()).Equal(2018)
				g.Assert(list.LastUpdated.Minute()).Equal(list.DateCreated.Minute())
				g.Assert(list.Version).Equal(1)
				g.Assert(len(list.Title) > 5).IsTrue()
				g.Assert(list.Category).Equal("")

				g.Assert(list.Status).Equal(service.ListStatusOpen)

				g.Assert(len(list.Attributes)).Equal(0)
			}

		})
	})
}
