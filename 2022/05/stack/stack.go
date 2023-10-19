package stack

type Stack []byte

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Pop() byte {
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}

func (s Stack) Peek() byte {
	return s[len(s)-1]
}

func (s *Stack) Reverse() {
	var output Stack

	for len(*s) > 0 {
		output.Push(s.Pop())
	}

	*s = output
}
