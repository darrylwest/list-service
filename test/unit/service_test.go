//
// cache tests
//
// @author darryl.west@ebay.com
// @created 2018-01-16 12:57:59
//

package unit

import (
	"lister"

	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		log := lister.CreateLogger()
		log.SetLevel(3)

		g.It("should create a lister service struct")

	})
}
