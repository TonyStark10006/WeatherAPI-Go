package mysql1

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
			defer DB.Close()
		}
	}
}

type STATEMENT struct {
	table         string
	selectFields  string
	where         string
	update        string
	insert        string
	fullStatement string
	desc          string
	order         string
}

type ErrorMsg struct {
	msg string
}

func (e ErrorMsg) Error() string {
	return e.msg
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
	s.table = table + " "
	return s
}

// 返回所有结果, sql.Rows指针
func (s *STATEMENT) Get() (*sql.Rows, error) {
	s.gather()
	rows, err := DB.Query(s.fullStatement)
	return rows, err
}

// 返回单行结果
func (s *STATEMENT) First() (*sql.Rows, error) {
	s.gather()
	s.fullStatement += "LIMIT 1"
	fmt.Println(s.fullStatement)
	rows, err := DB.Query(s.fullStatement)
	return rows, err
}

// UPDATE或者INSERT语句调用Exec函数
func (s *STATEMENT) Exec() (sql.Result, error) {
	s.gather()
	fmt.Println(s.fullStatement)
	num, err := DB.Exec(s.fullStatement)
	return num, err
}

// 返回update语句
func (s *STATEMENT) Update(set map[string]string) *STATEMENT {
	for key, val := range set {
		s.update += "`" + key + "` = \"" + val + "\", "
	}
	s.update = delSomeCharacters("SET "+s.update, 2) + " "
	return s
}

// 返回insert语句
func (s *STATEMENT) Insert(insert map[string]string) *STATEMENT {
	var keys string
	var vals string
	for key, val := range insert {
		keys += "`" + key + "`, "
		vals += "\"" + val + "\", "
	}
	s.insert = "(" + delSomeCharacters(keys, 2) + ") VALUES (" + delSomeCharacters(vals, 2) + ") "
	return s
}

// 根据操作类型拼凑最终查询语句
func (s *STATEMENT) gather() {
	if s.insert != "" {
		s.fullStatement = "INSERT INTO " + s.table + s.insert
	}

	if s.update != "" {
		s.fullStatement = "UPDATE " + s.table + s.update + s.where
	}

	if s.selectFields != "" {
		s.fullStatement = s.selectFields + "FROM " + s.table + s.where + s.order + s.desc
	}
}

// 是否倒序
func (s *STATEMENT) Desc() *STATEMENT {
	s.desc = " DESC"
	return s
}

// 拼凑ORDER语句
func (s *STATEMENT) Order(order []string) *STATEMENT {
	for _, val := range order {
		s.order = val + ", "
	}
	s.order = " ORDER BY `" + delSomeCharacters(s.order, 2) + "`"
	return s
}

// 删除拼凑表达式的过程中多余的AND和等号之类的字符
func delSomeCharacters(str string, num int) string {
	if num >= 1 {
		rs := []rune(str)
		length := len(rs) - num
		return string(rs[:length])
	} else {
		//var e ErrorMsg
		//e.msg = "num不可以小于1"
		//return "", e
		panic("num不可以小于1")
	}
}
