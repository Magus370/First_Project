package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	
		MakeRequest()
	
}

func MakeRequest() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	go suck()
	fmt.Println(string(body))
}

func suck()  {
	_, err = db.Exec("insert into posts.post (userId, id, title,body) values (?, ?, ?,?)", 
	body, body, body, body)

  if err != nil {
	  log.Println(err)
  }
}