package componentpreview

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	t1 "github.com/senforsce/tndr"
	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

type CSSContextFile struct {
	IsDir    bool
	Name     string
	Contents string
}

func Handler(c *router.Context) error {
	// add a comment

	// TODO: get Payload
	// TODO: get Templ.Component from component registry
	// TODO: render this component with loaded component as children
	err := c.Request.ParseForm()
	if err != nil {
		log.Fatalf("error")
	}

	cssRoot := c.Get("CSSROOT").(string)

	fmt.Printf("root directory %s\n", cssRoot)

	// Replace 'filename' with the actual filename you want to use
	_, r := os.ReadDir(cssRoot)
	if r != nil {
		fmt.Printf("Cannot get root directory %s\n", cssRoot)
	}
	var csscontents []byte
	e := filepath.Walk(cssRoot, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("dir %s->%s\n", cssRoot, path)

		if err != nil {
			fmt.Printf("Cannot get root wp 1 %s\n", cssRoot)

			fmt.Println(err)
			return nil
		}

		f, err := os.Open(path)

		if err != nil {
			fmt.Printf("Cannot get root wp 2 %s\n", cssRoot)

			return err
		}
		defer f.Close()

		byteValue, _ := io.ReadAll(f)

		csscontents = append(csscontents, byteValue...)
		return nil
	})
	if e != nil {
		fmt.Printf("Cannot get root wp %s\n", cssRoot)

		log.Fatalln(e)
	}
	aggregatedcss := CSSContextFile{false, "preview_dependencies.css", string(csscontents)}
	fmt.Println(aggregatedcss)
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

	return c.Render(HTMX(c, preview(sb), cn, aggregatedcss))
}

func preview(text string) t1.Component {
	return t1.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := io.WriteString(w, text)
		return err
	})
}
