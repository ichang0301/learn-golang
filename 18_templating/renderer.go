// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates

// hotwire: https://hotwired.dev/ (e.g. https://github.com/wolfeidau/hotwire-golang-website)

// 'text/template' documantation: https://pkg.go.dev/text/template

// approval test: https://approvaltests.com/

package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS // Go source files that import "embed" can use the '//go:embed pattern' directive to initialize a variable of type string, []byte, or FS with the contents of files read from the package directory or subdirectories at compile time. 'embed' documantation: https://pkg.go.dev/embed
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r)) // By parsing more than one template into templ we now have to call ExecuteTemplate and specify which template we wish to render as appropriate.
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
