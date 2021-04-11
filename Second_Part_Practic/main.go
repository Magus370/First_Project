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
	
//ЧТЕНИЕ ФАЙЛА ОКОНЧЕНО

	db, err:= sql.Open("mysql", "root:dkvjeoc582@/posts")
     
	if err != nil {
		panic(err)
	} 
	defer db.Close()
	 
	
	result, err := db.Exec("insert into posts.post ", body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(result.LastInsertId())  // id добавленного объекта

	fmt.Println(string(body))
}
 