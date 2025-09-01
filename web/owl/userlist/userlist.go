package userlist

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
SELECT * WHERE {  ?s rdf:type <http://senforsce.com/o8/brain/coach-mj/User> .
	?s ?p ?o .
	OPTIONAL{
		?pic rdf:type :UserProfilePictureFile ;
	    :belongsToUser ?s ;
		:hasSalt ?ProfilePictureSalt  ;
		:hasFilename ?ProfilePictureFilename ;
		:hasFilepath ?ProfilePictureFilepath .
	}
} LIMIT 1000
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

	query := selectDetails
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	liste := sparql.ListOfSubjects(res.Results.Bindings)
	views := []TndrSimpleUser{}
	uc := userconfig.NewUserConfig()
	for i, _ := range liste {
		view := NewUser(liste[i])
		key, _ := b64.StdEncoding.DecodeString(view.HasProfilePicture.Salt)
		log.Println(filepath.Join(uc.SubjectStaticRoot, "uploads", view.HasProfilePicture.Filepath))

		decrypted, _ := file.DecryptFile(filepath.Join(uc.SubjectRoot, "uploads", view.HasProfilePicture.Filepath), key)
		view.HasProfilePicture.Data = b64.StdEncoding.EncodeToString(decrypted)
		views = append(views, view)

	}

	fmt.Println(res.Results.Bindings)

	return c.Render(HTMX(views, c))
}

type TndrSimpleUser struct {
	City              string
	Firstname         string
	Email             string
	Lastname          string
	HasProfilePicture ProfilePicture
	Subject           string
}

func NewProfilePicture(res []map[string]sparql.Binding, subject string) ProfilePicture {
	return ProfilePicture{
		Filename: sparql.GetValue("ProfilePictureFilename", res),
		Filepath: sparql.GetValue("ProfilePictureFilepath", res),
		Salt:     sparql.GetValue("ProfilePictureSalt", res),
		Subject:  subject,
	}
}

type ProfilePicture struct {
	Filename string
	Data     string
	Filepath string
	Salt     string
	Subject  string
}

func NewUser(results []map[string]sparql.Binding) TndrSimpleUser {
	return TndrSimpleUser{
		City:              sparql.FindObjectValueByPredicate("hasCity", results)["o"].Value,
		Firstname:         sparql.FindObjectValueByPredicate("hasFirstname", results)["o"].Value,
		Email:             sparql.FindObjectValueByPredicate("hasEmail", results)["o"].Value,
		Lastname:          sparql.FindObjectValueByPredicate("hasLastname", results)["o"].Value,
		HasProfilePicture: NewProfilePicture(results, sparql.FindObjectValueByPredicate("hasFirstname", results)["s"].Value),

		Subject: sparql.FindObjectValueByPredicate("hasEmail", results)["s"].Value,
	}
}
