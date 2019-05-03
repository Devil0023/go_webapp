package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
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
	http.Handle("/static/", http.FileServer(http.Dir("./")))
	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func booklist(w http.ResponseWriter, r *http.Request) {

	var id int
	var name string
	var email string
	var content string
	var create_time int

	db, err := sql.Open("mysql", "messager:messager123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	rows, err := db.Query("select id, name, email, content, create_time from book")
	CheckErr(err)

	var messageSlice []*Message

	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &content, &create_time)
		CheckErr(err)

		msg := new(Message)
		msg.Id = id
		msg.Name = name
		msg.Email = email
		msg.Content = content
		msg.CreateTime = time.Unix(int64(create_time), 0).Format("2006-01-02 15:04:05")

		messageSlice = append(messageSlice, msg)
	}

	t, err := template.ParseFiles("index.html")
	CheckErr(err)

	t.Execute(w, &messageSlice)
}

func add(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "messager:messager123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	name := r.FormValue("name")
	content := r.FormValue("content")
	email := r.FormValue("email")
	createtime := time.Now().Unix()
	updatetime := time.Now().Unix()

	stmt, err := db.Prepare("insert into book(name, email, content, create_time, update_time) values (?, ?, ?, ?, ?)")
	if _, err := stmt.Exec(name, email, content, createtime, updatetime); err == nil {
	}

	booklist(w, r)
	db.Close()
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars["id"][0]
	db, err := sql.Open("mysql", "messager:messager123456@tcp(localhost:3306)/message_book?charset=utf8")
	CheckErr(err)

	stmt, err := db.Prepare("delete from book where id = ?")
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
