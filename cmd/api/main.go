package main

import (
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	// "github.com/tidwall/gjson"
)

type Todo struct {
	UserID int `json:"userId"`
	ID int	`json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

// func main() {
// 	url := "https://jsonplaceholder.typicode.com/todos/7"

// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	if response.StatusCode == http.StatusOK {

// 		data, err := io.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		todo := gjson.Parse(string(data))
// 		fmt.Println(todo)
// 	}
// }

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/7"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		// todo := Todo{}

		data, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// json.Unmarshal([]byte(data), todo)
		// fmt.Printf("API: %+v\n", string(data))
		fmt.Println(string(data))
	}
}