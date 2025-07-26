package app

import (
	"container/list"
	"fmt"
)

type TopSort struct {
	vertex  int
	adjList map[int]*list.List
}

func (t *TopSort) TopSortInit(vertex int) {
	t.vertex = vertex
	t.adjList = make(map[int]*list.List, vertex)
	for i := 0; i < t.vertex; i++ {
		t.adjList[i] = list.New()
	}
}

func (t *TopSort) AddEdges(i, j int) {
	_, ok := t.adjList[i]
	if !ok {
		t.adjList[i] = list.New()
	}
	t.adjList[i].PushBack(j)
}

func (t *TopSort) TpSort() {
	visited := make([]bool, t.vertex)
	for i := 0; i < t.vertex; i++ {
		visited[i] = false
	}
	stack := list.New()
	for i := 0; i < t.vertex; i++ {
		if !visited[i] {
			t.TopSortUtil(visited, stack, i)
		}
	}
	for e := stack.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}

func (t *TopSort) TopSortUtil(visited []bool, stack *list.List, u int) {
	visited[u] = true
	lst := t.adjList[u]
	for e := lst.Front(); e != nil; e = e.Prev() {
		if !visited[e.Value.(int)] {
			t.TopSortUtil(visited, stack, e.Value.(int))
		}
	}
	stack.PushBack(u)
}
