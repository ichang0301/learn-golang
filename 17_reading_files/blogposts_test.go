package blogposts_test

import (
	"testing"
	"testing/fstest"

	blogposts "github.com/ichang0301/learn-golang/17_reading_files"
)

func TestBlogPosts(t *testing.T) {
	// Given
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte("hello, world")},
		"hello-go.md":    {Data: []byte("hello, go")},
	}

	// When
	posts := blogposts.PostsFromFS(fs)

	// Then
	if len(posts) != len(fs) {
		t.Errorf("expected %d posts, got %d posts", len(fs), len(posts))
	}
}
