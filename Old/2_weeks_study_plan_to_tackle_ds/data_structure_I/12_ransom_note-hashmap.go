/*
383. Ransom Note

Given two stings ransomNote and magazine, return true if ransomNote can be constructed from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.

*/

func canConstruct(ransomNote string, magazine string) bool {
	seen := make(map[byte]int)
	for _, v := range []byte(ransomNote) {
		if _, ok := seen[v]; ok {
			seen[v] += 1
		} else {
			seen[v] = 1
		}
	}
	// fmt.Println(seen)
	for _, v := range []byte(magazine) {
		if _, ok := seen[v]; ok {
			seen[v] -= 1
			if seen[v] == 0 {
				delete(seen, v)
			}
		}
	}
	if len(seen) == 0 {
		return true
	}
	return false
}

// Implemented