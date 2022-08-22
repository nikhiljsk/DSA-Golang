type MinStack struct {
	stack  []int
	minArr []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, minArr: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minArr) == 0 {
		this.minArr = append(this.minArr, val)
	} else {
		last := this.minArr[len(this.minArr)-1]
		if last < val {
			this.minArr = append(this.minArr, last)
		} else {
			this.minArr = append(this.minArr, val)
		}
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

// Approach explained
// The idea here is to maintain an array with the so-far-found-least-element.
// So if we're inserting 1, 3, 3, 5, 4
// The minArr looks      1, 1, 1, 1, 1
// OR
// If we're inserting 5, 6, 1, 3, 4
// minArr lokss       5, 5, 1, 1, 1

// ** Metadata **
// Intuition - Yes
// Approach - No
// Implementation - Yes
// Solution - Yes
// Hints - No
// Time Approach - 5M
// Time Implement - 15M