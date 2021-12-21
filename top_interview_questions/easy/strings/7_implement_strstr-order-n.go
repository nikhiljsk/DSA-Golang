/*
Implement strStr()
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().



Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Example 3:

Input: haystack = "", needle = ""
Output: 0


Constraints:

0 <= haystack.length, needle.length <= 5 * 104
haystack and needle consist of only lower-case English characters.

*/

func strStr(haystack string, needle string) int {
	n := len(needle)
	h := len(haystack)
	if n == 0 {
		return 0
	}

	for i := 0; i < h-n+1; i++ {
		for j := 0; j < n; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == n-1 {
				return i
			}
		}
	}
	return -1
}

// Approach two
func strStr(haystack string, needle string) int {
	n, h := len(needle), len(haystack)

	for i := 0; i <= h-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}
	return -1
}

// Had to see the solution to get the approaches above