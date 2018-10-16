package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func init() {
	if DB == nil {
		loadErr := godotenv.Load()
		connection := os.Getenv("DB_CONNECTION")
		host := os.Getenv("DB_HOST")
		databaseName := os.Getenv("DB_DATABASE")
		username := os.Getenv("DB_USERNAME")
		pwd := os.Getenv("DB_PASSWORD")
		port := os.Getenv("DB_PORT")
		var err error
		DB, err = sql.Open(connection,
			username+":"+pwd+"@tcp("+string(host)+":"+string(port)+")/"+databaseName)
		if err != nil {
			fmt.Println(loadErr)
			log.Fatal(err)
		}
	}
}

func main() {
	rows, err := DB.Query("select id, user_name, email from user")
	type POST struct {
		ID       int
		username string
		email    string
	}
	post := POST{}
	var posts []POST
	// var post map[string]string
	for rows.Next() {
		err = rows.Scan(&post.ID, &post.username, &post.email)
		posts = append(posts, post)
	}
	fmt.Println(posts)
	for key, value := range posts {
		fmt.Println(string(key) + ":" + value.email + "\n")
	}
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
}

func QueryAll() {

}
