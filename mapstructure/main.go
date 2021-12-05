package main

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func main(){
	type Person struct {
		Name   string
		Age    int
		Emails []string
		Extra  map[string]string
	}

	// This input can come from anywhere, but typically comes from
	// something like decoding JSON where we're not quite sure of the
	// struct initially.
	input := map[string]interface{}{
		"name":   "Mitchell",
		"age":    91,
		"emails": []string{"one", "two", "three"},
		"extra": map[string]string{
			"twitter": "mitchellh",
		},
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		panic(err)
	}

	fmt.Println("%#v", result.Name)
	_ = mapstructure.Decode(result, input)
	fmt.Println(input["age"])
}
