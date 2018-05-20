//
// service
//
// @author darryl.west@ebay.com
// @created 2017-11-27 12:57:59
//

package unit

import (
	"app"
	"fmt"
	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		log := app.CreateLogger()
		log.SetLevel(4)
		cfg := app.NewDefaultConfig()

		g.It("should create a service struct", func() {
			service, err := app.NewService(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", service)).Equal("*app.Service")
		})

		g.It("should initialize the router", func() {
			service, err := app.NewService(cfg)
			g.Assert(err).Equal(nil)

			router := service.CreateRoutes()

			log.Debug("router: %v", router)

		})

	})
}
