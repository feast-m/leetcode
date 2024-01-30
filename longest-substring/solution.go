package longestsubstring

// bad solution
func Do(s string) int {
	set := map[byte]struct{}{}
	max := 0
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if _, ok := set[s[j]]; ok {
				newMax := len(set)
				if newMax > max {
					max = newMax
				}
				set = map[byte]struct{}{}
				break
			}
			set[s[j]] = struct{}{}
			if j == len(s)-1 && len(set) > max {
				max = len(set)
			}
		}
	}
	return max
}

func lengthOfLongestSubstring(s string) int {
	store := make(map[uint8]int)
	var left, right, result int

	for right < len(s) {
		var r = s[right]
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}
