// Approach 1
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// Approach 2, (3^19)%n==0
func isPowerOfThree(n int) bool {
	return n > 0 && int(math.Mod(math.Pow(float64(3), float64(19)), float64(n))) == 0
}

// Approach 3 - Not working with NextAfter value
func isPowerOfThree(n int) bool {
	// epsilon := math.Nextafter(1, 3) - 1
	epsilon := 0.0000000001 // This should eliminate the uncertainity
	return math.Mod(((math.Log10(float64(n))/math.Log10(3))+epsilon), 1.0) <= 2*epsilon
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 10M
// Time Implement - 30M
