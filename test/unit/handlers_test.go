//
// handlers tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2017-07-27 08:35:20
//

package unit

import (
	"fmt"
	"service"
	"testing"

	. "github.com/franela/goblin"
)

func TestHandlers(t *testing.T) {
	g := Goblin(t)

	g.Describe("Handlers", func() {
		log := service.CreateLogger()
		log.SetLevel(3)
		cfg := service.NewDefaultConfig()

		g.It("should return a valid handler object", func() {
			hnd := service.NewHandlers(cfg)
			g.Assert(fmt.Sprintf("%T", hnd)).Equal("*service.Handlers")

		})

		g.It("should return the home page")
	})
}
