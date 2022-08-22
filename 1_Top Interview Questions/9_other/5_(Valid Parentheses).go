func isPair(i, j byte) bool {
	if (i == '{' && j == '}') || (i == '(' && j == ')') || (i == '[' && j == ']') {
		return true
	}
	return false
}

func isValid(s string) bool {
	var stack []byte
	var top int
	stack = append(stack, s[0])
	for i := 1; i < len(s); i++ {
		if top >= 0 && isPair(stack[top], s[i]) {
			stack = stack[:top]
			top--
			continue
		}
		stack = append(stack, s[i])
		top++
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - No
// Hints - No
// Time Approach - 5M
// Time Implement - 15M