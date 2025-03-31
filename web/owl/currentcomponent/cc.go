package currentcomponent

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/senforsce/sparql"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

var term = "<http://senforsce.com/o8/brain/coach-mj/%s>"
var selectDetails = "SELECT ?p ?o WHERE {  %s ?p ?o . } LIMIT 100"
var panicMessage = "failed sparql query %s"

func doIt(rc io.ReadCloser, buflen int64) string {
	defer rc.Close()
	buf := make([]byte, buflen)
	n, err := rc.Read(buf)
	if err != nil && err != io.EOF {

		log.Fatal(err)
	}
	return string(buf[:n])
}

func gippity(s string) (string, error) {
	// URL Decode
	decodedStr, err := url.QueryUnescape(s)
	if err != nil {
		fmt.Println("Error decoding URL:", err)
		return "", err
	}

	// Print the decoded string
	fmt.Println("Decoded String:", decodedStr)

	// Extract the JSON part (after the '=' sign)
	// We know the part after '=' contains the JSON
	jsonPart := decodedStr[strings.Index(decodedStr, "=")+1:]

	// Parse the JSON into a map or struct
	var metaPayload map[string]interface{}
	err = json.Unmarshal([]byte(jsonPart), &metaPayload)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return "", err
	}

	// Print the JSON data
	return metaPayload["Subject"].(string), nil
}

func Handler(c *router.Context) error {
	repo, err1 := sparql.NewRepo("http://localhost:3030/ds",
		sparql.DigestAuth("dba", "dba"),
		sparql.Timeout(time.Millisecond*1500),
	)
	if err1 != nil {
		panic("bad")
	}
	subj, err1 := gippity(doIt(c.Request.Body, c.Request.ContentLength))

	if err1 != nil {
		panic("bad")
	}
	query := fmt.Sprintf(selectDetails, fmt.Sprintf(term, subj))
	res, err := repo.Query(query)

	if err != nil {
		panic(fmt.Sprintf(panicMessage, query))
	}

	so := meta.ShouldInspect(c.Request.URL)
	fmt.Println(res.Results.Bindings)

	view := NewViewConfig(res.Results.Bindings)

	return c.Render(HTMX(res.Results.Bindings, view, meta.Config{
		TargetSelector: "",
		Inspect:        so,
	}, c))
}
