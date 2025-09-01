package userlistdrag

import (
	b64 "encoding/base64"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/senforsce/fh/file"
	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/userconfig"
)

var selectDetails = `
PREFIX rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#>
PREFIX : <http://senforsce.com/o8/brain/>
PREFIX mj: <http://senforsce.com/o8/brain/coach-mj/>

SELECT * WHERE {  ?s rdf:type <http://senforsce.com/o8/brain/coach-mj/User> .
	?s ?p ?o .
	?main rdf:type :UserMainInfo .
    ?main :belongsToUser ?s .
 	OPTIONAL{ ?main :hasBirthdaydate ?Birthdaydate . }
	OPTIONAL{ ?main mj:hasFirstname ?Firstname . }
	OPTIONAL{ ?main :hasLastname ?Lastname . }
	OPTIONAL{ ?main :hasEmail ?Email . }
	OPTIONAL{ ?main :hasTelephone ?Telephone . }
  	OPTIONAL{ ?main :joinedSince ?JoinedSince . }
	OPTIONAL{
		?pic rdf:type :UserProfilePictureFile ;
	    :belongsToUser ?s ;
		:hasSalt ?ProfilePictureSalt  ;
		:hasFilename ?ProfilePictureFilename ;
		:hasFilepath ?ProfilePictureFilepath .
	}
}
`
var panicMessage = "failed sparql query %s"

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}

	uc := userconfig.NewUserConfig()

	if err1 != nil {
		panic("bad")
	}
	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	liste := ListOfSubjects(res.Results.Bindings)
	views := []CompleteMainInfo{}
	for i, _ := range liste {
		view := NewCompleteMainInfo(liste[i])
		key, _ := b64.StdEncoding.DecodeString(view.HasProfilePicture.Salt)
		log.Println(filepath.Join(uc.SubjectStaticRoot, "uploads", view.HasProfilePicture.Filepath))

		decrypted, _ := file.DecryptFile(filepath.Join(uc.SubjectRoot, "uploads", view.HasProfilePicture.Filepath), key)
		view.HasProfilePicture.Data = b64.StdEncoding.EncodeToString(decrypted)
		views = append(views, view)

	}

	return c.Render(HTMX(views, c))
}
