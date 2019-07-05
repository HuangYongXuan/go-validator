package validator

var Message = make(map[string]interface{})

func init() {
	Message["accepted"] = ":attr 必须接受。"
	Message["active_url"] = ":attr 不是一个有效的网址。"
	Message["after"] = ":attr 必须要晚于 :date。"
	Message["after_or_equal"] = ":attr 必须要等于 :date 或更晚。"
	Message["alpha"] = ":attr 只能由字母组成。"
	Message["alpha_dash"] = ":attr 只能由字母、数字和斜杠组成。"
	Message["alpha_num"] = ":attr 只能由字母和数字组成。"
	Message["array"] = ":attr 必须是一个数组。"
	Message["before"] = ":attr 必须要早于 :date。"
	Message["before_or_equal"] = ":attr 必须要等于 :date 或更早。"

	between := make(map[string]string)
	between["numeric"] = ":attr 必须介于 :min - :max 之间。"
	between["string"] = ":attr 必须介于 :min - :max 个字符之间。"
	between["array"] = ":attr 必须只有 :min - :max 个单元。"
	Message["between"] = between

	Message["boolean"] = ":attr 必须为布尔值。"
	Message["confirmed"] = ":attr 两次输入不一致。"
	Message["date"] = ":attr 不是一个有效的日期。"
	Message["date_format"] = ":attr 的格式必须为 :format。"
	Message["different"] = ":attr 和 :other 必须不同。"
	Message["digits"] = ":attr 必须是 :digits 位的数字。"
	Message["digits_between"] = ":attr 必须是介于 :min 和 :max 位的数字。"
	Message["email"] = ":attr 不是一个合法的邮箱。"
	Message["filled"] = ":attr 不能为空。"
	Message["in"] = "已选的属性 :attr 非法。"
	Message["in_array"] = ":attr 没有在 :other 中。"
	Message["integer"] = ":attr 必须是整数。"
	Message["ip"] = ":attr 必须是有效的 IP 地址。"
	Message["ipv4"] = ":attr 必须是有效的 IPv4 地址。"
	Message["ipv6"] = ":attr 必须是有效的 IPv6 地址。"
	Message["json"] = ":attr 必须是正确的 JSON 格式。"

	max := make(map[string]string)
	max["numeric"] = ":attr 不能大于 :max。"
	max["string"] = ":attr 不能大于 :max 个字符。"
	max["array"] = ":attr 最多只有 :max 个单元。"
	Message["max"] = max

	min := make(map[string]string)
	min["numeric"] = ":attr 必须大于等于 :min。"
	min["string"] = ":attr 至少为 :min 个字符。"
	min["array"] = ":attr 至少有 :min 个单元。"
	Message["min"] = min
	Message["not_in"] = "已选的属性 :attr 非法。"
	Message["numeric"] = ":attr 必须是一个数字。"
	Message["present"] = ":attr 必须存在。"
	Message["match"] = ":attr 格式不匹配。"
	Message["regex"] = ":attr 格式不正确。"
	Message["required"] = ":attr 不能为空。"
	Message["required_if"] = "当 :other 为 :value 时 :attr 不能为空。"
	Message["required_unless"] = "当 :other 不为 :value 时 :attr 不能为空。"
	Message["required_with"] = "当 :values 存在时 :attr 不能为空。"
	Message["required_with_all"] = "当 :values 存在时 :attr 不能为空。"
	Message["required_without"] = "当 :values 不存在时 :attr 不能为空。"
	Message["required_without_all"] = "当 :values 都不存在时 :attr 不能为空。"
	Message["same"] = ":attr 和 :other 必须相同。"
	size := make(map[string]string)
	size["numeric"] = ":attr 大小必须为 :size。"
	size["string"] = ":attr 必须是 :size 个字符。"
	size["array"] = ":attr 必须为 :size 个单元。"
	Message["size"] = size
	Message["string"] = ":attr 必须是一个字符串。"
	Message["timezone"] = "必须是一个合法的时区值。"
	Message["url"] = ":attr 格式不正确。"
}
