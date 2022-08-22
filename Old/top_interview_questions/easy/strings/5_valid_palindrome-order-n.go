/*
Valid Palindrome

Solution
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

// Incase you are a pythonista! <3
// Data format is love in python. Here is a two line solution
// s = ''.join(s.translate(str.maketrans('', '', string.punctuation)).lower().split(' '))
// return s[-1::-1] == s

import (
	"strings"
	"unicode"
	"fmt"
)

func reverse(r []rune) []rune {
	n := len(r)
	for i := 0; i < n/2; i++ {
		r[i], r[n-i-1] = r[n-i-1], r[i]
	}
	return r
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := []rune{}
	for _, r := range []rune(s) {
		if unicode.IsNumber(r) || unicode.IsLetter(r) {
			n = append(n, r)
		}
	}
	k := append([]rune{}, n...)
	k = reverse(k)
	return string(n) == string(k)
}

// You could also traverse the string and directly check
func isPalindrome(s string) bool {
	st := []rune(strings.ToLower(s))
	i := 0
	j := len(st) - 1
	for i < j {
		if !unicode.IsNumber(st[i]) && !unicode.IsLetter(st[i]) {
			i++
		} else if !unicode.IsNumber(st[j]) && !unicode.IsLetter(st[j]) {
			j--
		} else if st[i] == st[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// Able to crack the approach and implemented but second one had to see the solution