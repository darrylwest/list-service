//
// handlers tests
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

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		app.CreateLogger()

		g.It("should create a config struct", func() {
			cfg := new(app.Config)
			g.Assert(cfg != nil).IsTrue()
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := app.NewDefaultConfig()
			g.Assert(fmt.Sprintf("%T", cfg)).Equal("*app.Config")
			g.Assert(cfg.Port).Equal(80)
			g.Assert(cfg.LogLevel).Equal(2)
		})

		g.It("should parse an empty command line and return default config", func() {
			cfg := app.ParseArgs()
			g.Assert(cfg != nil).IsTrue()
		})
	})
}
