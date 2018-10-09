package controllers //main

import (
	"fmt"
	"sync"
	"time"
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

	// 并发编程时如果要保证所有goroutine的任务都执行完毕再执行下一个任务的话可以用到sync包的wait group特性
	var wg sync.WaitGroup
	wg.Add(2)
	go printLetters2(&wg)
	go printNumbers2(&wg)
	wg.Wait()

	// for _, post := range PostByAuthor["nima1"] {
	// 	fmt.Println(post)
	// }

}

func store(post Post) {
	PostByID[post.ID] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)

}

func printNumbers() {
	for i := 0; i <= 10; i++ {
		// fmt.Printf("%d", i)
	}
}

func printLetters() {
	for i := 'A'; i < 'A'+10; i++ {
		// fmt.Printf("%c", i)
	}
}

func printNumbers1() {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Microsecond)
		// fmt.Printf("%d", i)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		// fmt.Printf("%c", i)
	}
}

func print1() {
	printNumbers()
	printLetters()
}

func goPrint1() {
	go printNumbers()
	go printLetters()
}

func print2() {
	printNumbers1()
	printLetters1()
}

func goPrint2() {
	go printNumbers1()
	go printLetters1()
}

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	wg.Done()
}

func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	wg.Done()
}
