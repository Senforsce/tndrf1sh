package object

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {  
	?s rdf:type <http://senforsce.com/o8/brain/PersonnalizedProgram> .
	?s ?p ?o .
} LIMIT 100`

var editDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
PREFIX rdfs: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>

SELECT * WHERE {
	BIND(<%s> AS ?s)
	BIND(<%s> AS ?p)
	BIND(%s AS ?o)
	?s ?p ?o .
	OPTIONAL { ?p rdf:type ?datatype }
} LIMIT 1000`
var panicMessage = "failed sparql query %s"

func ListHandler(c *router.Context) error {
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

	liste := ListOfSubjects(res.Results.Bindings)
	views := []ViewConfig{}
	for i, _ := range liste {
		view := NewViewConfig(liste[i])
		views = append(views, view)

	}

	return c.Render(List(views, c))
}

func ViewHandler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	objectname := strings.Replace(c.Param("objectname"), "/", "", 1)
	triple := strings.Split(objectname, "§")
	subject := triple[0]
	predicate := triple[1]
	object := `"` + triple[2] + `"`
	if predicate == "http://senforsce.com/o8/brain/persoProgDate" {
		object = `"` + triple[2] + `"` + "^^<http://www.w3.org/2001/XMLSchema#dateTime>"
	}

	if predicate == "http://senforsce.com/o8/brain/hasTrainingMethod" ||
		predicate == "http://senforsce.com/o8/brain/hasTrainingUser" ||
		predicate == "http://senforsce.com/o8/brain/hasEchauffement" ||
		predicate == "http://senforsce.com/o8/brain/hasSportMovement" ||
		predicate == "http://senforsce.com/o8/brain/hasCorpsDeSeance" ||
		predicate == "http://senforsce.com/o8/brain/hasRetourAuCalme" {
		object = `<` + triple[2] + `>`
	}
	query := fmt.Sprintf(editDetails, subject, predicate, object)
	log.Println(query)
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	log.Println(res)
	// liste := ListOfSubjects(res.Results.Bindings)
	// views := []ViewConfig{}
	// for i, _ := range liste {
	// 	view := NewViewConfig(liste[i])
	// 	views = append(views, view)

	// }

	return c.Render(SubjectView(c, res.Results.Bindings))
}

func EditHandler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	objectname := strings.Replace(c.Param("objectname"), "/", "", 1)

	triple := strings.Split(objectname, "§")
	subject := triple[0]
	predicate := triple[1]
	object := `"` + triple[2] + `"`
	if predicate == "http://senforsce.com/o8/brain/persoProgDate" {
		object = `"` + triple[2] + `"` + "^^<http://www.w3.org/2001/XMLSchema#dateTime>"
	}

	if predicate == "http://senforsce.com/o8/brain/hasTrainingMethod" ||
		predicate == "http://senforsce.com/o8/brain/hasTrainingUser" ||
		predicate == "http://senforsce.com/o8/brain/hasEchauffement" ||
		predicate == "http://senforsce.com/o8/brain/hasCorpsDeSeance" ||
		predicate == "http://senforsce.com/o8/brain/hasSportMovement" ||
		predicate == "http://senforsce.com/o8/brain/hasRetourAuCalme" {
		object = `<` + triple[2] + `>`
	}

	query := fmt.Sprintf(editDetails, subject, predicate, object)
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	// liste := ListOfSubjects(res.Results.Bindings)
	// views := []ViewConfig{}
	// for i, _ := range liste {
	// 	view := NewViewConfig(liste[i])
	// 	views = append(views, view)

	// }

	return c.Render(EditView(c, res.Results.Bindings))
}
