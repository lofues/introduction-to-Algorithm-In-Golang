package stack

import "errors"

type Stack struct {
	Data []int
	Size int
}

func (s *Stack) Init(n int) {
	s.Data = make([]int, n)
	s.Size = 0
}

func (s *Stack) Insert(x int) {
	s.Data[s.Size] = x
	s.Size++
}

func (s *Stack) Pop() (int, error) {
	if s.Size <= 0 {
		return 0, errors.New("without element")
	}
	top := s.Data[s.Size-1]
	s.Size--
	return top, nil
}

func (s *Stack) Top() (int, error) {
	if s.Size <= 0 {
		return 0, errors.New("without element")
	}
	return s.Data[s.Size-1], nil
}

func (s *Stack) Length() int {
	return s.Size
}
