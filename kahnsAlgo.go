package app

import (
	"container/list"
	"fmt"
)

type KahnsAlgo struct {
	vertex  int
	adjList map[int]*list.List
}

func (k *KahnsAlgo) KahnsAlgoInit(vertex int) {
	k.vertex = vertex
	k.adjList = make(map[int]*list.List)
	for i := 0; i < k.vertex; i++ {
		k.adjList[i] = list.New()
	}
}

func (k *KahnsAlgo) AddEdges(i, j int) {
	_, ok := k.adjList[i]
	if !ok {
		k.adjList[i] = list.New()
	}
	k.adjList[i].PushBack(j)
}

func (k *KahnsAlgo) KahnsAlgorithm() {
	indegree := make([]int, k.vertex)
	visited := make([]bool, k.vertex)

	for i := 0; i < k.vertex; i++ {
		visited[i] = false
	}

	for i := 0; i < k.vertex; i++ {
		indegree[i] = 0
	}

	for i := range k.vertex {
		lst := k.adjList[i]
		for e := lst.Front(); e != nil; e = e.Next() {
			indegree[e.Value.(int)] += 1
		}
	}

	queue := list.New()
	for i := 0; i < k.vertex; i++ {
		if indegree[i] == 0 {
			queue.PushBack(i)
		}
	}
	count := 0
	for queue.Len() != 0 {
		elm := queue.Front()
		val := elm.Value.(int)
		fmt.Println(val)
		queue.Remove(elm)
		lst := k.adjList[val]
		for e := lst.Front(); e != nil; e = e.Next() {
			indegree[e.Value.(int)] = indegree[e.Value.(int)] - 1
			if indegree[e.Value.(int)] == 0 {
				queue.PushBack(e.Value.(int))
			}
		}
		count += 1
	}

	if count == k.vertex {
		fmt.Println("Topological sorting is possible.")
	} else {
		fmt.Println("Topological sorting not possible.")
	}
}
