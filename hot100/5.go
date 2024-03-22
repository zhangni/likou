package main

// 给你一个字符串 s，找到 s 中最长的回文子串
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, end = 0, 1
	for i := 1; i < len(s); i++ {
		//当前字符为中间节点
		j := 0
		for ; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] != s[i+j] {
				break
			}
		}
		if 2*j-1 > end-start {
			start = i - j + 1
			end = i + j
		}
		//当前字符和前一个字符相同
		if s[i-1] == s[i] {
			j = 0
			for ; i-j-1 >= 0 && i+j < len(s); j++ {
				if s[i-j-1] != s[i+j] {
					break
				}
			}
			if 2*j > end-start {
				start = i - j
				end = i + j
			}
		}
	}
	return s[start:end]
}

func main() {
	s := "cbbd"
	println(longestPalindrome(s))
}
