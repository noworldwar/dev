package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj_list map[*Vertax](*list.List)
}

type Vertax struct {
	Val interface{}
}

func NewGraph() *Graph {
	g := new(Graph)
	g.adj_list = make(map[*Vertax](*list.List))
	return g
}

func (g *Graph) NewVertax(number int) *Vertax {
	v := new(Vertax)
	v.Val = number
	g.adj_list[v] = list.New()
	return v
}

func (g *Graph) AddEdge(start, end *Vertax) {
	g.adj_list[start].PushBack(end)
	g.adj_list[end].PushBack(start)
}

func (g *Graph) EasyTraversal() {
	for k, v := range g.adj_list {
		fmt.Printf("[%v] -> ", k.Val)
		for current := v.Front(); current != nil; current = current.Next() {
			fmt.Printf("%v -> ", current.Value.(*Vertax).Val)
		}
		fmt.Println("nil")
	}
}

func (g *Graph) BFS(start *Vertax) {
	visited := make(map[*Vertax]bool)
	path_queue := queue.New()
}

func main() {
	g := NewGraph()
	v0 := g.NewVertax(0)
	v1 := g.NewVertax(1)
	v2 := g.NewVertax(2)
	v3 := g.NewVertax(3)
	v4 := g.NewVertax(4)
	v5 := g.NewVertax(5)
	g.AddEdge(v0, v1)
	g.AddEdge(v0, v2)
	g.AddEdge(v0, v3)
	g.AddEdge(v1, v5)
	g.AddEdge(v3, v5)
	g.AddEdge(v4, v5)
	g.EasyTraversal()

}
