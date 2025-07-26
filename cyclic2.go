package app

import "container/list"

type NewCyclic struct {
	vertex  int
	adjList map[int]*list.List
}

func (n *NewCyclic) NewCyclicInit(vertex int) {
	n.vertex = vertex
	n.adjList = make(map[int]*list.List)
	for i := 0; i < n.vertex; i++ {
		n.adjList[i] = list.New()
	}
}

func (n *NewCyclic) AddEdges(i, j int) {
	_, ok := n.adjList[i]
	if !ok {
		n.adjList[i] = list.New()
	}
	n.adjList[i].PushBack(j)
}

func (n *NewCyclic) IsCyclic() bool {
	recstack := make([]bool, n.vertex)
	visited := make([]bool, n.vertex)
	for i := 0; i < n.vertex; i++ {
		recstack[i] = false
	}
	for i := 0; i < n.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < n.vertex; i++ {
		if !visited[i] {
			if n.IsCyclicUtil(visited, i, recstack) {
				return true
			}
		}
	}
	return false
}

func (n *NewCyclic) IsCyclicUtil(visited []bool, u int, recstack []bool) bool {
	visited[u] = true
	recstack[u] = true
	lst := n.adjList[u]
	for e := lst.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		if !visited[val] {
			if n.IsCyclicUtil(visited, val, recstack) {
				return true
			}
		} else if recstack[val] {
			return true
		}
	}
	return false
}
