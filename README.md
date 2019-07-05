# go-validator

### 数据验证
* 使用

```$go
    import "github.com/Ysll233/go-validator/validator"
    
    data := make(map[string]interface{})
    data["ID"] = 1
    data["Name"] = "lin han"
    data["Disabled"] = true
    data["Money"] = 10.65
    
    rules := make(map[string]interface{})
    rules["ID"] = "required|min:0|integer"
    rules["Name"] = "required|min:0|integer|active_url"
    rules["age"] = "required|integer|min:0|max:120"
    
    validator := validate.Make(data, rules)
    if validator.Fails() {
        fmt.Println(validator.GetErrors())
    }
```

* 自定义错误消息

```$go
    import "github.com/Ysll233/go-validator/validator"
    
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
```

* 自定义验证字段名称
```$go
    import "github.com/Ysll233/go-validator/validator"
    
    data := make(map[string]interface{})
    data["name"] = "1890000000"
    
    rules := make(map[string]interface{})
    rules["name"] = []string{"nullable", "match:^1([38][0-9]|14[57]|5[^4])\\d{8}$"}
    
    customMessages := make(map[string]interface{})
    customMessages["match"] = "格式不正确"
    customNames := make(map[string]string)
    customNames["name"] = "名称"
    v := validator.MakeAndCustom(data, rules, customMessages, customNames)
    v.Fails()
    fmt.Println(v.GetErrors())
```