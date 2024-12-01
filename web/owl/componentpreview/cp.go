package componentpreview

import (
	"context"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	t1 "github.com/senforsce/tndr"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

func Handler(c *router.Context) error {
	// add a comment

	// TODO: get Payload
	// TODO: get Templ.Component from component registry
	// TODO: render this component with loaded component as children
	err := c.Request.ParseForm()
	if err != nil {
		log.Fatalf("error")
	}

	d := c.Request.PostForm.Get("ScenePayload")

	dir := filepath.Dir(d)
	previewPath := strings.Replace(dir, "owl", "preview", 1)
	host := c.Request.Host
	protocol := "http://" //for now

	log.Println(protocol + host + "/" + previewPath)
	resp, err := http.Get(protocol + host + "/" + previewPath)
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)

	conf := meta.NewComponentConfig(c)
	cn := conf.WithMetaConfig(meta.Config{
		TargetSelector: "#metaPayload",

		Inspect: true,
		Payload: map[string]string{
			"Subject": previewPath,
		},
	})

	return c.Render(HTMX(c, preview(sb), cn))
}

func preview(text string) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, text)
		return err
	})
}
