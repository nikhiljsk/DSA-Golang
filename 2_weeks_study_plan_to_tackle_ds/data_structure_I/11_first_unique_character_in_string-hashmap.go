/*
First Unique Character in a String

Solution
Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.



Example 1:

Input: s = "leetcode"
Output: 0
Example 2:

Input: s = "loveleetcode"
Output: 2
Example 3:

Input: s = "aabb"
Output: -1


Constraints:

1 <= s.length <= 105
s consists of only lowercase English letters.
*/

func firstUniqChar(s string) int {
	seen := make(map[byte]int)
	for _, v := range []byte(s) {
		if _, ok := seen[v]; ok {
			seen[v] += 1
		} else {
			seen[v] = 1
		}
	}
	for i, v := range []byte(s) {
		if _, ok := seen[v]; ok && seen[v] == 1 {
			return i
		}
	}
	return -1
}

// You can also use an array instead of hashmap to store the freq. This could be even faster I believe
// The array indices will be from 0-25, and a will be 0 while z will be 25. So O(1) lookup time!

// Able to crack the approach and implement