package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/ichang0301/learn-golang/17_reading_files"
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fileSystem := fstest.MapFS{
			"hello-world.md": {Data: []byte(`Title: Hello, TDD world!
Description: file description
Tags: tdd, go
---
file content
with new line`)},
			// "hello-go.md":    {Data: []byte("Title: Hello, go!")},
		}

		// When
		posts, err := blogposts.PostsFromFS(fileSystem)

		// Then
		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fileSystem) {
			t.Errorf("expected %d posts, got %d posts", len(fileSystem), len(posts))
		}

		assertPost(t, posts[0], blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "file description",
			Tags:        []string{"tdd", "go"},
			Body:        "file content\nwith new line",
		})
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(FailingFS{})
		if err == nil {
			t.Error("expected an error, didn't get one")
		}
	})
}

func assertPost(t *testing.T, got, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %#v, want %#v", got, want)
	}
}

type FailingFS struct {
}

func (f FailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("this function always fail")
}
