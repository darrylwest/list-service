//
// access tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"fmt"
	"data"
    "service"
	"testing"

	. "github.com/franela/goblin"
)

func TestData(t *testing.T) {
	g := Goblin(t)

	g.Describe("Access", func() {
		log := service.CreateLogger()

		g.It("should create a DOI struct", func() {
            doi := data.DOI{}
            log.Info("%v", doi)
			g.Assert(fmt.Sprintf("%T", doi)).Equal("data.DOI")
            g.Assert(doi.ID).Equal("")
            g.Assert(doi.Version).Equal(uint64(0))
		})

        g.It("should create a new populated DOI", func() {
            doi := data.NewDOI()
            log.Info("%v", doi)
			g.Assert(fmt.Sprintf("%T", doi)).Equal("data.DOI")
            g.Assert(len(doi.ID)).Equal(26)
            g.Assert(doi.Version).Equal(uint64(0))
        })
	})
}
