package app

type Stack []int

func GetStack() Stack {
	m := make(Stack, 0)
	return m
}

func (s *Stack) PushData(num int) {
	if *s == nil {
		*s = GetStack()
	}
	*s = append(*s, num)
}

func (s *Stack) PopData() int {
	index := len(*s)
	num := (*s)[index-1]
	*s = (*s)[:index-1]
	return num
}

func (s *Stack) StackLength() int {
	return len(*s)
}
