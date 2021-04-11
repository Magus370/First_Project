package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (Post) TableName() string {
	return "post"
}

func main() {

	go MakeRequest()

	var input string
	fmt.Scanln(&input)
}

func MakeRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=1")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	//post := make(Post, 0)
	post := new(Post)

	err = json.Unmarshal(body, &post)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post)

	db, err := sql.Open("mysql", "root:dkvjeoc582@/Posts")

	result, err := db.Exec("insert into posts.post ", post)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.LastInsertId())
}
