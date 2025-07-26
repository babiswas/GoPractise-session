package app

import (
	"container/list"
)

type MotherVertex struct {
	vertex  int
	adjList map[int]*list.List
}

func (m *MotherVertex) MotherVertexInit(vertex int) {
	m.vertex = vertex
	m.adjList = make(map[int]*list.List)
	for i := 0; i < m.vertex; i++ {
		m.adjList[i] = list.New()
	}
}

func (m *MotherVertex) AddEdges(i, j int) {
	_, ok := m.adjList[i]
	if !ok {
		m.adjList[i] = list.New()
	}
	m.adjList[i].PushBack(j)
}

func (m *MotherVertex) MotherV() int {
	var vrtx int
	visited := make([]bool, m.vertex)
	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}
	for i := 0; i < m.vertex; i++ {
		if !visited[i] {
			m.MotherVUtil(visited, i)
			vrtx = i
		}
	}

	for i := 0; i < m.vertex; i++ {
		visited[i] = false
	}

	m.MotherVUtil(visited, vrtx)

	for i := 0; i < m.vertex; i++ {
		if !visited[i] {
			return -1
		}
	}
	return vrtx
}

func (m *MotherVertex) MotherVUtil(visited []bool, u int) {
	visited[u] = true
	lst := m.adjList[u]
	for e := lst.Front(); e != nil; e = e.Next() {
		if !visited[e.Value.(int)] {
			m.MotherVUtil(visited, e.Value.(int))
		}
	}
}
