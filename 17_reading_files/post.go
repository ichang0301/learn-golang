package blogposts

import (
	"bufio"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

func newPost(blogFile io.Reader) Post {
	scanner := bufio.NewScanner(blogFile)

	readline := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readline(titlePrefix)
	description := readline(descriptionPrefix)
	tags := strings.Split(readline(tagsPrefix), ", ")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}
}
