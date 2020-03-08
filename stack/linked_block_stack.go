package stack

const (
	blocksize = 4
)

type LinkedBlockStack struct {
	items [blocksize]Item
	top   int
	prev  *LinkedBlockStack
}

func NewLinkedBlockStack() *LinkedBlockStack {
	return &LinkedBlockStack{
		top: -1,
	}
}

func (s *LinkedBlockStack) Empty() bool {
	return s.top == -1 && s.prev == nil
}

func (s *LinkedBlockStack) Push(item Item) {
	if s.top >= len(s.items)-1 {
		prev := NewLinkedBlockStack()
		prev.items = s.items
		prev.top = s.top
		prev.prev = s.prev
		var newitems [blocksize]Item
		s.items = newitems
		s.prev = prev
		s.top = -1
	}
	s.top++
	s.items[s.top] = item
}

func (s *LinkedBlockStack) Pop() Item {
	item := s.items[s.top]
	s.top--
	if s.top == -1 && s.prev != nil {
		s.items = s.prev.items
		s.top = s.prev.top
		s.prev = s.prev.prev
	}
	return item
}

func (s *LinkedBlockStack) Peek() Item {
	return s.items[s.top]
}

func (s *LinkedBlockStack) Traverse() [][]Item {
	res := make([][]Item, 0, 1)
	ptr := s
	for ptr != nil {
		if ptr.top >= 0 {
			res = append(res, ptr.items[0:ptr.top+1])
		}
		ptr = ptr.prev
	}
	return res
}
