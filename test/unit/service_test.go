//
// service
//
// @author darryl.west@ebay.com
// @created 2017-11-27 12:57:59
//

package unit

import (
	"fmt"
	"service"
	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		log := service.CreateLogger()
		log.SetLevel(4)
		cfg := service.NewDefaultConfig()

		g.It("should create a service struct", func() {
			service, err := service.NewService(cfg)
			g.Assert(err).Equal(nil)
			g.Assert(fmt.Sprintf("%T", service)).Equal("*service.Service")
		})

		g.It("should initialize the router", func() {
			service, err := service.NewService(cfg)
			g.Assert(err).Equal(nil)

			router := service.CreateRoutes()

			log.Debug("router: %v", router)

		})

	})
}
