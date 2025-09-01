package service

import (
	"fmt"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
PREFIX : <http://senforsce.com/o8/brain/> 

SELECT * WHERE {  ?s rdf:type :UserBookingRequest .
	?s  :hasBookingStatusString "DemandeEnvoyee";
		:belongsToUser ?BelongsToUser ;
		:hasBookingType ?BookingType ;
			:hasBookingRequestDate ?BookingRequestDate ; 
			:hasBookingComment ?BookingComment .   
} LIMIT 10000`
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	liste := sparql.ListOfSubjects(res.Results.Bindings)
	views := []UserBooking{}
	for _, elem := range liste {

		view := NewUserBooking(elem)
		views = append(views, view)

	}
	return c.Render(HTMX(c, views))
}

type UserBooking struct {
	BelongsToUser      string
	BookingType        string
	BookingRequestDate string
	BookingComment     string
	Subject            string
}

func NewUserBooking(res []map[string]sparql.Binding) UserBooking {
	return UserBooking{
		BelongsToUser:      sparql.GetValue("BelongsToUser", res),
		BookingType:        sparql.GetValue("BookingType", res),
		BookingRequestDate: sparql.GetValue("BookingRequestDate", res),
		BookingComment:     sparql.GetValue("BookingComment", res),
		Subject:            sparql.GetValue("s", res),
	}
}
