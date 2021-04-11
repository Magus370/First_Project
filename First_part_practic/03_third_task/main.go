package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	
		go MakeRequest1()
		go MakeRequest2()
		go MakeRequest3()
		go MakeRequest4()
		go MakeRequest5()
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
	
	fmt.Println(string(body))
}

func MakeRequest2() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/2")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	fmt.Println(string(body))
}

func MakeRequest3() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/3")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	fmt.Println(string(body))
}

func MakeRequest4() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/4")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	fmt.Println(string(body))
}

func MakeRequest5() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/5")
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	
	fmt.Println(string(body))
}