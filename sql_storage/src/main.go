package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"errors"
)

type Post struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=12345678 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, content, author FROM posts LIMIT $1", limit)
	if err != nil {
		return
	}

	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("SELECT id, content, author FROM posts WHERE id=$1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("SELECT id, content, author FROM comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post) Create() (err error) {
	statement := "INSERT INTO posts(content, author) VALUES($1, $2) RETURNING id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error) {
	_, err = Db.Exec("UPDATE posts SET content=$2, author=$3 WHERE id=$1", post.Id, post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error) {
	_, err = Db.Exec("DELETE FROM posts WHERE id=$1", post.Id)
	return
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("INSERT INTO comments(content, author, post_id) VALUES($1, $2, $3) RETURNING id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Thuan Nguyen"}
	fmt.Println(post)
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Jo Jo", Post: &post}
	comment.Create()

	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)

	// readPost.Content = "Bonjour Monde!"
	// readPost.Author = "Pierre"
	// readPost.Update()
	
	// posts, _ := Posts(10)
	// fmt.Println(posts)

	// readPost.Delete()
}