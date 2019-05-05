package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Message struct {
	Id         int
	Name       string
	Email      string
	Content    string
	CreateTime string
}

func main() {
	http.HandleFunc("/", booklist)
	http.HandleFunc("/list", booklist)
	http.HandleFunc("/add", add)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/update", update)
	http.Handle("/static/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/detail", detail)

	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func update(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	name := r.FormValue("name")
	content := r.FormValue("content")

	updatetime := time.Now().Format("2006-01-02 15:04:05")

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	stmt, err := db.Prepare("update message  set name = ?, content = ?, updated_at = ? where id = ?")
	if _, err := stmt.Exec(name, content, updatetime, id); err == nil {
	}

	CheckErr(err)

	booklist(w, r)
	db.Close()

}

func detail(w http.ResponseWriter, r *http.Request) {

	var id int
	var name string
	var email string
	var content string

	vars := r.URL.Query()
	id, err := strconv.Atoi(vars["id"][0])
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	row := db.QueryRow("select id, name, email, content from message where id = ?", id)

	err = row.Scan(&id, &name, &email, &content)
	CheckErr(err)

	msg := new(Message)

	msg.Id = id
	msg.Name = name
	msg.Email = email
	msg.Content = content

	t, err := template.ParseFiles("detail.html")
	CheckErr(err)

	t.Execute(w, &msg)

}

func booklist(w http.ResponseWriter, r *http.Request) {

	var id int
	var name string
	var email string
	var content string
	var created_at string

	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	rows, err := db.Query("select id, name, email, content, created_at from message")
	CheckErr(err)

	var messageSlice []*Message

	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &content, &created_at)
		CheckErr(err)

		msg := new(Message)
		msg.Id = id
		msg.Name = name
		msg.Email = email
		msg.Content = content
		msg.CreateTime = created_at

		messageSlice = append(messageSlice, msg)
	}

	t, err := template.ParseFiles("index.html")
	CheckErr(err)

	t.Execute(w, &messageSlice)
}

func add(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	name := r.FormValue("name")
	content := r.FormValue("content")
	email := r.FormValue("email")
	createtime := time.Now().Format("2006-01-02 15:04:05")
	updatetime := time.Now().Format("2006-01-02 15:04:05")

	stmt, err := db.Prepare("insert into message(name, email, content, created_at, updated_at) values (?, ?, ?, ?, ?)")
	if _, err := stmt.Exec(name, email, content, createtime, updatetime); err == nil {
	}

	booklist(w, r)
	db.Close()
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars["id"][0]
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	stmt, err := db.Prepare("delete from message where id = ?")
	if _, err := stmt.Exec(id); err == nil {
	}

	booklist(w, r)
	db.Close()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("Something Wrong")
		panic(err)
	}
}
