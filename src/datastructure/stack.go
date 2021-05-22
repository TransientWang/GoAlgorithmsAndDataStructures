package datastructure

import (
	"errors"
	"fmt"
)

type Stack struct {
	l   []interface{}
	Len int
}

func (s *Stack) Push(val interface{}) error {
	if len(s.l) == s.Len {
		return fmt.Errorf("stack over flow ,len %d", s.Len)
	}
	s.l = append(s.l, val)
	return nil
}

func (s *Stack) Pop() (val interface{}, err error) {
	if len(s.l) == 0 {
		return nil, errors.New("stack empty")
	}
	val = s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return
}


