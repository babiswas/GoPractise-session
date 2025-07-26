package app

import "container/list"

type Cyclic struct {
	vertex  int
	adjList map[int]*list.List
}

func (c *Cyclic) CyclicInit(vertex int) {
	c.vertex = vertex
	c.adjList = make(map[int]*list.List)
	for i := 0; i < c.vertex; i++ {
		c.adjList[i] = list.New()
	}
}

func (c *Cyclic) AddEdges(i, j int) {
	_, ok := c.adjList[i]
	if !ok {
		c.adjList[i] = list.New()
	}
	c.adjList[i].PushBack(j)
	_, ok = c.adjList[j]
	if !ok {
		c.adjList[i] = list.New()
	}
	c.adjList[j].PushBack(i)
}

func (c *Cyclic) IsCyclicUtil(visited []bool, u int, parent int) bool {
	visited[u] = true
	lst := c.adjList[u]
	for e := lst.Front(); e != nil; e = e.Next() {
		if !visited[e.Value.(int)] {
			if c.IsCyclicUtil(visited, e.Value.(int), u) {
				return true
			} else if e.Value.(int) != parent {
				return true
			}
		}
	}
	return false
}

func (c *Cyclic) IsCyclic() bool {
	visited := make([]bool, c.vertex)
	for i := 0; i < c.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < c.vertex; i++ {
		if !visited[i] {
			if c.IsCyclicUtil(visited, i, -1) {
				return true
			}
		}
	}
	return false
}
