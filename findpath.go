package app

import (
	"container/list"
	"fmt"
)

type FindPath struct {
	vertex  int
	adjList map[int]*list.List
	allPath [][]int
}

func (f *FindPath) FindPathInit(vertex int) {
	f.vertex = vertex
	f.allPath = make([][]int, 0)
	f.adjList = make(map[int]*list.List)
	for i := 0; i < f.vertex; i++ {
		f.adjList[i] = list.New()
	}
}

func (f *FindPath) AddEdges(i, j int) {
	_, ok := f.adjList[i]
	if !ok {
		f.adjList[i] = list.New()
	}
	f.adjList[i].PushBack(j)
}

func (f *FindPath) DisplayGraph() {
	for i := 0; i < f.vertex; i++ {
		lst := f.adjList[i]
		fmt.Printf("%d-->", i)
		for e := lst.Front(); e != nil; e = e.Next() {
			fmt.Printf("%d-", e.Value.(int))
		}
		fmt.Println("")
	}
}

func (f *FindPath) FPath(u, v int) {
	temp := make([]int, 0)
	visited := make([]bool, f.vertex)
	for i := 0; i < f.vertex; i++ {
		visited[i] = false
	}
	f.FPath_util(u, v, visited, temp)
	fmt.Println("All path from source to destination is:")
	fmt.Println(f.allPath)
}

func (f *FindPath) FPath_util(i, target int, visited []bool, temp []int) {
	visited[i] = true
	temp = append(temp, i)
	fmt.Println(temp)
	if i == target {
		newlist := make([]int, len(temp))
		copy(newlist, temp)
		f.allPath = append(f.allPath, newlist)
	} else {
		lst := f.adjList[i]
		for e := lst.Front(); e != nil; e = e.Next() {
			if !visited[e.Value.(int)] {
				f.FPath_util(e.Value.(int), target, visited, temp)
			}
		}
	}
	visited[i] = false
	temp = temp[:len(temp)-1]
	fmt.Println(temp)
}
