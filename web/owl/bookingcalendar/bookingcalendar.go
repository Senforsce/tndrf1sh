package bookingcalendar

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/senforsce/owl/connected/booking"
	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

//go:embed calendarQuery.sparql
var query string

func Handler(c *router.Context) error {
	repo, ok := c.Get("sparqlRepo").(*sparql.Repo)

	if !ok {
		panic("bad repo context")
	}

	res, err := repo.Query(query)
	if err != nil {
		panic(fmt.Sprintf("bad %s", query))
	}
	log.Println("[sen:TndrCalendar(c, res.Results.Bindings)] : Rendered")

	userlist := booking.ListOfUsers(res.Results.Bindings)

	c.Render(HTMX(c, userlist))
	return nil
}

func TableHandler(c *router.Context) error {
	repo, ok := c.Get("sparqlRepo").(*sparql.Repo)

	if !ok {
		panic("bad repo context")
	}

	res, err := repo.Query(query)
	if err != nil {
		panic(fmt.Sprintf("bad %s", query))
	}
	log.Println("[sen:TndrCalendar(c, res.Results.Bindings)] : Rendered")

	userlist := booking.ListOfUsers(res.Results.Bindings)

	c.Render(Table(c, userlist))
	return nil
}
