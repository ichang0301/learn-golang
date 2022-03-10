package blogposts

import (
	"io/fs"
)

func PostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")

	var posts []Post
	for _, f := range dir {
		post := makePostFromFile(fileSystem, f.Name())
		posts = append(posts, post)
	}

	return posts
}

func makePostFromFile(fileSystem fs.FS, fileName string) Post {
	blogFile, _ := fileSystem.Open(fileName)
	return newPost(blogFile)
}
