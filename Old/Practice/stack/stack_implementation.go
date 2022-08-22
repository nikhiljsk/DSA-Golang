package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top      int
	elements []interface{}
}

func (s *Stack) isEmpty() bool {
	return len(s.elements) == 0
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

func main() {
	var s Stack
	s.Push(10)
	s.Push("11")
	s.Push(12)
	fmt.Printf("Here is the stack %v and top is pointing at index: %v\n", s, s.top-1)
	for i := 0; i < 4; i++ {
		r, err := s.Pop()
		if err != nil {
			fmt.Println("ERROR!", err)
		}
		fmt.Printf("Popped value: %v Stack: %v Top at index: %v\n", r, s, s.top-1)
	}
	fmt.Println("-- The End --")
}
