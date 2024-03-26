package main

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	var i = 0
	var j = len(nums) - 1
	for {
		if i > j {
			break
		}
		if nums[i] == val {
			if nums[j] == val {
				j--
			} else {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		} else {
			i++
		}
	}
	nums = nums[:i]
	return i
}
