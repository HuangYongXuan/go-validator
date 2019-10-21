package test

import (
	"fmt"
	"testing"

	"github.com/Ysll233/go-validator/validator"
)

func TestCustomMessages(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "1890000000"

	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "match:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}

	customMessages := make(map[string]interface{})
	customMessages["match"] = "格式不正确"
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestCustomNames(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "1890000000"

	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "match:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}

	customMessages := make(map[string]interface{})
	customMessages["match"] = ":attr 格式不正确"
	customNames := make(map[string]string)
	customNames["name"] = "名称"
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateSometimes(t *testing.T) {
	// TODO
}

func TestValidateBail(t *testing.T) {
	// TODO
}

func TestValidateRequired(t *testing.T) {
	// TODO
}

func TestValidatePresent(t *testing.T) {
	// TODO
}

func TestValidateFilled(t *testing.T) {
	// TODO
}

func TestValidateRequiredWith(t *testing.T) {
	// TODO
}

func TestValidateRequiredWithAll(t *testing.T) {
	// TODO
}

func TestValidateRequiredWithout(t *testing.T) {
	// TODO
}

func TestValidateRequiredWithoutAll(t *testing.T) {
	// TODO
}

func TestValidateRequiredIf(t *testing.T) {
	// TODO
}

func TestValidateRequiredUnless(t *testing.T) {
	// TODO
}

func TestValidateMatch(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "18900000001"

	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "match:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	if v.Fails() {
		err := v.GetErrors()
		fmt.Println(err)
	}
}

func TestValidateRegex(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "18900000002"

	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "regex:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAccepted(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = true
	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "accepted"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateArray(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = []int{}
	rules := make(map[string]interface{})
	rules["name"] = []string{"nullable", "array"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateConfirmed(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 1212121
	data["name_confirmation"] = 1212121
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "confirmed"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateSame(t *testing.T) {
	// TODO
}

func TestValidateDifferent(t *testing.T) {
	// TODO
}

func TestValidateDigits(t *testing.T) {
	// TODO
}

func TestValidateDigitsBetween(t *testing.T) {
	// TODO
}

func TestValidateSize(t *testing.T) {
	// TODO
}
func TestValidateBetween(t *testing.T) {

	data := make(map[string]interface{})
	data["name"] = "ABCDEFGH"
	data["age"] = 10.25
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "between:10,100"}
	rules["age"] = []string{"required", "min:11", "max:100"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateMin(t *testing.T) {
	// TODO
}

func TestValidateMax(t *testing.T) {
	// TODO
}

func TestValidateIn(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "bcc"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "in:asc,bcc"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateNotIn(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 12
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "not_in:asc,bcc"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateNumeric(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "s"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "numeric"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateInteger(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = 1
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "integer"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateString(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = true
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "string"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateEmail(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "754060604@qq.com.cn"
	rules := make(map[string]interface{})
	rules["name"] = []string{"required", "email"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateIp(t *testing.T) {
	data := make(map[string]interface{})
	data["ip"] = "255.255.255.256"
	rules := make(map[string]interface{})
	rules["ip"] = []string{"required", "ip"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateUrl(t *testing.T) {
	data := make(map[string]interface{})
	data["url"] = "https://www.baidu.com"
	rules := make(map[string]interface{})
	rules["url"] = []string{"required", "url"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAlpha(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "aA"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "Alpha"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAlphaNum(t *testing.T) {

	data := make(map[string]interface{})
	data["data"] = "eqw2314123sad3124dsa"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "alpha_num"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAlphaDash(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "123213123~!@#$%^&*("
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "alpha_dash"}
	v := validator.Make(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateBefore(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "2019-12-25 12:00:00"
	data["data2"] = "1562396373"
	data["data3"] = 1562396373
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "before:2019-12-25 12:00:00"}
	rules["data2"] = []string{"required", "before:2019-04-25 12:00:00"}
	rules["data3"] = []string{"required", "before:2019-04-25 12:00:00"}
	v := validator.Make(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateBeforeOrEqual(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "2019-12-25 12:00:00"
	data["data2"] = "1562396373"
	data["data3"] = 1562396373
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "before_or_equal:2019-12-25 12:00:00"}
	rules["data2"] = []string{"required", "before_or_equal:2019-04-25 12:00:00"}
	rules["data3"] = []string{"required", "before_or_equal:2019-04-25 12:00:00"}
	v := validator.Make(data, rules)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateAfter(t *testing.T) {
	// TODO
}

func TestValidateAfterOrEqual(t *testing.T) {
	// TODO
}

func TestValidateDate(t *testing.T) {
	// TODO
}

func TestValidateBoolean(t *testing.T) {
	// TODO
}

func TestValidateJson(t *testing.T) {
	msg := `{"name":"zhangsan", "age":18, "id":122463, "sid":122464}`

	data := make(map[string]interface{})
	data["data"] = msg
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "json"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}

func TestValidateActiveUrl(t *testing.T) {
	data := make(map[string]interface{})
	data["data"] = "http://www.asdsadasdwedswadsadwdsadsdas.com/"
	rules := make(map[string]interface{})
	rules["data"] = []string{"required", "active_url"}

	customMessages := make(map[string]interface{})
	customNames := make(map[string]string)
	v := validator.MakeAndCustom(data, rules, customMessages, customNames)
	v.Fails()
	fmt.Println(v.GetErrors())
}
