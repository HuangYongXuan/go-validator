package validator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
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

	types := []string{"uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "int", "uint", "float32", "float64"}
	valType := GetInterfaceType(param)
	b, _ := StringArrayIndex(types, valType)
	return b
}

func InterInterfaceTOInt(val interface{}) int {
	var t2 int
	switch val.(type) {
	case uint:
		t2 = int(val.(uint))
		break
	case int8:
		t2 = int(val.(int8))
		break
	case uint8:
		t2 = int(val.(uint8))
		break
	case int16:
		t2 = int(val.(int16))
		break
	case uint16:
		t2 = int(val.(uint16))
		break
	case int32:
		t2 = int(val.(int32))
		break
	case uint32:
		t2 = int(val.(uint32))
		break
	case int64:
		t2 = int(val.(int64))
		break
	case uint64:
		t2 = int(val.(uint64))
		break
	case float32:
		t2 = int(val.(float32))
		break
	case float64:
		t2 = int(val.(float64))
		break
	case string:
		t2, _ = strconv.Atoi(val.(string))
		break
	default:
		t2 = val.(int)
		break
	}
	return t2
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

func parseStringsToDate(layout string, dates []string) ([]int64, error) {
	if len(dates) != 2 {
		return []int64{}, errors.New("dates is not length to 2")
	}

	first, err := time.Parse(layout, dates[0])
	if err != nil {
		return []int64{}, err
	}
	end, err := time.Parse(layout, dates[1])
	if err != nil {
		return []int64{}, err
	}

	return []int64{first.Unix(), end.Unix()}, nil
}

func ToString(data interface{}) string {
	return fmt.Sprintf("%v", data)
}
