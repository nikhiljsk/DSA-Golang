/*
Permutation in String
Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.

In other words, return true if one of s1's permutations is the substring of s2.

 

Example 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
 

Constraints:

1 <= s1.length, s2.length <= 104
s1 and s2 consist of lowercase English letters.
*/

import "reflect"

func compareFreqEqual(s1, s2 string) bool {
    m1, m2 := make(map[rune]int, len(s1)), make(map[rune]int, len(s2))
    for _, v := range s1{
        m1[v] += 1
    }
    for _, v := range s2{
        m2[v] += 1
    }
    return reflect.DeepEqual(m1, m2)
}

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    n, m := len(s1), len(s2)
    for i:=0; i<m-n+1; i++{
        if compareFreqEqual(s1, s2[i:i+n]) {
            return true
        }
    }
    return false
}

// Approach 2
import "reflect"

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    
    n, m := len(s1), len(s2)
    m1, m2 := make(map[byte]int, len(s1)), make(map[byte]int, len(s1))

    for i, _ := range s1{
        m1[s1[i]] += 1 // Calculate frequency of chars in s1
    }
    for i, _ := range s2[:n-1]{
        m2[s2[i]] += 1 // Calculate frequency of chars in s2 till n chars
    }
    
    for i:=n-1; i<m; i++{ // Window method
        m2[s2[i]] += 1
        if reflect.DeepEqual(m1, m2) {
            return true
        }
        
        if m2[s2[i-n+1]]-1 == 0 { // Delete the freq of first char in the window
            delete(m2, s2[i-n+1])
        } else {
            m2[s2[i-n+1]] -= 1
        }
    }
    return false
}

// Approach 3
import "reflect"

func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    n, m := len(s1), len(s2)
    m1, m2 := make([]int, 26), make([]int, 26)

    for i, _ := range s1{
        m1[s1[i]-'a'] += 1 // Calculate frequency of chars in s1
    }
    for i, _ := range s2[:n-1]{
        m2[s2[i]-'a'] += 1 // Calculate frequency of chars in s2 till n chars
    }
    
    for i:=n-1; i<m; i++{ // Window method
        m2[s2[i]-'a'] += 1
        if reflect.DeepEqual(m1, m2) {
            return true
        }
        m2[s2[i-n+1]-'a'] -= 1 // Delete the freq of first char in the window
    }
    return false
}

// Approach 4 - Better implementation of above method
func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    n, m := len(s1), len(s2)
    m1, m2 := make([]int, 26), make([]int, 26)
    for i:=0; i<n; i++ {
        m1[s1[i]-'a'] += 1      // Calculate frequency of chars in s1
        m2[s2[i]-'a'] += 1      // Calculate frequency of chars in s2 till n chars
    }
    
    i:=n
    for { // Window method
        if reflect.DeepEqual(m1, m2) {
            return true
        }
        if i >= m {
			return false
		}
        m2[s2[i-n]-'a']--      // Delete the freq of first char in the window
        m2[s2[i]-'a']++        // New char freq
        i++
    }
    return false
}

/*
Testcases
"adc"
"dcda"
"ab"
"eidbaooo"
"ab"
"eidboaoo"
"aaa"
"a"
"a"
"aaa"
"ab"
"ba"
"ab"
"ab"
"aabss"
"ssbaa"
*/

// Implemented 
// Had to see the solution for concept, Number of chars == Perm number of chars