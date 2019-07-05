package validator

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Validator struct {
	data           map[string]interface{}
	rules          []ruleStruct
	errors         map[string]interface{}
	customMessages map[string]interface{}
	customNames    map[string]string
}

type ruleStruct struct {
	name  string
	rules []ruleParamsStruct
}

type ruleParamsStruct struct {
	name   string
	params []string
}

type RuleMethodMapsType map[string]reflect.Value

var ruleMethodMaps RuleMethodMapsType

type validatorParams struct {
	name      string
	value     interface{}
	params    []string
	validator *Validator
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

func MakeAndCustom(data map[string]interface{}, rules map[string]interface{}, customMessages map[string]interface{}, customNames map[string]string) Validator {
	var validator Validator
	validator.data = data
	validator.rules = validator.parseRules(rules)
	validator.customMessages = customMessages
	validator.customNames = customNames
	return validator
}

func Make(data map[string]interface{}, rules map[string]interface{}) Validator {
	var validator Validator
	validator.data = data
	validator.rules = validator.parseRules(rules)
	return validator
}

func (c *Validator) parseRules(rules map[string]interface{}) []ruleStruct {
	var arr []ruleStruct
	for k, v := range rules {
		if GetInterfaceType(v) == "[]string" {
			arr = append(arr, ruleStruct{
				name:  k,
				rules: c.parseItemRulesArray(v.([]string)),
			})
		} else {
			arr = append(arr, ruleStruct{
				name:  k,
				rules: c.parseItemRules(v.(string)),
			})
		}
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
		var params []string
		if args[0] == "regex" {
			params = []string{args[1]}
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

func (c *Validator) hasData(name string) bool {
	return c.data[name] != nil
}

func (c *Validator) hasRule(name string, rules []string) bool {
	rule := c.getRule(name, rules)
	return !reflect.DeepEqual(rule, ruleStruct{})
}

func (c *Validator) getRule(name string, rulesToCheck []string) ruleStruct {
	var rule ruleStruct
	for _, value := range c.rules {
		if value.name == name {
			rule = value
			break
		}
	}
	if reflect.DeepEqual(rule, ruleStruct{}) {
		return ruleStruct{}
	}

	var rules ruleStruct
	rules.name = rule.name
	for _, value := range rule.rules {
		b, _ := StringArrayIndex(rulesToCheck, value.name)
		if b {
			rules.rules = append(rules.rules, value)
		}
	}

	return rules
}

func (c *Validator) requireParameterCount(count int, params []string, rule string) error {
	if len(params) < count {
		return errors.New("Validation rule" + rule + " requires at least " + string(count) + " parameters")
	}
	return nil
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
			name:      name,
			value:     value,
			params:    rule.params,
			validator: c,
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
	msg = c.doReplacements(msg, name, rule)
	errMsg := make([]string, 0)
	if c.hasError(name) {
		errMsg = c.errors[name].([]string)
	}
	c.errors[name] = append(errMsg, msg)
}

func (c *Validator) doReplacements(msg, name string, rule ruleParamsStruct) string {
	newMsg := strings.TrimSpace(msg)
	if newMsg == "" {
		return ""
	}
	// 获取名称映射
	attr := c.getAttr(name)
	msg = strings.ReplaceAll(msg, ":ATTR", attr)
	msg = strings.ReplaceAll(msg, ":Attr", attr)
	msg = strings.ReplaceAll(msg, ":attr", attr)

	reg, _ := regexp.Compile(":[a-zA-Z]{3,6}")

	data := reg.FindAllString(msg, 2)

	for key, value := range data {
		msg = strings.Replace(msg, value, rule.params[key], 1)
	}

	return msg
}

func (c *Validator) getAttr(name string) string {
	attr := c.customNames[name]
	fmt.Println("attr:", attr)
	if attr != "" {
		return attr
	}
	return name
}

func (c *Validator) getErrorMessage(name string, rule ruleParamsStruct) string {
	key := SnakeString(rule.name)
	message, b := c.getErrorMessages(key, name, c.customMessages)
	if b {
		return message
	}
	message, b = c.getErrorMessages(key, name, Message)
	if b {
		return message
	}
	return key + " Not error message"
}

func (c *Validator) getErrorMessages(key, name string, messages map[string]interface{}) (string, bool) {
	if GetInterfaceType(messages[key]) == "string" {
		message, ok := messages[key].(string)
		if ok {
			return message, true
		}
	} else if IsArray(messages[key]) {
		msg := messages[key].(map[string]string)
		value := c.data[name]

		if InterfaceIsInteger(value) {
			return msg["numeric"], true
		} else if InterfaceIsNumeric(value) {
			return msg["numeric"], true
		} else if IsArray(value) {
			return msg["array"], true
		} else if GetInterfaceType(value) == "string" {
			return msg["string"], true
		}
	}
	return "", false
}

func (c *Validator) hasError(name string) bool {
	err := c.errors[name]
	if err == nil {
		return false
	}
	return true
}

func (c *Validator) getError(name string) interface{} {
	return c.errors[name]
}

func (c *Validator) GetErrors() map[string]interface{} {
	return c.errors
}

func findRuleMethod(rule ruleParamsStruct) (reflect.Value, bool) {
	method := ruleMethodMaps["Validate"+rule.name]
	return method, method.IsValid()
}

func (c *Validator) ValidateSometimes(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateBail(params validatorParams) bool {
	return true
}

func (c *Validator) shouldStopValidating(name string) bool {
	if !c.hasRule(name, []string{"Bail"}) {
		return false
	}
	return c.hasError(name)
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

func (c *Validator) ValidatePresent(params validatorParams) bool {
	return c.data[params.name] != nil
}

func (c *Validator) ValidateFilled(params validatorParams) bool {
	if c.hasData(params.name) {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) anyFailingRequired(names []string) bool {
	result := false
	for _, value := range names {
		var params validatorParams
		params.name = value
		params.value = c.getValue(value)
		params.params = []string{}
		if !c.ValidateRequired(params) {
			result = true
			break
		}
	}
	return result
}

func (c *Validator) allFailingRequired(names []string) bool {
	result := true
	for _, name := range names {
		var params validatorParams
		params.name = name
		params.value = c.getValue(name)
		params.params = []string{}
		if c.ValidateRequired(params) {
			result = false
			break
		}
	}
	return result
}

func (c *Validator) ValidateRequiredWith(params validatorParams) bool {
	if c.allFailingRequired(params.params) {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) ValidateRequiredWithAll(params validatorParams) bool {
	if !c.anyFailingRequired(params.params) {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) ValidateRequiredWithout(params validatorParams) bool {
	if c.anyFailingRequired(params.params) {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) ValidateRequiredWithoutAll(params validatorParams) bool {
	if c.allFailingRequired(params.params) {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) ValidateRequiredIf(params validatorParams) bool {
	err := c.requireParameterCount(2, params.params, "required_if")
	if err != nil {
		return false
	}
	var data = c.getValue(params.params[0]).(string)

	var values = params.params[1:]

	b := false
	for _, value := range values {
		if value == data {
			b = true
		}
	}
	if b {
		return c.ValidateRequired(params)
	}

	return true
}

func (c *Validator) ValidateRequiredUnless(params validatorParams) bool {
	err := c.requireParameterCount(2, params.params, "required_unless")
	if err != nil {
		return false
	}
	var data = c.getValue(params.params[0]).(string)

	var values = params.params[1:]

	b := false
	for _, value := range values {
		if value == data {
			b = true
		}
	}
	if !b {
		return c.ValidateRequired(params)
	}
	return true
}

func (c *Validator) ValidateMatch(params validatorParams) bool {
	if params.value == nil {
		return false
	}
	var re = params.params[0]
	b, err := regexp.MatchString(re, params.value.(string))
	if err != nil {
		return false
	}
	return b
}

func (c *Validator) ValidateRegex(params validatorParams) bool {
	return c.ValidateMatch(params)
}

func (c *Validator) ValidateAccepted(params validatorParams) bool {
	if GetInterfaceType(params.value) == "bool" {
		return params.value.(bool)
	} else if GetInterfaceType(params.value) == "string" {
		value := params.value.(string)
		if value == "yes" || value == "on" || value == "1" || value == "true" {
			return true
		}
	} else if GetInterfaceType(params.value) == "int" {
		value := params.value.(int)
		if value == 1 {
			return true
		}
	}
	return false
}

func (c *Validator) ValidateArray(params validatorParams) bool {
	if IsArray(params.value) {
		return true
	}
	return false
}

func (c *Validator) ValidateConfirmed(params validatorParams) bool {
	params.params = []string{params.name + "_confirmation"}
	return c.ValidateSame(params)
}

func (c *Validator) ValidateSame(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "same")
	if err != nil {
		return false
	}

	var other = params.validator.getValue(params.params[0])

	return other != nil && params.value == other
}

func (c *Validator) ValidateDifferent(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "different")
	if err != nil {
		return false
	}

	var other = c.data[params.params[0]]

	return other != nil && params.value != other
}

func (c *Validator) ValidateDigits(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "digits")
	if err != nil {
		return false
	}
	length, _ := strconv.Atoi(params.params[0])
	str, _ := params.value.(string)
	return c.ValidateNumeric(params) && len(str) == length
}

func (c *Validator) ValidateDigitsBetween(params validatorParams) bool {
	err := c.requireParameterCount(2, params.params, "digits_between")
	if err != nil {
		return false
	}
	return true
}

func (c *Validator) ValidateSize(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "size")
	if err != nil {
		return false
	}
	return true
}

func (c *Validator) ValidateBetween(params validatorParams) bool {
	err := c.requireParameterCount(2, params.params, "between")
	if err != nil {
		return false
	}
	min, err := strconv.Atoi(params.params[0])
	max, err := strconv.Atoi(params.params[1])
	if err != nil {
		fmt.Println(err)
		return false
	}
	paramSize := c.getSize(params.name, params.value)
	return min <= paramSize && paramSize <= max
}

func (c *Validator) ValidateMin(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "min")
	if err != nil {
		return false
	}
	size, err := strconv.Atoi(params.params[0])
	if err != nil {
		fmt.Println(err)
	}

	return c.getSize(params.name, params.value) >= size
}

func (c *Validator) ValidateMax(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "max")
	if err != nil {
		return false
	}
	size, err := strconv.Atoi(params.params[0])
	if err != nil {
		fmt.Println(err)
	}

	return c.getSize(params.name, params.value) <= size
}

func (c *Validator) getSize(name string, value interface{}) int {
	if InterfaceIsInteger(value) {
		return value.(int)
	} else if IsArray(value) {
		size := 0
		switch GetInterfaceType(value) {
		case "[]string":
			size = len(value.([]string))
		case "[]int":
			size = len(value.([]int))
		}
		return size
	} else if GetInterfaceType(value) == "string" {
		return len(value.(string))
	}
	return 0
}

func (c *Validator) ValidateIn(params validatorParams) bool {
	value, err := params.value.(string)
	if !err {
		return false
	}

	b, _ := StringArrayIndex(params.params, value)
	return b
}

func (c *Validator) ValidateNotIn(params validatorParams) bool {
	return !c.ValidateIn(params)
}

func (c *Validator) ValidateNumeric(params validatorParams) bool {
	return InterfaceIsNumeric(params.value)
}

func (c *Validator) ValidateInteger(params validatorParams) bool {
	return InterfaceIsInteger(params.value)
}

func (c *Validator) ValidateString(params validatorParams) bool {
	return GetInterfaceType(params.value) == "string"
}

func (c *Validator) ValidateEmail(params validatorParams) bool {
	params.params = []string{`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`}
	return c.ValidateMatch(params)
}

func (c *Validator) ValidateIp(params validatorParams) bool {
	ipv4, b := params.value.(string)
	if !b {
		return false
	}

	return net.ParseIP(ipv4) != nil
}

func (c *Validator) ValidateUrl(params validatorParams) bool {
	params.params = []string{`(https?|ftp):\/\/[^\s\/$.?#].[^\s]*`}
	return c.ValidateMatch(params)
}

func (c *Validator) ValidateAlpha(params validatorParams) bool {
	params.params = []string{`^[\pL\pM]+$`}
	return c.ValidateMatch(params)
}

func (c *Validator) ValidateAlphaNum(params validatorParams) bool {
	params.params = []string{`^[\pL\pM\pN]+$`}
	return c.ValidateMatch(params)
}

/**
验证属性是否仅包含字母数字字符，短划线和下划线。
*/
func (c *Validator) ValidateAlphaDash(params validatorParams) bool {
	params.params = []string{`^[\pL\pM\pN_-]+$`}
	return c.ValidateMatch(params)
}

func (c *Validator) ValidateBefore(params validatorParams) bool {
	err := c.requireParameterCount(1, params.params, "before")
	if err != nil {
		return false
	}
	return true
}

func (c *Validator) ValidateBeforeOrEqual(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateAfter(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateAfterOrEqual(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateDate(params validatorParams) bool {
	return true
}

func (c *Validator) ValidateBoolean(params validatorParams) bool {

	if GetInterfaceType(params.value) == "bool" {
		return true
	}
	if InterfaceIsInteger(params.value) {
		num := params.value.(int)
		return num == 0 || num == 1
	}
	if GetInterfaceType(params.value) == "string" {
		str := params.value.(string)
		b, _ := StringArrayIndex([]string{"0", "1"}, str)
		return b
	}
	return false
}

func (c *Validator) ValidateJson(params validatorParams) bool {
	if GetInterfaceType(params.value) != "string" {
		return false
	}
	son := make(map[string]interface{})
	err := json.Unmarshal([]byte(params.value.(string)), &son)
	if err != nil {
		return false
	}
	return true
}

func (c *Validator) ValidateActiveUrl(params validatorParams) bool {
	if GetInterfaceType(params.value) != "string" {
		return false
	}
	webUrl := params.value.(string)

	urls, err := url.Parse(webUrl)
	if err != nil {
		return false
	}
	ns, err := net.LookupHost(urls.Host)
	if err != nil {
		return false
	}
	return len(ns) > 0
}