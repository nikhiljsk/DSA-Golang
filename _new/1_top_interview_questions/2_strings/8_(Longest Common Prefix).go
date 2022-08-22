func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	n := 201
	for _, v := range strs {
		if n > len(v) {
			n = len(v)
		}
	}

	var res []byte
	for i := 0; i < n; i++ {
		for j := 0; j < len(strs)-1; j++ {
			if strs[j][i] != strs[j+1][i] {
				return string(res)
			}
		}
		res = append(res, strs[0][i])
	}
	return string(res)
}

// ** Metadata **
// Intuition - Yes
// Approach - No
// Implementation - No
// Solution - Yes
// Hints - No
// Time Approach - 10M  (Got confused by thinking the  solution is asking for longest substring)
// Time Implement - 30M