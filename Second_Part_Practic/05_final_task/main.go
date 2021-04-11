package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	Postid int    `json:"postId"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func main() {
	fmt.Println("Запрос постов")
	posts, err := fetchPosts("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Открытие базы данных")
	db, err := sql.Open("mysql", "root:dkvjeoc582@/Posts")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Очистка базы")
	err = cleanupDB(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Сохранение постов")
	err = savePosts(db, posts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Получение и сохранение комментов")
	err = saveComments(db, posts)
	if err != nil {
		fmt.Println("Получил ошибку и в комментах")
		log.Fatal(err)
	}

}

func fetchPosts(url string) ([]Post, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	return posts, err
}

func savePosts(db *sql.DB, posts []Post) error {
	for _, post := range posts {
		_, err := db.Exec("INSERT INTO post (id, userid, title, body) VALUES (?, ?, ?, ?)",
			post.ID, post.Userid, post.Title, post.Body)
		if err != nil {
			return err
		}
	}

	return nil
}

func saveComments(db *sql.DB, posts []Post) error {

	for _, post := range posts {
		comments, err := fetchComments(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v/comments", post.ID))

		if err != nil {
			return err
		}

		for _, comment := range comments {
			_, err = db.Exec("INSERT INTO comments (ID, postid, name,email, body) VALUES (?, ?, ?, ?, ?)",
				comment.ID, comment.Postid, comment.Name, comment.Email, comment.Body)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func fetchComments(url string) ([]Comment, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var Comments []Comment
	err = json.NewDecoder(resp.Body).Decode(&Comments)

	return Comments, err
}

func cleanupDB(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM comments")
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM post")
	if err != nil {
		return err
	}

    return nil
}
