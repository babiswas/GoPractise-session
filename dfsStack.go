package app

import (
	"container/list"
	"fmt"
)

type StackGraph struct {
	vertex  int
	adjList map[int]*list.List
}

func (g *StackGraph) StackGraphInit(vertex int) {
	g.vertex = vertex
	g.adjList = make(map[int]*list.List)
	for i := 0; i < g.vertex; i++ {
		g.adjList[i] = list.New()
	}
}

func (g *StackGraph) AddEdges(i, j int) {
	_, ok := g.adjList[i]
	if !ok {
		g.adjList[i] = list.New()
	}
	g.adjList[i].PushBack(j)
}

func (g *StackGraph) DFSStack(i int) {
	stack := list.New()
	visited := make([]bool, g.vertex)
	for i := 0; i < g.vertex; i++ {
		visited[i] = false
	}
	stack.PushBack(i)
	for stack.Len() != 0 {
		elm := stack.Back()
		val := elm.Value.(int)
		if !visited[val] {
			fmt.Println(val)
			visited[val] = true
		}
		stack.Remove(elm)
		lst := g.adjList[val]
		for e := lst.Front(); e != nil; e = e.Next() {
			if !visited[e.Value.(int)] {
				stack.PushBack(e.Value.(int))
			}
		}
	}
}
