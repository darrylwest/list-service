//
// handlers tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"fmt"
	"service"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		service.CreateLogger()

		g.It("should create a config struct", func() {
			cfg := new(service.Config)
			g.Assert(cfg != nil).IsTrue()
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := service.NewDefaultConfig()
			g.Assert(fmt.Sprintf("%T", cfg)).Equal("*service.Config")
			g.Assert(cfg.Port).Equal(80)
			g.Assert(cfg.LogLevel).Equal(2)
			g.Assert(cfg.DbFilename != "").IsTrue()
		})

		g.It("should parse an empty command line and return default config", func() {
			cfg := service.ParseArgs()
			g.Assert(cfg != nil).IsTrue()
		})
	})
}
