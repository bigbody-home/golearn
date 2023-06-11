package main

import "fmt"

func main() {
	g := NewGraph(6).CreateGraph(6)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	//g.Adj(2).PrintNode()
	g.Dfs(0)
	fmt.Println()
	g.Bfs(0)

}
