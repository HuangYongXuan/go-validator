package commit

import (
	"fmt"
	"reflect"
	"strings"
)

type Validator struct {
	data           map[string]interface{}
	rules          []ruleStruct
	failedRules    map[string]interface{}
	errors         map[string]interface{}
	customRules    interface{}
	customMessages interface{}
	customNames    interface{}
	customValues   interface{}
}

type ruleStruct struct {
	name  string
	rules []ruleParamsStruct
}

type ruleParamsStruct struct {
	name   string
	params interface{}
}

type RuleMethodMapsType map[string]reflect.Value

var ruleMethodMaps RuleMethodMapsType

type validatorParams struct {
	name   string
	value  interface{}
	params interface{}
}

func init() {
	if len(ruleMethodMaps) != 0 {
		return
	}
	var validator Validator
	ruleMethodMaps = make(RuleMethodMapsType, 0)
	vf := reflect.ValueOf(&validator)
	vft := vf.Type()
	mNum := vf.NumMethod()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		if strings.Contains(mName, "Validate") {
			ruleMethodMaps[mName] = vf.Method(i)
		}
	}
}

func ValidatorMake(data map[string]interface{}, rules map[string]interface{}) Validator {
	var validator Validator
	validator.data = data
	validator.rules = validator.parseRules(rules)
	return validator
}

func (c *Validator) parseRules(rules map[string]interface{}) []ruleStruct {
	var arr []ruleStruct
	for k, v := range rules {
		arr = append(arr, ruleStruct{
			name:  k,
			rules: c.parseItemRules(v.(string)),
		})
	}
	return arr
}

func (c *Validator) parseItemRules(itemRules string) []ruleParamsStruct {
	return c.parseItemRulesArray(strings.Split(itemRules, "|"))
}

func (c *Validator) parseItemRulesArray(itemRules []string) []ruleParamsStruct {
	var rules []ruleParamsStruct

	for _, v := range itemRules {
		var rule = strings.TrimSpace(v)
		args := strings.Split(rule, ":")
		var params interface{}
		if args[0] == "regex" {
			params = args[1]
		} else {
			if len(args) > 1 {
				params = strings.Split(args[1], ",")
			} else {
				params = []string{}
			}
		}
		rules = append(rules, ruleParamsStruct{
			name:   c.titleCase(args[0]),
			params: params,
		})
	}
	return rules
}

func (*Validator) titleCase(str string) string {
	return CamelString(str)
}

func (c *Validator) passes() bool {
	c.errors = make(map[string]interface{})
	c.failedRules = make(map[string]interface{})

	for _, rule := range c.rules {
		var name = rule.name
		if c.isEmptyValueAndContainsNullableRule(rule) {
			continue
		}
		for _, rules := range rule.rules {
			if rules.name == "Nullable" {
				continue
			}
			c.validate(name, rules)
		}
	}
	return false
}

func (c *Validator) isEmptyValueAndContainsNullableRule(ruleStruct ruleStruct) bool {
	value := c.getValue(ruleStruct.name)
	hasNullable := c.hasNullable(ruleStruct.rules)
	return value == nil && hasNullable
}

func (c *Validator) hasNullable(paramsStruct []ruleParamsStruct) bool {
	ok := false
	for _, r := range paramsStruct {
		if r.name == "Nullable" {
			ok = true
			break
		}
	}
	return ok
}

func (c *Validator) getValue(name string) interface{} {
	if c.data[name] == nil {
		return nil
	}
	return c.data[name]
}

func (c *Validator) Fails() bool {
	return !c.passes()
}

func (c *Validator) validate(name string, rule ruleParamsStruct) {
	value := c.getValue(name)
	method, ok := findRuleMethod(rule)
	if !ok {
		fmt.Println("验证方法 Validate" + rule.name + " 无法找到")
	} else {
		data := validatorParams{
			name:   name,
			value:  value,
			params: rule.params,
		}
		params := []reflect.Value{reflect.ValueOf(data)}
		returnValue := method.Call(params)
		OK := returnValue[0].Bool()
		if !OK {
			c.addFailure(name, rule)
		}
	}
}

func (c *Validator) addFailure(name string, rule ruleParamsStruct) {
	c.addError(name, rule)
}

func (c *Validator) addError(name string, rule ruleParamsStruct) {
	var msg = c.getErrorMessage(name, rule)
	errMsg := make([]string, 0)
	if c.hasError(name) {
		errMsg = c.errors[name].([]string)
	}
	c.errors[name] = append(errMsg, msg)
}

func (c *Validator) getErrorMessage(name string, rule ruleParamsStruct) string {
	key := SnakeString(rule.name)
	message, ok := Message[key].(string)
	if ok {
		return message
	}
	return ""
}

func (c *Validator) hasError(name string) bool {
	err := c.errors[name]
	if err == nil {
		return false
	}
	return true
}

func (c *Validator) GetErrors() map[string]interface{} {
	return c.errors
}

func findRuleMethod(rule ruleParamsStruct) (reflect.Value, bool) {
	method := ruleMethodMaps["Validate"+rule.name]
	return method, method.IsValid()
}

func (c *Validator) ValidateRequired(params validatorParams) bool {
	if params.value == nil {
		return false
	}
	val, ok := params.value.(string)
	if ok {
		val = strings.TrimSpace(val)
		if val == "" {
			return false
		}
	}

	arr, ok := params.value.([]string)
	if ok {
		if len(arr) < 1 {
			return false
		}
	}

	return true
}

func (c *Validator) ValidateMin(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateInteger(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateActiveUrl(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateMax(params validatorParams) bool {
	return true
}
