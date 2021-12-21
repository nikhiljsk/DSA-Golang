/*
Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	n := 201
	for _, i := range strs {
		if n > len(i) {
			n = len(i)
		}
	}

	prefix := ""
	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i] {
				return prefix
			}
		}
		prefix += string(strs[0][i])
	}
	return prefix
}

// Multiple other approaches defined in solution. Need to check
// Able to crack the approach and implemented