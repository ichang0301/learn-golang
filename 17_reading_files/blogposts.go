// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files

// File system abstractions introduced in Go 1.16 : https://pkg.go.dev/io/fs

package blogposts

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := makePostFromFile(fileSystem, f.Name())
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, fileName string) (Post, error) {
	blogFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer blogFile.Close()
	return getPost(blogFile)
}
