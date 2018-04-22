//
// user tests
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

func TestUser(t *testing.T) {
	g := Goblin(t)

	g.Describe("User", func() {
		log := service.CreateLogger()

		g.It("should create a User struct", func() {
            user := data.User{}
            log.Info("%v", user)
			g.Assert(fmt.Sprintf("%T", user)).Equal("data.User")
            g.Assert(user.ID).Equal("")
            g.Assert(user.Version).Equal(uint64(0))
            g.Assert(user.Username).Equal("")
            g.Assert(user.Email).Equal("")
            g.Assert(user.SMS).Equal("")
            g.Assert(user.Info).Equal("")
            g.Assert(user.Status).Equal("")
		})

        g.It("should create a new populated DOI", func() {
            doi := data.NewDOI()
            name := "fredv"

            user := data.NewUser(doi, name)
            log.Info("%v", user)

			g.Assert(fmt.Sprintf("%T", user)).Equal("data.User")
            g.Assert(user.ID).Equal(doi.ID)
            g.Assert(user.DateCreated).Equal(doi.DateCreated)
            g.Assert(user.LastUpdated).Equal(doi.LastUpdated)
            g.Assert(user.Version).Equal(uint64(0))
            g.Assert(user.Username).Equal(name)
            g.Assert(user.Email).Equal("")
            g.Assert(user.SMS).Equal("")
            g.Assert(user.Info).Equal("")
            g.Assert(user.Status).Equal("")
        })
	})
}
