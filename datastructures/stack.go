package datastructures

// The Stack datastructure can also be
// implemented using a linked list

// Stack datastructure
type Stack struct {
	items []interface{}
}

// Push add an item to a stack
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop removes the last item from the stack
func (s *Stack) Pop() interface{} {
	if !s.isEmpty() {
		item := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return item
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if s.isEmpty() {
		panic("empty stack!")
	}
	return s.items[len(s.items)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.items) < 1
}
