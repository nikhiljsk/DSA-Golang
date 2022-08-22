/*
Longest Substring Without Repeating Characters

Given a string s, find the length of the longest substring without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/

func isValid(seen map[byte]bool, r int, s string) bool {
	if _, ok := seen[s[r]]; ok {
		return false
	}
	return true
}

// Time - O(2N)
func lengthOfLongestSubstring(s string) int {
	l, r, count, temp := 0, 0, 0, 0
	seen := make(map[byte]bool, len(s))

	for r < len(s) {
		if isValid(seen, r, s) {
			seen[s[r]] = true
			r++
			temp++
		} else {
			for !isValid(seen, r, s) {
				delete(seen, s[l])
				l++
				temp--
			}
		}
		if count < temp {
			count = temp
		}
	}
	return count
}

/*
Testcases
"abcabcbb"
"bbbbb"
"pwwkew"
"sldga"
"a"
"aaaaaa"
""
"dvdf"
"abcdefghijklmnopqrstuvwxyz"
*/

// Implemented