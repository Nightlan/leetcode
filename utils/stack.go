package utils

import (
	"strconv"
	"strings"
)

type Stack struct {
	buf []int
	sum int
}

func NewStack(cap int)*Stack{
	return &Stack{
		buf:make([]int,0, cap),
		sum: 0,
	}
}

func (s *Stack)Push(value int){
	s.buf = append(s.buf, value)
	s.sum += value
}

func (s *Stack)Pop()int{
	if len(s.buf) == 0 {
		return -1
	}
	value := s.buf[len(s.buf)-1]
	s.buf = s.buf[:len(s.buf)-1]
	s.sum -= value
	return value
}

func (s *Stack)Top()int{
	if len(s.buf) == 0 {
		return -1
	}
	return s.buf[len(s.buf)-1]
}

func (s *Stack)Sum()int{
	return s.sum
}

func (s *Stack)String()string{
	var build strings.Builder
	for i := range s.buf {
		build.WriteString(strconv.Itoa(s.buf[i]))
		if i != len(s.buf)-1 {
			build.WriteString("-")
		}
	}
	return build.String()
}
