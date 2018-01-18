//
// handlers tests
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

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		lister.CreateLogger()

		g.It("should create a config struct", func() {
			cfg := new(lister.Config)
			g.Assert(cfg != nil).IsTrue()
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := lister.NewDefaultConfig()
			g.Assert(fmt.Sprintf("%T", cfg)).Equal("*lister.Config")
			g.Assert(cfg.Port).Equal(9040)
			g.Assert(cfg.LogLevel).Equal(2)
			g.Assert(cfg.DbFilename != "").IsTrue()
		})

		g.It("should parse an empty command line and return default config", func() {
			cfg := lister.ParseArgs()
			g.Assert(cfg != nil).IsTrue()
		})
	})
}
