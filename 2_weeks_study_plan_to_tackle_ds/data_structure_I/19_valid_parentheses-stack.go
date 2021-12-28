// ---- Stack implementation ----
type Stack struct {
	top      int
	elements []interface{}
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
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

func any(char byte, set []byte) bool {
	for _, c := range set {
		if c == char {
			return true
		}
	}
	return false
}

func isPair(i byte, j byte) bool {
	if (i == '{' && j == '}') || (i == '(' && j == ')') || (i == '[' && j == ']') {
		return true
	}
	return false
}

func isValid(s string) bool {
	var stack Stack
	openSet := []byte{'(', '{', '['}

	for _, char := range []byte(s) {
		if any(char, openSet) {
			stack.Push(char)
		} else {
			if !stack.isEmpty() && isPair(stack.elements[stack.top-1].(byte), char) {
				_, err := stack.Pop()
				if err != nil {
					return false
				}
			} else {
				return false
			}
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}

// implemented