package main

import (
	"fmt"
)

//罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
//字符          数值
//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000
//例如， 罗马数字 2 写做 II ，即为两个并列的 1 。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
//
//通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
//
//I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//给定一个罗马数字，将其转换成整数。

func romanToInt(s string) int {
	var mapping = map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	var result int
	for i := 0; i < len(s); i++ {
		if s[i] == 'X' {
			if i+1 < len(s) && s[i+1] == 'L' {
				result += 40
				i++
			} else if i+1 < len(s) && s[i+1] == 'C' {
				result += 90
				i++
			} else {
				result += 10
			}
		} else if s[i] == 'I' {
			if i+1 < len(s) && s[i+1] == 'V' {
				result += 4
				i++
			} else if i+1 < len(s) && s[i+1] == 'X' {
				result += 9
				i++
			} else {
				result += 1
			}
		} else if s[i] == 'C' {
			if i+1 < len(s) && s[i+1] == 'D' {
				result += 400
				i++
			} else if i+1 < len(s) && s[i+1] == 'M' {
				result += 900
				i++
			} else {
				result += 100
			}
		} else {
			result += mapping[s[i]]
		}
		fmt.Println(result)
	}
	return result
}
func main() {
	println(romanToInt("DCXXI"))
}
