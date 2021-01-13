package common

type Stack struct {
	slot []interface{}
}

func NewStack() *Stack {
	return &Stack{
		slot: make([]interface{}, 0),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.slot) == 0
}

func (s *Stack) Size() int {
	return len(s.slot)
}

func (s *Stack) Pop() interface{} {
	l := len(s.slot)
	if l == 0 {
		return nil
	}

	out := s.slot[l-1]

	s.slot = s.slot[:l-1]

	return out
}

func (s *Stack) Push(value interface{}) {

	s.slot = append(s.slot, value)

}

func (s *Stack) Peek() interface{} {
	l := len(s.slot)
	if l == 0 {
		return nil
	}
	return s.slot[l-1]
}
