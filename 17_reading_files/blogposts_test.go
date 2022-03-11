package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/ichang0301/learn-golang/17_reading_files"
)

const (
	theFirstFileBody = `Title: Hello, TDD world!
Description: the first file description
Tags: tdd, hello
---
the 1st file content
with new line`
	theSecondFileBody = `Title: Hello, go!
Description: the second file description
Tags: tdd, go
---
the 2nd file content
with new line`
)

func TestPostsFromFS(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Given
		fileSystem := fstest.MapFS{ // A MapFS is a simple in-memory file system for use in tests, represented as a map from path names (arguments to Open) to information about the files or directories they represent.
			"1_hello-world.md": {Data: []byte(theFirstFileBody)},
			"2_hello-go.md":    {Data: []byte(theSecondFileBody)},
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
			Description: "the first file description",
			Tags:        []string{"tdd", "hello"},
			Body:        "the 1st file content\nwith new line",
		})
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFS(StubFailingFS{})
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

type StubFailingFS struct {
}

func (f StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("this function always fail")
}
