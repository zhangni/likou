package hot100

func lengthOfLongestSubstring(s string) int {
	var maxLen = 0
	var right = 0
	var byteIndex = make(map[byte]int, 0)
	for left := 0; left < len(s); left++ {
		if left != 0 {
			delete(byteIndex, s[left-1])
		}
		for right+1 < len(s) && byteIndex[s[right+1]] == 0 {
			byteIndex[s[right+1]]++
			right++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}

	}
	return maxLen
}
