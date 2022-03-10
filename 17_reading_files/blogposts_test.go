package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/ichang0301/learn-golang/17_reading_files"
)

func TestBlogPosts(t *testing.T) {
	// Given
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte("Title: Hello, TDD world!")},
		// "hello-go.md":    {Data: []byte("Title: Hello, go!")},
	}

	// When
	posts := blogposts.PostsFromFS(fs)

	// Then
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
	}

	expectedFirstPost := blogposts.Post{Title: "Hello, TDD world!"}
	if !reflect.DeepEqual(posts[0], expectedFirstPost) {
		t.Errorf("got %#v, want %#v", posts[0], expectedFirstPost)
	}
}
