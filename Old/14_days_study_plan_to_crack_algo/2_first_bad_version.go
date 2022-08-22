/*
First Bad Version
You are a product manager and currently leading a team to develop a new product. Unfortunately, the latest version of your product fails the quality check. Since each version is developed based on the previous version, all the versions after a bad version are also bad.

Suppose you have n versions [1, 2, ..., n] and you want to find out the first bad one, which causes all the following ones to be bad.

You are given an API bool isBadVersion(version) which returns whether version is bad. Implement a function to find the first bad version. You should minimize the number of calls to the API.



Example 1:

Input: n = 5, bad = 4
Output: 4
Explanation:
call isBadVersion(3) -> false
call isBadVersion(5) -> true
call isBadVersion(4) -> true
Then 4 is the first bad version.
Example 2:

Input: n = 1, bad = 1
Output: 1


Constraints:

1 <= bad <= n <= 231 - 1
*/

func firstBadVersion(n int) int {
	l, r := 1, n
	var mid int

	for l <= r {
		mid = l + (r-l)/2
		if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l // OR r+1
}

// What overflow condition?
// https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html
// Why mid = l+(r-l)/2
// https://qr.ae/pG6yGi
// We know mid is somewhere after low, so:
// mid = low + x
// ( high + low ) / 2 = low + x // substituting mid
// low + x = ( high + low ) / 2
// x = ( high + low ) / 2 - low
// x = ( high + low - 2 * low ) / 2
// x = ( high - low ) / 2

// Implemented, but had to see the solution for l+(r-l)/2
