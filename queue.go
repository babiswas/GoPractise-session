package app

type Queue []int

func GetQueue() Queue {
	m := make(Queue, 0)
	return m
}

func (q *Queue) Enqueue(num int) {
	if *q == nil {
		*q = make(Queue, 0)
	}
	*q = append(*q, num)
}

func (q *Queue) Dequeue() int {
	if len(*q) == 0 {
		return -1
	} else if len(*q) == 1 {
		num := (*q)[0]
		*q = (*q)[:0]
		return num
	} else {
		num := (*q)[0]
		*q = (*q)[1:]
		return num
	}
}

func (q *Queue) QueueLength() int {
	return len(*q)
}
