// Approach 1
import "math/rand"

type Solution struct {
	elements []int
}

func Constructor(nums []int) Solution {
	return Solution{elements: nums}
}

func (this *Solution) Reset() []int {
	return this.elements
}

func (this *Solution) Shuffle() []int {
	newArr := make([]int, len(this.elements))
	copy(newArr, this.elements)

	rand.Shuffle(len(newArr), func(i, j int) {
		newArr[i], newArr[j] = newArr[j], newArr[i]
	})
	return newArr
}

// Approach 2 - Fisher Yates Method of shuffling
func (this *Solution) Shuffle() []int {
	n := len(this.elements)
	newArr := make([]int, n)
	copy(newArr, this.elements)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		temp := rand.Intn(n-i) + i
		newArr[temp], newArr[i] = newArr[i], newArr[temp]
	}
	return newArr
}

// ** Metadata **
// Intuition - Yes
// Approach - Yes
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 5M
// Time Implement - 15M