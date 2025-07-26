package app

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adjList map[int]*list.List
	vertex  int
}

func (g *Graph) GraphInit(vertex int) {
	g.adjList = make(map[int]*list.List)
	g.vertex = vertex
	for i := 0; i < g.vertex; i++ {
		g.adjList[i] = list.New()
	}
}

func (g *Graph) AddEdges(i, j int) {
	_, ok := g.adjList[i]
	if !ok {
		g.adjList[i] = list.New()
	}
	g.adjList[i].PushBack(j)
}

func (g *Graph) BFS(i int) {
	queue := list.New()
	visited := make([]bool, g.vertex)
	for i := 0; i < g.vertex; i++ {
		visited[i] = false
	}
	queue.PushBack(i)
	visited[i] = true
	for queue.Len() > 0 {
		elmt := queue.Front()
		val := elmt.Value.(int)
		queue.Remove(elmt)
		fmt.Println(val)
		lst := g.adjList[val]
		for e := lst.Front(); e != nil; e = e.Next() {
			if !visited[e.Value.(int)] {
				queue.PushBack(e.Value.(int))
				visited[e.Value.(int)] = true
			}
		}
	}
}
