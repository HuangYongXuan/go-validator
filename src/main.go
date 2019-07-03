package main

import (
	"fmt"
	"validator/commit"
)

type Data struct {
	ID       int
	Name     string
	Disabled bool
	Money    float64
}

func main() {
	data := make(map[string]interface{})
	data["ID"] = 1
	data["Name"] = "lin han"
	data["Disabled"] = true
	data["Money"] = 10.65

	rules := make(map[string]interface{})
	rules["ID"] = "required|min:0|integer"
	rules["Name"] = "required|min:0|integer|active_url"
	rules["age"] = "required|integer|min:0|max:120"

	validator := commit.ValidatorMake(data, rules)
	fmt.Println(validator.Fails())
	fmt.Println(validator.GetErrors())
}
