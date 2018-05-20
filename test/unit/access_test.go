//
// access tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"app"
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestData(t *testing.T) {
	g := Goblin(t)

	g.Describe("Access", func() {
		log := app.CreateLogger()

		g.It("should create a DOI struct", func() {
			doi := app.DOI{}
			log.Info("%v", doi)
			g.Assert(fmt.Sprintf("%T", doi)).Equal("app.DOI")
			g.Assert(doi.ID).Equal("")
			g.Assert(doi.Version).Equal(uint64(0))
		})

		g.It("should create a new populated DOI", func() {
			doi := app.NewDOI()
			log.Info("%v", doi)
			g.Assert(fmt.Sprintf("%T", doi)).Equal("app.DOI")
			g.Assert(len(doi.ID)).Equal(26)
			g.Assert(doi.Version).Equal(uint64(0))
		})
	})
}
