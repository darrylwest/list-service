//
// user tests
//
// @author darryl.west <darwest@ebay.com>
// @created 2018-01-17 08:35:20
//

package unit

import (
	"fmt"
	"app"
	"testing"

	. "github.com/franela/goblin"
)

func TestUser(t *testing.T) {
	g := Goblin(t)

	log := app.CreateLogger()
	log.SetLevel(2)

	g.Describe("User", func() {

		g.It("should create a User struct", func() {
			user := app.User{}
			log.Info("%v", user)
			g.Assert(fmt.Sprintf("%T", user)).Equal("app.User")
			g.Assert(user.ID).Equal("")
			g.Assert(user.Version).Equal(uint64(0))
			g.Assert(user.Username).Equal("")
			g.Assert(user.Email).Equal("")
			g.Assert(user.SMS).Equal("")
			g.Assert(user.Info).Equal("")
			g.Assert(user.Status).Equal("")
		})

		g.It("should create a new populated DOI", func() {
			doi := app.NewDOI()
			name := "fredv"

			user := app.NewUser(doi, name)
			log.Info("%v", user)

			g.Assert(fmt.Sprintf("%T", user)).Equal("app.User")
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

	g.Describe("UserDao", func() {
		dao := app.NewUserDao()

		g.It("should create a user dao", func() {
			g.Assert(fmt.Sprintf("%T", dao)).Equal("app.UserDao")
			g.Assert(dao.Table).Equal("users")
		})

		g.It("should create a user table", func() {
			stmt := dao.CreateSchema()
			g.Assert(stmt != "").IsTrue()
			log.Info("%s", stmt)
		})

		g.It("should create a query statement", func() {
			stmt := dao.CreateQuery("status='active'")
			g.Assert(stmt).Equal("select * from users where status='active'")
		})

		g.It("should create a query with sort", func() {
			stmt := dao.CreateQuerySort("status='active'", "LastUpdated")
			g.Assert(stmt).Equal("select * from users where status='active' order by LastUpdated")
		})

		g.It("should create the user schema", func() {
			stmt := dao.CreateSchema()
			g.Assert(stmt != "").IsTrue()
		})
	})
}
