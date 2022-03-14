// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates

// hotwire: https://hotwired.dev/ (e.g. https://github.com/wolfeidau/hotwire-golang-website)

// 'text/template' documantation: https://pkg.go.dev/text/template

// approval test: https://approvaltests.com/

package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS // Go source files that import "embed" can use the '//go:embed pattern' directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time. 'embed' documantation: https://pkg.go.dev/embed
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {

	if err := r.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
