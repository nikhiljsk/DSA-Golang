func firstUniqChar(s string) int {
	find := make(map[rune]int)
	for _, v := range s {
		if _, ok := find[v]; ok {
			find[v] += 1
		} else {
			find[v] = 1
		}
	}
	for i, v := range s {
		if find[v] == 1 {
			return i
		}
	}
	return -1
}

// Other approaches
// You can also use an array instead of hashmap to store the freq. This could be even faster I believe
// The array indices will be from 0-25, and a will be 0 while z will be 25. So O(1) lookup time!

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No (Forgot r>7<-8 edge case)
// Solution - No
// Hints - No
// Time Approach - 1M
// Time Implement - 2M