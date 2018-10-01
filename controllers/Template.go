package controllers

import (
	"html/template"
	"net/http"
	"time"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/template.html") //, "templates/content.html"
	daysOfWeek := []string{"5", "6", "7", "8"}
	panic(err)
	t.Execute(w, daysOfWeek)
}

func formateData(t time.Time) string {
	layout := "2016-01-01"
	return t.Format(layout)
}

// 自定义模板函数
func Process1(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdata": formateData}
	t := template.New("templates/content.html").Funcs(funcMap)
	t.ParseFiles("templates/content.html")
	t.Execute(w, "")
}

func Process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/template.html", "templates/content1.html")
	t.ExecuteTemplate(w, "template", []string{"1", "2", "3", "4"})
}
