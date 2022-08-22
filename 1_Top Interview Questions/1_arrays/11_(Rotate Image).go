// Approach 1
// Refer this to understand this: https://dev.to/seanpgallivan/solution-rotate-image-cpp
func rotate(matrix [][]int) {
	// The four element swap in a ring approach
	n := len(matrix)
	depth := n / 2 // Since only floor of n/2 rings exists, for odd matrix lenght middle element is not rotated
	for i := 0; i < depth; i++ {
		opp := n - i - 1        // Opposite side of i
		jLen := n - (2 * i) - 1 // Since for every ring the dimention is reduced by 2 and last element of row is already swapped by first element
		for j := 0; j < jLen; j++ {
			temp := matrix[i][i+j]
			matrix[i][i+j] = matrix[opp-j][i]
			matrix[opp-j][i] = matrix[opp][opp-j]
			matrix[opp][opp-j] = matrix[i+j][opp]
			matrix[i+j][opp] = temp
		}
	}
}

// Approach 2
func rotate(matrix [][]int) {
	n := len(matrix)
	depth := n / 2
	t, b, l, r := 0, n-1, 0, n-1
	for i := 0; i < depth; i++ {
		jLen := n - (2 * i) - 1
		for j := 0; j < jLen; j++ {
			temp := matrix[t][l+j]
			matrix[t][l+j] = matrix[b-j][l]
			matrix[b-j][l] = matrix[b][r-j]
			matrix[b][r-j] = matrix[t+j][r]
			matrix[t+j][r] = temp
		}
		r--
		l++
		b--
		t++
	}
}

// ** Metadata **
// Intuition - Yes
// Approach - No
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 20M
// Time Implement - 1H
