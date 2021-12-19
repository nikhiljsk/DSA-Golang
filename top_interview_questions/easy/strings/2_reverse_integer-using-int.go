/*
Reverse Integer

Solution
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).



Example 1:

Input: x = 123
Output: 321
Example 2:

Input: x = -123
Output: -321
Example 3:

Input: x = 120
Output: 21


Constraints:

-231 <= x <= 231 - 1
*/

func reverse(x int) int {
	res := 0
	max := 2147483647
	min := -2147483648
	for x != 0 {
		r := x % 10
		if (res > max/10) || (res == max/10 && r > 7) {
			return 0
		}
		if (res < min/10) || (res == min/10 && r < -8) {
			return 0
		}
		x /= 10
		res = (res * 10) + r
	}
	return res
}

// Able to crack the approach and implement half way through. But the edge cases got me, had to see solution