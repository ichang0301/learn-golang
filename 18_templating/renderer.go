package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	if _, err := fmt.Fprintf(w, "<h1>%s</h1>\n", post.Title); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "<p>%s</p>\n", post.Description); err != nil {
		return err
	}
	if _, err := fmt.Fprint(w, "Tags: <ul>"); err != nil {
		return err
	}
	for _, tag := range post.Tags {
		if _, err := fmt.Fprintf(w, "<li>%s</li>", tag); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprint(w, "</ul>"); err != nil {
		return err
	}
	return nil
}
