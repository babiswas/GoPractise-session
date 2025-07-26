package app

import (
	"container/list"
	"fmt"
)

type LvlTrav struct {
	vertex  int
	adjList map[int]*list.List
}

func (l *LvlTrav) LvlTravInit(vertex int) {
	l.vertex = vertex
	l.adjList = make(map[int]*list.List)
	for i := 0; i < l.vertex; i++ {
		l.adjList[i] = list.New()
	}
}

func (l *LvlTrav) AddEdges(i, j int) {
	_, ok := l.adjList[i]
	if !ok {
		l.adjList[i] = list.New()
	}
	l.adjList[i].PushBack(j)
	_, ok = l.adjList[j]
	if !ok {
		l.adjList[j] = list.New()
	}
	l.adjList[j].PushBack(i)
}

func (l *LvlTrav) FindPath(source, target int) {
	prev := l.ShortestPath(source)
	fmt.Println(prev)
	path := make([]int, 0)
	path = append(path, target)
	temp := prev[target]
	for temp != -1 {
		path = append(path, temp)
		temp = prev[temp]
	}
	start_index := 0
	end_index := len(path) - 1
	for start_index < end_index {
		path[start_index], path[end_index] = path[end_index], path[start_index]
		start_index += 1
		end_index -= 1
	}
	fmt.Println(path)
	if (path[0] == source) && (path[len(path)-1] == target) {
		fmt.Println("Shortest path exists:")
		fmt.Println(path)
	} else {
		fmt.Println("Path do not exist:")
	}

}

func (l *LvlTrav) ShortestPath(source int) []int {
	queue := list.New()
	visited := make([]bool, l.vertex)
	for i := 0; i < l.vertex; i++ {
		visited[i] = false
	}
	prev := make([]int, l.vertex)
	for i := 0; i < l.vertex; i++ {
		prev[i] = -1
	}
	queue.PushBack(source)
	visited[source] = true
	for queue.Len() != 0 {
		elm := queue.Front()
		val := elm.Value.(int)
		queue.Remove(elm)
		lst := l.adjList[val]
		for e := lst.Front(); e != nil; e = e.Next() {
			if !visited[e.Value.(int)] {
				queue.PushBack(e.Value.(int))
				visited[e.Value.(int)] = true
				prev[e.Value.(int)] = val
			}
		}
	}
	return prev
}
