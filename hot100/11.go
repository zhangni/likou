package main

func maxArea(height []int) int {
	var i, j = 0, len(height) - 1
	var res = 0
	for i < j {
		if height[i] <= height[j] {
			if res < (j-i)*height[i] {
				res = (j - i) * height[i]
			}
			i++
		} else {
			if res < (j-i)*height[j] {
				res = (j - i) * height[j]
			}
			j--
		}
	}
	return res
}
