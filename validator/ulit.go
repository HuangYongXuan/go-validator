package validator

import (
	"reflect"
	"strings"
)

// snake string, XxYy to xx_yy , XxYY to xx_yy
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}

// camel string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func GetInterfaceType(param interface{}) string {
	if param == nil {
		return ""
	}
	return reflect.TypeOf(param).String()
}

func StringArrayIndex(array []string, value string) (bool, int) {
	length := len(array)
	b := false
	index := -1
	for i := 0; i < length; i++ {
		if array[i] == value {
			b = true
			index = i
			break
		}
	}
	return b, index
}

func InterfaceIsInteger(param interface{}) bool {
	if param == nil {
		return false
	}

	types := []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "int", "uint"}
	valType := GetInterfaceType(param)
	b, _ := StringArrayIndex(types, valType)
	return b
}

func InterfaceIsNumeric(param interface{}) bool {
	if param == nil {
		return false
	}

	types := []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "int", "uint", "float32", "float64", "complex64", "complex128"}
	valType := GetInterfaceType(param)

	b, _ := StringArrayIndex(types, valType)
	return b
}

func IsArray(param interface{}) bool {
	if param == nil {
		return false
	}
	ty := reflect.ValueOf(param).Kind().String()
	return ty == "slice" || ty == "map"
}
