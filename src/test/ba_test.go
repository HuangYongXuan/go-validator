package test

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"testing"
	"validator/commit"
)

func TestStr(t *testing.T) {
	data := make(map[string]interface{})
	data["id"] = 12
	var id string = "id2"
	fmt.Println(data[id])
}

func TestInterfaceType(t *testing.T) {
	fmt.Println(commit.GetInterfaceType(int8(6)))
}

func TestArraySlice(t *testing.T) {
	var array interface{}
	array = make(map[string]string)
	fmt.Println(reflect.TypeOf(array).Kind().String())
}

func BenchmarkValidator(b *testing.B) {
	b.ReportAllocs()

	data := make(map[string]interface{})
	data["ID"] = 1
	data["Name"] = "lin han"
	data["Disabled"] = true
	data["Money"] = 10.65

	rules := make(map[string]interface{})
	rules["ID"] = "required|min:0|integer"
	rules["Name"] = "required|min:0|max:2"
	rules["Age"] = "required|integer|min:0|max:120"

	for i := 0; i < b.N; i++ {
		validator := commit.ValidatorMake(data, rules)
		if validator.Fails() {
		}
	}
}

func TestMegTemplate(t *testing.T) {
	msg := "name 必须介于 :min - :max 之间。"
	params := []string{"1", "2"}

	reg, err := regexp.Compile(":[a-zA-Z]{3,6}")
	if err != nil {
		fmt.Println(err)
		return
	}

	data := reg.FindAllString(msg, 2)

	for key, value := range data {
		msg = strings.Replace(msg, value, params[key], 1)
	}
	fmt.Println(msg)
}
