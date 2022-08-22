/*
Count Primes
Given an integer n, return the number of prime numbers that are strictly less than n.



Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0


Constraints:

0 <= n <= 5 * 106
*/

import "math"

// Approach 1 - Naive - Timeout
func isPrime(v int) bool {
	i := 2
	for float64(i) <= math.Sqrt(float64(v)) {
		if v%i == 0 {
			return false
		}
		i++
	}
	return true
}

func countPrimes(n int) int {
	var c int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			c++
		}
	}
	return c
}

// Approach 2
// Sieve of Eratosthenes
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}

	sieve := make([]bool, n)
	for i := 2; i < n; i++ {
		sieve[i] = true
	}

	for i := 2; i*i < n; i++ {
		for j := i * i; j < n; j += i {
			if !sieve[i] {
				continue
			} else {
				sieve[j] = false
			}
		}
	}

	var c int
	for _, v := range sieve {
		if v {
			c++
		}
	}
	return c
}

// Able to crack both the approaches. But had to look at solution (hints) for optimizations