package test

import (
	"fmt"
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
	var a interface{}
	a = []string{"a"}
	fmt.Println(commit.InterfaceType(a))
}
