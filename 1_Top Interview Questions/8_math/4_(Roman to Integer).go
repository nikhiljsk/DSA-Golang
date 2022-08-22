func romanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var prev, res int
	for i := len(s) - 1; i >= 0; i-- {
		temp := s[i]
		if prev > roman[temp] {
			res -= roman[temp]
		} else {
			res += roman[temp]
		}
		prev = roman[temp]
	}
	return res
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 10M
// Time Implement - 20M