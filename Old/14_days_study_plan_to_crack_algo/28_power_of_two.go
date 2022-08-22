/*
Power of Two
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.



Example 1:

Input: n = 1
Output: true
Explanation: 20 = 1
Example 2:

Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false


Constraints:

-231 <= n <= 231 - 1


Follow up: Could you solve it without loops/recursion?
*/

// Approach 1
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	var count int
	for n != 0 {
		if n&1 == 1 {
			count++
		}
		if count > 1 {
			return false
		}
		n >>= 1
	}
	return true
}

// Approach 2
func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// Approach 3
// Negative number are stored as: (Flip all values of absolute and then Add 1 bit to it)
// 4 & -4 == 4
// 100 & (011+1) -> 100 & 100 == 100 // True
// 4 & -3
// 100 & (100+1) -> 100 & 101 != 100 // False
func isPowerOfTwo(n int) bool {
	return n > 0 && n&-n == n
}

// Implemented approach 1
// Solution approach 2 and 3