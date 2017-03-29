package main

import (
	"errors"
	"fmt"
	"sync"
)

type stackslice []int

func (s stackslice) Push(v int) stackslice {
	return append(s, v)
}

func (s stackslice) Pop() (stackslice, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}

type stack struct {
	lock sync.Mutex
	s    []int
}

func NewStack() *stack {
	return &stack{sync.Mutex{}, []int{}}
}

func (s *stack) Push(v int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, v)
}

func (s *stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	l := len(s.s)
	if l <= 0 {
		return 0, errors.New("Stack is empty")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func main() {
	//Implementation 1
	ss := make(stackslice, 0)
	ss = ss.Push(1)
	ss = ss.Push(2)
	ss = ss.Push(3)
	ss, v := ss.Pop()
	fmt.Println(v)
	ss, v = ss.Pop()
	fmt.Println(v)
	ss, v = ss.Pop()
	fmt.Println(v)

	//Implementation 2
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
