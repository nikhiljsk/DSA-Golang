// ---- Stack implementation ----
type Stack struct {
	top      int
	elements []interface{}
}

func (s *Stack) isEmpty() bool {
	return s == nil || len(s.elements) == 0
}

func (s *Stack) Push(v interface{}) {
	s.elements = append(s.elements, v)
	s.top = len(s.elements)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.isEmpty() {
		return -1, errors.New("empty stack")
	}
	r := s.elements[s.top-1]
	s.top--
	s.elements = s.elements[:s.top]
	return r, nil
}

// ---- Problem related code ----
var s1, s2 *Stack

type MyQueue struct {
	front int
}

func Constructor() MyQueue {
	s1 = &Stack{}
	s2 = &Stack{}
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if s1.isEmpty() {
		this.front = x
	}
	s1.Push(x)
}

func (this *MyQueue) Pop() int {
	if s2.isEmpty() {
		for !s1.isEmpty() {
			r, _ := s1.Pop()
			s2.Push(r)
		}
	}
	r, _ := s2.Pop()
	return r.(int)
}

func (this *MyQueue) Peek() int {
	if !s2.isEmpty() {
		return s2.elements[s2.top-1].(int)
	}
	return this.front
}

func (this *MyQueue) Empty() bool {
	return s1.isEmpty() && s2.isEmpty()
}

// Approach cracked. Implementation solution