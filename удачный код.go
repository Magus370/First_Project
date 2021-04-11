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
	// импровизированная проверка на успешность возврата
	// которую я отключил чтоб не засоряла консоль
	// fmt.Println(post)
	// подключение к БД
	db, err := sql.Open("mysql", "root:dkvjeoc582@/Posts")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(post)
	// Смотрю на то какой тип данных на данный момент
	fmt.Printf("%T\n", post)
	//запись в БД, в которой я ОЧЕНЬ не уверен

	for _, fuck := range post {
		//	result, err := db.Exec("insert into post values(?)", fuck)
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//	fmt.Println(result.LastInsertId())
		fmt.Print(fuck)
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
	// проверка получения данных и какой тип данных
	//fmt.Println(body)
	//fmt.Printf("%T\n", body)
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
