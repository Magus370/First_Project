package main


type Tests []*Test
 
type Test struct {
    PostID int    `json:"postId"`
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Body   string `json:"body"`
}
 
func (Test) TableName() string {
    return "test"
}
 
resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=1")
    if err != nil {
        log.Fatal(err)
    }
 
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    tests:= make(Tests,0)
    err := json.Unmarshal(body, &tests )
    if err != nil {
        fmt.Println(err)
    }
    for test := tests {
         res = gormDb.Create(test)
         if res.Error != nil  {
            panic(res.Error)
          }
    }


    package main
	package main

	import (
		"encoding/json"
		"fmt"
		"io/ioutil"
		"log"
		"net/http"
	
		//"database/sql"
		_ "github.com/go-sql-driver/mysql"
		"gorm.io/driver/mysql"
		"gorm.io/gorm"
	)
	
	type Tests []*Test
	
	type Test struct {
		PostID int    `json:"postId"`
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Body   string `json:"body"`
	}
	
	func (Test) TableName() string {
		return "test"
	}
	func main() {
	
		resp, err := http.Get("https://jsonplaceholder.typicode.com/comments?postId=1")
		if err != nil {
			log.Fatal(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
	
		dsn := "root:dkvjeoc582@tcp(127.0.0.1:3306)/posts?charset=utf8mb4&parseTime=True&loc=Local"
		gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
		}
	
		tests := make(Tests, 0)
		err = json.Unmarshal(body, &tests)
		if err != nil {
			fmt.Println(err)
		}
	
		res := gormDb.AutoMigrate(tests) // на этом этапе в базе данных создаётся наша таблица
		if res.Error != nil {
			panic(res.Error)
		}
	
		for test := range tests {
			gormDb.Create(test)
			if res.Error != nil {
				panic(res.Error)
			}
		}
	}






	

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

// Создал структуру и интерфейс
type Post struct {
	Userid int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (Post) TableName() interface{} {
	return "post"
}

var (
	db *sql.DB
)

// Инициилизировал переменную post. JSON записывает array только в []interface{}
// Но []interface{} не получается записать в БД
var post []Post

func main() {
	// задал body для использования в функциях
	var body []byte
	// функция для запроса на сайт, вернёт декодированные значения
	makeRequest(body)

	// подключение к БД
	db, err := sql.Open("mysql", "root:password@/Posts")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(post)


	//запись в БД, в которой я ОЧЕНЬ не уверен	
		result, err := db.Exec("insert into post values(?)", post)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.LastInsertId())
		
	}

}

func makeRequest(body []byte) interface{} {
	//Запрос на сервер
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts?id=1")
	if err != nil {
		log.Fatal(err)
	}
	//чтение результата
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	fmt.Println("Отправка на декодирование")

	//переход к функции декодирования
	return makeDecoding(body)
}

func makeDecoding(body []byte) interface{} {

	//декодирование
	err := json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
	}
	//Тест успешности декодирования
	fmt.Println(post)
	fmt.Println("Возвращение декодированных значений")
	return post
}
