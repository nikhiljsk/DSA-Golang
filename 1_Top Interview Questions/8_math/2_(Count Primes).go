// Approach 1
func countPrimes(n int) int {
	sieve := make([]int, n+1)
	for i := 2; i*i <= n; i++ {
		if sieve[i] == 1 {
			continue
		}
		for j := i + i; j <= n; j += i {
			sieve[j] = 1
		}
	}
	var res int
	for i := 2; i < n; i++ {
		if sieve[i] == 0 {
			res++
		}
	}
	return res
}

// Approach 2
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	primes := make([]bool, n+1, n+1)

	result := n - 2

	for i := 2; i*i < n; i++ {
		if primes[i] == true {
			continue
		}
		for j := i * i; j < n; j += i {
			if !primes[j] {
				result--
			}
			primes[j] = true
		}
	}

	return result
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 1M
// Time Implement - 15M