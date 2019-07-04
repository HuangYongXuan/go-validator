package main

import (
	"fmt"
	"time"
	"validator/commit"
)

func main() {
	bT := time.Now()
	data := make(map[string]interface{})
	data["data"] = `{"name":"zhangsan", "age":18, "id":122463, "sid":122464}`
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "json"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
	eT := time.Since(bT)
	fmt.Println("Run time: ", eT)
}
