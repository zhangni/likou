package main

func twoSum(nums []int, target int) []int {
	var numsIndex = make(map[int]int, 0)
	for i, k := range nums {
		if index, ok := numsIndex[target-k]; ok {
			return []int{index, i}
		}
		numsIndex[k] = i
	}
	return nil
}
