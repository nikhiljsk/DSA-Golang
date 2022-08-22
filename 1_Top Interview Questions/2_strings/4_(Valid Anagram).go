// Approach 1
func isAnagram(s string, t string) bool {
	find := make(map[rune]int)
	for _, v := range s {
		if _, ok := find[v]; ok {
			find[v] += 1
		} else {
			find[v] = 1
		}
	}
	for _, v := range t {
		if freq, ok := find[v]; ok {
			if freq == 0 {
				return false
			}
			find[v] -= 1
		} else {
			return false
		}
	}
	for _, v := range find {
		if v > 0 {
			return false
		}
	}
	return true
}

// Approach 2
func isAnagram(s string, t string) bool {
	find := make([]int, 26)
	for _, v := range s {
		find[v-'a'] += 1
	}
	for _, v := range t {
		find[v-'a'] -= 1
	}
	for _, v := range find {
		if v != 0 {
			return false
		}
	}
	return true
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 2M
// Time Implement - 15M