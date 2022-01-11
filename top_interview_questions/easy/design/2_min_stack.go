/*
Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.


Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2


Constraints:

-231 <= val <= 231 - 1
Methods pop, top and getMin operations will always be called on non-empty stacks.
At most 3 * 104 calls will be made to push, pop, top, and getMin.
*/

type MinStack struct {
	elements []int
	top      int
	minArr   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	this.top = len(this.elements) - 1
	if len(this.minArr) != 0 {
		if this.minArr[len(this.minArr)-1] < val {
			this.minArr = append(this.minArr, this.minArr[len(this.minArr)-1])
		} else {
			this.minArr = append(this.minArr, val)
		}
	} else {
		this.minArr = append(this.minArr, val)
	}
}

func (this *MinStack) Pop() {
	fmt.Println(this.elements, this.minArr, len(this.elements))
	this.elements = this.elements[:len(this.elements)-1]
	this.minArr = this.minArr[:len(this.minArr)-1]
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minArr[len(this.minArr)-1]
}

// Had to see to push the same min element into the minArr Solution