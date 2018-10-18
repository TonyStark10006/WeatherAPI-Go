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

type STATEMENT struct {
	table        string
	selectFields string
	where        string
	full         string
}

func main() {
	good := &STATEMENT{}
	good.Select([]string{"id", "username"}).
		Table("nima").
		//Where(map[string]string{"wo": "detian"}).
		First()
	// rows, err := DB.Query("select id, user_name, email from user")
	// type POST struct {
	// 	ID       int
	// 	username string
	// 	email    string
	// }
	// post := POST{}
	// var posts []POST
	// // var post map[string]string
	// for rows.Next() {
	// 	err = rows.Scan(&post.ID, &post.username, &post.email)
	// 	posts = append(posts, post)
	// }
	// fmt.Println(posts)
	// for key, value := range posts {
	// 	fmt.Println(string(key) + ":" + value.email + "\n")
	// }
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// rows.Close()
}

// 返回select组合语句
func (s *STATEMENT) Select(fields []string) *STATEMENT {
	for _, value := range fields {
		s.selectFields += "`" + value + "`, "
	}
	s.selectFields = delSomeCharacters("SELECT "+s.selectFields, 2) + " "
	return s
}

// 返回where组合语句
func (s *STATEMENT) Where(where map[string]string) *STATEMENT { //where string
	for key, val := range where {
		s.where += "`" + key + "` = \"" + val + "\" AND "
	}
	s.where = "WHERE " + delSomeCharacters(s.where, 4)
	return s
}

// 返回from组合语句
func (s *STATEMENT) Table(table string) *STATEMENT {
	s.table = "FROM " + table + " "
	return s
}

// 返回所有结果
func (s *STATEMENT) Get() {
	s.full = s.selectFields + s.table + s.where
	fmt.Println(s.full)
}

// 返回单行结果
func (s *STATEMENT) First() {
	s.full = s.selectFields + s.table + s.where + "LIMIT 1"
	fmt.Println(s.full)
}

func delSomeCharacters(str string, num int) string {
	if num >= 1 {
		rs := []rune(str)
		length := len(rs) - num
		return string(rs[:length])
	} else {
		panic("num不可以小于1")
	}
}
