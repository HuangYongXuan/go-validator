package commit

var Message = make(map[string]interface{})

func init() {
	Message["accepted"] = "必须接受。"
	Message["active_url"] = "不是一个有效的网址。"
	Message["after"] = "必须要晚于 :date。"
	Message["after_or_equal"] = "必须要等于 :date 或更晚。"
	Message["alpha"] = "只能由字母组成。"
	Message["alpha_dash"] = "只能由字母、数字和斜杠组成。"
	Message["alpha_num"] = "只能由字母和数字组成。"
	Message["array"] = "必须是一个数组。"
	Message["before"] = "必须要早于 :date。"
	Message["before_or_equal"] = "必须要等于 :date 或更早。"

	between := make(map[string]string)
	between["numeric"] = "必须介于 :min - :max 之间。"
	between["file"] = "必须介于 :min - :max kb 之间。"
	between["string"] = "必须介于 :min - :max 个字符之间。"
	between["array"] = "必须只有 :min - :max 个单元。"
	Message["between"] = between

	Message["boolean"] = "必须为布尔值。"
	Message["confirmed"] = "两次输入不一致。"
	Message["date"] = "不是一个有效的日期。"
	Message["date_format"] = "的格式必须为 :format。"
	Message["different"] = "和 :other 必须不同。"
	Message["digits"] = "必须是 :digits 位的数字。"
	Message["digits_between"] = "必须是介于 :min 和 :max 位的数字。"
	Message["dimensions"] = "图片尺寸不正确。"
	Message["distinct"] = "已经存在。"
	Message["email"] = "不是一个合法的邮箱。"
	Message["exists"] = "不存在。"
	Message["file"] = "必须是文件。"
	Message["filled"] = "不能为空。"
	Message["image"] = "必须是图片。"
	Message["in"] = "已选的属性 非法。"
	Message["in_array"] = "没有在 :other 中。"
	Message["integer"] = "必须是整数。"
	Message["ip"] = "必须是有效的 IP 地址。"
	Message["ipv4"] = "必须是有效的 IPv4 地址。"
	Message["ipv6"] = "必须是有效的 IPv6 地址。"
	Message["json"] = "必须是正确的 JSON 格式。"

	max := make(map[string]string)
	max["numeric"] = "不能大于 :max。"
	max["file"] = "不能大于 :max kb。"
	max["string"] = "不能大于 :max 个字符。"
	max["array"] = "最多只有 :max 个单元。"
	Message["max"] = max

	Message["mimes"] = "必须是一个 :values 类型的文件。"
	Message["mimetypes"] = "必须是一个 :values 类型的文件。"
	min := make(map[string]string)
	min["numeric"] = "必须大于等于 :min。"
	min["file"] = "大小不能小于 :min kb。"
	min["string"] = "至少为 :min 个字符。"
	min["array"] = "至少有 :min 个单元。"
	Message["min"] = min
	Message["not_in"] = "已选的属性 非法。"
	Message["numeric"] = "必须是一个数字。"
	Message["present"] = "必须存在。"
	Message["regex"] = "格式不正确。"
	Message["required"] = "不能为空。"
	Message["required_if"] = "当 :other 为 :value 时 不能为空。"
	Message["required_unless"] = "当 :other 不为 :value 时 不能为空。"
	Message["required_with"] = "当 :values 存在时 不能为空。"
	Message["required_with_all"] = "当 :values 存在时 不能为空。"
	Message["required_without"] = "当 :values 不存在时 不能为空。"
	Message["required_without_all"] = "当 :values 都不存在时 不能为空。"
	Message["same"] = "和 :other 必须相同。"
	size := make(map[string]string)
	size["numeric"] = "大小必须为 :size。"
	size["file"] = "大小必须为 :size kb。"
	size["string"] = "必须是 :size 个字符。"
	size["array"] = "必须为 :size 个单元。"
	Message["size"] = size
	Message["string"] = "必须是一个字符串。"
	Message["timezone"] = "必须是一个合法的时区值。"
	Message["unique"] = "已经存在。"
	Message["uploaded"] = "上传失败。"
	Message["url"] = "格式不正确。"
}
