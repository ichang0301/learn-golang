// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates

// 'text/template' documantation: https://pkg.go.dev/text/template

package blogrenderer

import (
	"html/template"
	"io"
)

const (
	postTemplate = `<h1>{{.Title}}</h1>
<p>{{.Description}}</p>
Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>` // .Title == post.Title, The templating language is very similar to Mustache. Mastache: https://mustache.github.io/
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	templ, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, post); err != nil {
		return err
	}

	return nil
}
