package mysql1

import (
	"fmt"
	"testing"
)

var test STATEMENT

func TestDelSomeCharacters(t *testing.T) {
	delSomeCharacters("wodetian", 2)
}

func TestSelect(t *testing.T) {
	test.Select([]string{"id", "user_name"})
}

func TestWhere(t *testing.T) {
	test.Where(map[string]string{"id": "6", "email": "666@666.com"})
}

func TestTable(t *testing.T) {
	test.Table("user")
}

func TestUpdate(t *testing.T) {
	test.Update(map[string]string{"id": "nima", "email": "good@gmail.com"})
}

func TestInsert(t *testing.T) {
	test.Insert(map[string]string{"id": "654", "email": "nima@gmail.com"})
}

func TestDesc(t *testing.T) {
	test.gather()
}

func TestOrder(t *testing.T) {
	test.Order([]string{"id"})
}
func TestGather(t *testing.T) {
	test.gather()
	fmt.Println(test.fullStatement)
}

func TestGet(t *testing.T) {
	rows, err := test.Get()
	fmt.Println("表达式：" + test.fullStatement + "\n 错误信息：")
	fmt.Println(err)
	fmt.Println(rows)
}
