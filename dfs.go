package app

import (
	"container/list"
	"fmt"
)

type DGraph struct {
	adjList map[int]*list.List
	vertex  int
}

func (g *DGraph) DGraphInit(vertex int) {
	g.vertex = vertex
	g.adjList = make(map[int]*list.List)
	for i := 0; i < g.vertex; i++ {
		g.adjList[i] = list.New()
	}
}

func (g *DGraph) AddEdges(i, j int) {
	_, ok := g.adjList[i]
	if !ok {
		g.adjList[i] = list.New()
	}
	g.adjList[i].PushBack(j)
}

func (g *DGraph) DFS() {
	visited := make([]bool, g.vertex)
	for i := 0; i < g.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < g.vertex; i++ {
		if !visited[i] {
			g.DFSUtil(visited, i)
		}
	}
}

func (g *DGraph) DFSUtil(visited []bool, i int) {
	visited[i] = true
	fmt.Println(i)
	lst := g.adjList[i]
	for e := lst.Front(); e != nil; e = e.Next() {
		if !visited[i] {
			g.DFSUtil(visited, e.Value.(int))
		}
	}
}
