package app

type StackContainer []int

func GetStackContainer() StackContainer {
	s := make(StackContainer, 0)
	return s
}

func (s *StackContainer) Push(val int) {
	if *s == nil {
		*s = GetStackContainer()
	}
	*s = append(*s, val)
}

func (s *StackContainer) Length() int {
	return len(*s)
}

func (s *StackContainer) Pop() int {
	num := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return num
}
