package hot100

func findKthLargest(nums []int, k int) int {
	var n = len(nums)
	for i := 0; i < k; i++ {
		for j := 0; j < n-i; j++ {
			if nums[j] > nums[n-i-1] {
				nums[j], nums[n-i-1] = nums[n-i-1], nums[j]
			}
		}
	}
	return nums[n-k]
}
