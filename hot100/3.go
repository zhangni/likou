package main

import (
	"fmt"
)

// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//
//	请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// 提示：
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var mapIndex = make(map[byte]int, 0)
	var start, maxSize = 0, 1
	for i := 0; i < len(s); i++ {
		if k, ok := mapIndex[s[i]]; ok {
			if k >= start {
				start = k + 1
			}
		}
		mapIndex[s[i]] = i
		if maxSize < i-start+1 {
			maxSize = i - start + 1
		}
	}
	return maxSize
}

func main() {
	s := "abcabcbb"
	size := lengthOfLongestSubstring(s)
	fmt.Println(size)
}
