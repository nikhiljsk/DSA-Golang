// Approach 1 - No use of built-in functions
func isAlphaNumeric(b byte) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	} else if b >= '0' && b <= '9' {
		return true
	} else {
		return false
	}
}

func isEqual(a, b byte) bool {
	if (a >= 'a' && a <= 'z') && (b >= 'A' && b <= 'Z') {
		return a == (b + 32)
	}
	if (b >= 'a' && b <= 'z') && (a >= 'A' && a <= 'Z') {
		return (a + 32) == b
	}
	return a == b
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if !isAlphaNumeric(s[start]) {
			start++
			continue
		}
		if !isAlphaNumeric(s[end]) {
			end--
			continue
		}
		if isEqual(s[start], s[end]) {
			start++
			end--
			continue
		}
		return false
	}
	return true
}

// Approach 2
func reverse(r []rune) []rune {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return r
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := []rune{}
	for _, r := range []rune(s) {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			n = append(n, r)
		}
	}
	k := append([]rune{}, n...)
	k = reverse(k)
	return string(n) == string(k)
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 2M
// Time Implement - 20M