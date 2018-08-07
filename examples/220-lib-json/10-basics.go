package main

import (
	"encoding/json"
	"fmt"
)

// Person We use tags on struct field declarations to customize the encoded JSON key names
// private fields won't be exported
// tags
type Person struct {
	Name        string   `json:"name"`
	Surname     string   `json:"surname"`
	Age         int      `json:"age"`
	Hobbies     []string `json:"hobbies"`
	NotExported int      `json:"-"`
}

func main() {
	//
	// Encoding
	//
	// encoding basic data types to JSON
	boolByte, _ := json.Marshal(true)
	fmt.Println(string(boolByte))

	intByte, _ := json.Marshal(10)
	fmt.Println(string(intByte))

	floatByte, _ := json.Marshal(3.14)
	fmt.Println(string(floatByte))

	stringByte, _ := json.Marshal("hello")
	fmt.Println(string(stringByte))

	sliceByte, _ := json.Marshal([]string{"foo", "bar", "baz"})
	fmt.Println(string(sliceByte))

	mapByte, _ := json.Marshal(map[string]int{"foo": 5, "bar": 7})
	fmt.Println(string(mapByte))

	// custom data type
	person := &Person{
		Name:        "John",
		Surname:     "Doe",
		Age:         22,
		Hobbies:     []string{"programing", "cycling", "traveling"},
		NotExported: 10,
	}
	personByte, _ := json.Marshal(person)
	fmt.Println(string(personByte))

	//
	// Decoding
	//
	// data structure to store decoded data
	// 	// map[string]interface{}, for JSON objects
	var data map[string]interface{} // interface{} means that values will be various types
	// our json
	jsonByte := []byte(`{"number":6.13,"strings":["a","b"]}`)
	// decoding
	if err := json.Unmarshal(jsonByte, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	// we need to cast values to their appropriate type
	// because we used map interface{}
	n := data["number"].(float64)
	fmt.Println(n)

	// access to embedded data
	// multicasts
	strings := data["strings"].([]interface{})
	s1 := strings[0].(string)
	s2 := strings[1].(string)
	fmt.Println(s1, s2)

	// custom data type
	personByte1 := []byte(`{"name":"John","surname":"Doe","age":22,"hobbies":["programing","cycling","traveling"]}`)
	res := Person{}
	json.Unmarshal(personByte1, &res)
	fmt.Println(res)
	fmt.Println(res.Name)
	fmt.Println(res.Hobbies[0])
}
