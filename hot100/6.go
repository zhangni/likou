package main

import (
	"fmt"
)

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
// 请你实现这个将字符串进行指定行数变换的函数：
// string convert(string s, int numRows);
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// 示例 3：
// 输入：s = "A", numRows = 1
// 输出："A"
func convert(s string, numRows int) string {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	var zunumlen = numRows*2 - 2
	var res = make([]byte, 0)
	for i := 0; i < numRows; i++ {
		for j := 0; j <= len(s)/zunumlen; j++ {
			if zunumlen*j+i < len(s) {
				res = append(res, s[zunumlen*j+i])
			}
			if i != 0 && i != zunumlen/2 && zunumlen*(j+1)-i < len(s) {
				res = append(res, s[zunumlen*(j+1)-i])
			}
		}
	}
	return string(res)
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows))
}
