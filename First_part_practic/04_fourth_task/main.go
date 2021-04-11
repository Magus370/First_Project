package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	MakeRequest1()

	var input string
	fmt.Scanln(&input)
}

func MakeRequest1() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	var post struct {
		Userid int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}

	err = json.Unmarshal(body, &post)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("1.txt", post, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Open file for reading
	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// ioutil.ReadAll() will read every byte
	// from the reader (in this case a file),
	// and return a slice of unknown slice

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
