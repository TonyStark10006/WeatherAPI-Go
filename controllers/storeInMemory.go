package main

import (
	"fmt"
)

type Post struct {
	ID      int
	Content string
	Author  string
}

var PostByID map[int]*Post
var PostByAuthor map[string][]*Post

func main() {
	PostByID = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{ID: 1, Content: "Hello World", Author: "nima"}
	post2 := Post{ID: 2, Content: "Hello World2", Author: "nima2"}
	post3 := Post{ID: 3, Content: "Hello World3", Author: "nima3"}
	post4 := Post{ID: 4, Content: "Hello World4", Author: "nima"}
	store(post1)
	store(post2)
	store(post3)
	store(post4)
	fmt.Println(PostByID[1])
	fmt.Println(PostByAuthor["nima"])

	for _, post := range PostByAuthor["nima"] {
		fmt.Println(post)
	}

	// for _, post := range PostByAuthor["nima1"] {
	// 	fmt.Println(post)
	// }

}

func store(post Post) {
	PostByID[post.ID] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)

}
