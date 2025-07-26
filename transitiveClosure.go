package app

import (
	"container/list"
	"fmt"
)

type TransitiveClosure struct {
	vertex  int
	adjList map[int]*list.List
}

func (t *TransitiveClosure) TransitiveClosureInit(vertex int) {
	t.vertex = vertex
	t.adjList = make(map[int]*list.List)
	for i := 0; i < t.vertex; i++ {
		t.adjList[i] = list.New()
	}
}

func (t *TransitiveClosure) AddEdges(i, j int) {
	_, ok := t.adjList[i]
	if !ok {
		t.adjList[i] = list.New()
	}
	t.adjList[i].PushBack(j)
}

func (t *TransitiveClosure) TClosure() {
	M := make([][]int, t.vertex)
	for i := 0; i < t.vertex; i++ {
		M[i] = make([]int, t.vertex)
	}
	for i := range t.vertex {
		t.TclosureUtil(M, i, i)
	}
	fmt.Println(M)
}

func (t *TransitiveClosure) TclosureUtil(M [][]int, i, j int) {
	M[i][j] = 1
	lst := t.adjList[j]
	for e := lst.Front(); e != nil; e = e.Next() {
		if M[i][e.Value.(int)] != 1 {
			t.TclosureUtil(M, i, e.Value.(int))
		}
	}
}
