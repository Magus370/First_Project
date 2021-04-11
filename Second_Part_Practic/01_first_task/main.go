package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Post struct {
	Userid int    `json:"userId"`
	ID     int    `gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Comment struct {
	Postid int    `json:"postId"`
	ID     int    `gorm:"primaryKey"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}

func (Post) TableName() string {
	return "post"
}

func main() {
	fmt.Println("Запрос постов")
	posts, err := fetchPosts("https://jsonplaceholder.typicode.com/posts?userId=7")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Открытие базы данных")
	dsn := "root:dkvjeoc582@tcp(127.0.0.1:3306)/posts?charset=utf8mb4&parseTime=True&loc=Local"
	gormDb, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
	}
	gormDb.Exec("DELETE FROM comments")
	gormDb.Exec("DELETE FROM post")

	fmt.Println("Сохранение постов")
	err = savePosts(gormDb, posts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Получение и сохранение комментов")
	err = saveComments(gormDb, posts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Наконец-то")
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

func savePosts(gormDb *gorm.DB, posts []Post) error {

	for _, post := range posts {
		result := gormDb.Create(post)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func saveComments(gormDb *gorm.DB, posts []Post) error {

	for _, post := range posts {
		comments, err := fetchComments(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v/comments", post.ID))

		if err != nil {
			return err
		}

		for _, comment := range comments {
			result := gormDb.Create(comment)
			if result.Error != nil {
				return result.Error
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
