package main

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 示例 1：
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 示例 2：
// 输入：digits = ""
// 输出：[]
// 示例 3：
// 输入：digits = "2"
// 输出：["a","b","c"]
// 提示：
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。
var mapping = map[byte][]byte{
	'2': []byte{'a', 'b', 'c'},
	'3': []byte{'d', 'e', 'f'},
	'4': []byte{'g', 'h', 'i'},
	'5': []byte{'j', 'k', 'l'},
	'6': []byte{'m', 'n', 'o'},
	'7': []byte{'p', 'q', 'r', 's'},
	'8': []byte{'t', 'u', 'v'},
	'9': []byte{'w', 'x', 'y', 'z'},
}
var res []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	res = make([]string, 0)
	combine(digits, 0, []byte{})
	return res
}
func combine(digits string, index int, tmp []byte) {
	if index == len(digits) && len(tmp) == len(digits) {
		res = append(res, string(tmp))
		return
	}
	for i := 0; i < len(mapping[digits[index]]); i++ {
		tmp = append(tmp, mapping[digits[index]][i])
		combine(digits, index+1, tmp)
		tmp = tmp[:len(tmp)-1]
	}
}
