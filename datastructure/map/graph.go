package main

import (
	"container/list"
	"fmt"
	"log"
)

type GraphInterface interface {
	CreateGraph(v uint)
	V() uint
	E() uint
	AddEdge(v uint, w uint)
	Adj(v uint) []uint
}

func (g *Graph) CreateGraph(v uint) *Graph {
	bag := g.b
	g.v = v
	for i := 0; uint(i) < v; i++ {
		b := NewBag()
		bag[i] = *b
	}
	return g
}

func (g *Graph) V() uint {
	return g.v
}

func (g *Graph) E() uint {
	return g.e
}

func (g *Graph) Bfs(s uint) {
	log.Println("Use Bfs start")
	g.visit = make([]bool, g.V())
	g.bfs(s)
}

func (g *Graph) bfs(s uint) {
	q := list.New()
	q.PushBack(s)
	for {
		//fmt.Print(q.Len())
		if q.Len() > 0 {
			first := q.Front()
			g.visit1(first.Value.(uint))
			q.Remove(first)

			adj := g.Adj(first.Value.(uint))
			cur := adj.first
			for cur != nil {
				if !g.visit[cur.val] {
					q.PushBack(cur.val)
				}

				cur = cur.next

			}
		} else {
			break
		}

	}
}

func (g *Graph) Dfs(s uint) {
	log.Println("Use Dfs Start")
	g.visit = make([]bool, g.V())
	g.dfs(s)
}

func (g *Graph) dfs(s uint) {
	g.visit1(s)
	adj := g.Adj(s)
	cur := adj.root
	for cur != nil {
		if !g.visit[cur.val] {
			g.dfs(cur.val)
			cur = cur.next
		} else {
			cur = cur.next
		}
	}
}
func (g *Graph) visit1(s uint) {
	if !g.visit[s] {
		g.visit[s] = true
		fmt.Print("-->", s, "-->")
	}

}
func (g *Graph) AddEdge(v uint, w uint) {
	n := &Node{val: w}
	g.b[v].Add(n)
	n2 := &Node{val: v}
	g.b[w].Add(n2)
	g.e++
}

func (g *Graph) Adj(v uint) *Bag {
	return &g.b[v]
}

type Graph struct {
	v     uint
	e     uint
	b     []Bag
	visit []bool
}

func NewGraph(n uint) *Graph {

	return &Graph{
		v: n,
		b: make([]Bag, n),
	}
}
