/*
Valid Anagram

Solution
Given two strings s and t, return true if t is an anagram of s, and false otherwise.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false


Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

func isAnagram(s string, t string) bool {
	seen := make(map[byte]int)
	for _, v := range []byte(s) {
		if _, ok := seen[v]; ok {
			seen[v] += 1
		} else {
			seen[v] = 1
		}
	}
	// fmt.Println(seen)
	for _, v := range []byte(t) {
		if _, ok := seen[v]; ok {
			seen[v] -= 1
			if seen[v] == 0 {
				delete(seen, v)
			}
		} else {
			return false
		}
	}
	if len(seen) == 0 {
		return true
	}
	return false
}

// Approach 2 - Using Array - Way faster!
func isAnagram(s string, t string) bool {
	seen := make([]int, 26)
	for _, v := range []byte(s) {
		seen[v-'a'] += 1
	}
	fmt.Println(seen)

	for _, v := range []byte(t) {
		seen[v-'a'] -= 1
		if seen[v-'a'] == -1 {
			return false
		}
	}
	for _, v := range seen {
		if v != 0 {
			return false
		}
	}
	return true
}

// Able to crack the approach and implement