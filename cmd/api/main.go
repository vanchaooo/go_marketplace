package main

import (
	"fmt"
	// "encoding/json"
	"log"
	// "github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	json := `{
		"name": {"first": "Tom", "last": "Anderson"},
		"age": 37,
		"children": ["Sara","Alex","Jack"],
		"fav.movie": "Deer Hunter",
		"friends": [
			{"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
			{"first": "Roger", "last": "Craig", "age": 68, "nets": ["fb", "tw"]},
			{"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		]
	}`

	// 	value, err := sjson.Set(json, "name.first", "Ivan")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	value, err = sjson.Delete(json, "friends.3")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(value)

	value, err := sjson.Delete(json, "friends")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
