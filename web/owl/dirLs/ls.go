package dirls

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/senforsce/tndr0cean/router"
	"github.com/senforsce/tndrf1sh/web/layout/v1/meta"
)

type File struct {
	isDir bool
	Name  string
}

func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

func Handler(c *router.Context) error {
	// add a comment

	so := meta.ShouldInspect(c.Request.URL)

	u := Must(url.Parse(c.Request.URL.String()))

	files := make(map[string]File, 0)

	var root string
	o8Root := c.Get("O8ROOT").(string)
	// Replace 'filename' with the actual filename you want to use
	_, r := os.ReadDir(o8Root)
	if r != nil {
		fmt.Println("Cannot get root directory")
	}

	root = o8Root + u.Path // careful path comes from frontend

	e := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			fmt.Println(err)
			return nil
		}

		file := File{info.IsDir(), info.Name()}
		files[info.Name()] = file

		return nil
	})

	if e != nil {
		log.Fatal(e)
	}

	f, err2 := os.OpenFile(root, os.O_RDONLY, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}
	fx, err3 := f.Readdirnames(0)
	if err2 != nil {
		log.Fatal(err3)
	}
	res := make([]File, 0)
	for i := range fx {
		g := fx[i]
		res = append(res, files[g])
	}
	return c.Render(HTMX(meta.Config{
		TargetSelector: "",
		Inspect:        so,
	}, c, u.Path, res))
}
