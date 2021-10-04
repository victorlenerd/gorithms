package misc

type Stack struct {
	data []string
}

func (s *Stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *Stack) pop() string {
	top := len(s.data) - 1
	val := s.data[top]
	s.data = s.data[:top]
	return val
}

func (s *Stack) top() string {
	top := len(s.data) - 1
	val := s.data[top]
	return val
}

func (s *Stack) len() int {
	return len(s.data)
}

type StackInt struct {
	data []int
}

func (s *StackInt) push(str int) {
	s.data = append(s.data, str)
}

func (s *StackInt) pop() int {
	top := len(s.data) - 1
	val := s.data[top]
	s.data = s.data[:top]
	return val
}

func (s *StackInt) top() int {
	top := len(s.data) - 1
	val := s.data[top]
	return val
}

func (s *StackInt) len() int {
	return len(s.data)
}
