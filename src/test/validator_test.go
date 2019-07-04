package test

import (
	"fmt"
	"testing"
	"validator/commit"
)

func TestValidateMatch(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "18916444882"

	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "match:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAccepted(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = true
	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "accepted"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateArray(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = []int{}
	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "array"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateConfirmed(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 1212121
	data["name_confirmation"] = 1212121
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "confirmed"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateBetween(t *testing.T) {

	data := make(map[string]interface{})
	data["name"] = "ABCDEFGH"
	data["age"] = 10.25
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "between:10,100"}
	rules["age"] = []string{"required", "min:11", "max:100"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateIn(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "bcc"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "in:asc,bcc"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateNotIn(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 12
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "not_in:asc,bcc"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateNumeric(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "s"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "numeric"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateInteger(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 's'
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "integer"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateString(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = true
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "string"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateEmail(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "754060604@qq.com.cn"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "email"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateIp(t *testing.T) {
	data := make(map[string]interface{})
	data["ip"] = "255.255.255.256"
	rules := make(map[string]interface{})
	rules["ip"] = []string{"required", "ip"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateUrl(t *testing.T) {
	data := make(map[string]interface{})
	data["url"] = "https://www.baidu.com"
	rules := make(map[string]interface{})
	rules["url"] = []string{"required", "url"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAlpha(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "aA"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "Alpha"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAlphaNum(t *testing.T) {

	data := make(map[string]interface{})
	data["data"] = "eqw2314123sad3124dsa"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "alpha_num"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateJson(t *testing.T) {
	msg := `{"name":"zhangsan", "age":18, "id":122463, "sid":122464}`

	data := make(map[string]interface{})
	data["data"] = msg
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "json"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateActiveUrl(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "http://www.asdsadasdwedswadsadwdsadsdas.com/"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "active_url"}

	v := commit.ValidatorMake(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}
