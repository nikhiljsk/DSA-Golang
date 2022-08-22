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

import "strconv"

func reverse(x int) int {
	neg := false
	if x < 0 {
		x *= -1
		neg = true
	}

	s := []byte(strconv.Itoa(x))
	till := len(s) / 2
	for i := 0; i < till; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
	si := strings.Trim(string(s), "0")
	fmt.Println(si)
	v, _ := strconv.ParseInt(si, 10, 64)

	if int(v) > 2147483647 || int(v) < -2147483648 {
		return 0
	}

	if neg {
		return -1 * int(v)
	}
	return int(v)
}

// Able to crack the approach and implement