//
// list model tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"fmt"
	"lister"
	"testing"

	. "github.com/franela/goblin"
)

func TestListModel(t *testing.T) {
	g := Goblin(t)

	g.Describe("ListModel", func() {
		lister.CreateLogger()

		g.It("should create a list struct", func() {
			model := lister.List{}
			g.Assert(fmt.Sprintf("%T", model)).Equal("lister.List")
		})

        g.It("should serialize a list object to json")
	})
}
