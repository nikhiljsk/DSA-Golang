import "unicode"

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0
	}

	isNegative := 1
	index := 0
	if s[0] == '-' {
		isNegative = -1
		index++
	} else if s[0] == '+' {
		isNegative = 1
		index++
	}

	res := 0
	max := 2147483647
	min := 2147483648
	for _, v := range s[index:] {
		if !unicode.IsDigit(v) {
			break
		}
		r := int(v) - 48
		if (res > max/10) || (res == max/10 && r > 7) || (res > min/10) || (res == min/10 && r > 8) {
			if isNegative == -1 {
				return -1 * min
			} else {
				return max
			}
		}
		res = (res * 10) + r

	}
	return res * isNegative
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 5M
// Time Implement - 30M