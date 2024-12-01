package texteditor

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

var cachedCss []byte

func Handler(c *router.Context) error {
	var b []byte

	u, err := url.Parse(c.Request.URL.String())
	if err != nil {
		panic(err)
	}

	var root string
	// Replace 'filename' with the actual filename you want to use
	dir, r := os.Getwd()
	if r != nil {
		fmt.Println("Cannot get current directory")
	}

	root = dir + u.Path // careful path comes from frontend

	path := root

	fmt.Printf("path is %s ->/ ,%s %s", u.Path, dir, path)

	full := strings.Replace(path, "/file/open", "", 1)
	if len(path) < 2 {
		return nil
	}
	if len(cachedCss) == 0 {
		//TODO: local hosting of the text editor's css
		res, err := http.Get("https://cdn.jsdelivr.net/npm/ace-builds@1.32.5/css/ace.min.css")

		if err != nil {
			panic("Unable to parse required library")
		}
		b, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()

		cachedCss = b
	} else {
		b = cachedCss
	}

	f, err := os.ReadFile(full)
	if err != nil {
		log.Println(full)
		panic("Unable to parse required file")
	}

	display := strings.Replace(u.Path, "/file/open/", "", 1)

	// string(f), string(b), strings.ReplaceAll(display, "/", " > "), meta.Config{
	// 	TargetSelector: "",
	// 	Inspect:        so,
	// }, c
	so := meta.ShouldInspect(c.Request.URL)

	conf := meta.NewComponentConfig(c)
	cn := conf.WithMetaConfig(meta.Config{
		TargetSelector: "",
		Inspect:        so,
		Payload: map[string]string{
			"Subject": display,
		},
	})
	cte := cn.WithTextEditorConfig(meta.TextEditorConfig{
		RelativeFilePath: strings.ReplaceAll(display, "/", " > "),
		PayloadSelector:  "#sceneTrigger",
		Payload:          display,
		Contents:         string(f),
		LibraryCss:       string(b),
	})

	return c.Render(HTMX(cte.Ctx, cte))
}
