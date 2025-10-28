package programme

import (
	"fmt"
	"log"
	"strings"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
)

func Handler(c *router.Context) error {
	c.Render(HTMX(c))
	return nil
}

var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {  
	?s rdf:type <http://senforsce.com/o8/brain/PersonnalizedProgram> .
	?s ?p ?o .
} LIMIT 100`

var programmeDetails = `
PREFIX : <http://senforsce.com/o8/brain/>
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT * WHERE {  
	?s rdf:type <http://senforsce.com/o8/brain/PersonnalizedProgram> .
	?s :persoProgName ?persoProgName .
	?s :hasTrainingMethod ?TrainingMethod .
	?s :persoProgDate ?persoProgDate .
	?s :hasTrainingUser ?TrainingUser .

} LIMIT 1000`

var editDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
SELECT ?s ?p ?o WHERE {
	BIND(<%s> AS ?s)
	?s rdf:type <http://senforsce.com/o8/brain/PersonnalizedProgram> .
	?s ?p ?o .
} LIMIT 100`
var panicMessage = "failed sparql query %s"

func ListHandler(c *router.Context) error {
	repo, ok := c.Get("sparqlRepo").(*sparql.Repo)

	if !ok {
		log.Println("wrong user context")
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

func TableHandler(c *router.Context) error {
	repo, ok := c.Get("sparqlRepo").(*sparql.Repo)

	if !ok {
		log.Println("wrong user context")
	}

	query := programmeDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	log.Println(res.Results.Bindings)

	return c.Render(Table(c, res.Results.Bindings))
}

func EditHandler(c *router.Context) error {
	repo, ok := c.Get("sparqlRepo").(*sparql.Repo)

	if !ok {
		log.Println("wrong user context")
	}

	progname := strings.Replace(c.Param("progname"), "/", "", 1)

	query := fmt.Sprintf(editDetails, progname)
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

	return c.Render(EditProgramme(c, res.Results.Bindings))
}
