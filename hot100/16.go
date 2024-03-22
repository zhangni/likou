package main

import (
	"math"
	"sort"
)

// 给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
// 返回这三个数的和。
//
// 假定每组输入只存在恰好一个解。
//
// 示例 1：
//
// 输入：nums = [-1,2,1,-4], target = 1
// 输出：2
// 解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
// 示例 2：
//
// 输入：nums = [0,0,0], target = 1
// 输出：0
//
// 提示：
//
// 3 <= nums.length <= 1000
// -1000 <= nums[i] <= 1000
// -104 <= target <= 104
func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] == target {
				return target
			} else if nums[i]+nums[j]+nums[k] > target {
				res = minLen(target, res, nums[i]+nums[j]+nums[k])
				k--
				for k > j && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] < target {
				res = minLen(target, res, nums[i]+nums[j]+nums[k])
				j++
				for k > j && nums[j] == nums[j-1] {
					j++
				}
			}
		}
	}
	return res
}
func minLen(target, res, tmp int) int {
	if math.Abs(float64(res-target)) > math.Abs(float64(tmp-target)) {
		return tmp
	}
	return res
}
