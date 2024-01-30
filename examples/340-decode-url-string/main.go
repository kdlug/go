package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	s := "%7B%22name%22%3A%22user1%22%2C%22age%22%3A30%7D" // URL encoded value

	in, err := url.QueryUnescape(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("query escaped %s \n", in)
	var u User
	// unmarshalling
	json.Unmarshal([]byte(in), &u)
	fmt.Printf("username : %s,  age %d \n", u.Name, u.Age)

	//marshalling
	bytes, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bytes))
}
