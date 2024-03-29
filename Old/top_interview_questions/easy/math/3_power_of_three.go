/*
 Power of Three

Given an integer n, return true if it is a power of three. Otherwise, return false.

An integer n is a power of three, if there exists an integer x such that n == 3x.



Example 1:

Input: n = 27
Output: true
Example 2:

Input: n = 0
Output: false
Example 3:

Input: n = 9
Output: true


Constraints:

-231 <= n <= 231 - 1


Follow up: Could you solve it without loops/recursion?
*/
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n != 1 && n != -1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}

// Approach 2
// https://leetcode.com/problems/power-of-three/solution/
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

// Approach 3 - Not working with NextAfter value
func isPowerOfThree(n int) bool {
	// epsilon := math.Nextafter(1, 3) - 1
	epsilon := 0.0000000001 // This should eliminate the uncertainity
	return math.Mod(((math.Log10(float64(n))/math.Log10(3))+epsilon), 1.0) <= 2*epsilon
}

// Implemented. Solution for second approach