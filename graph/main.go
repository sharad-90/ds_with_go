package main

import (
	"fmt"
	"math"
)

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key      int
	adjacent []*Vertex
}

func (g *Graph) AddVertex(k int) {
	if g.contain(g.vertices, k) {
		fmt.Printf("Vertex %d is not added because it is an existing key", k)
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}

}

func (g *Graph) AddEdge(source, sink int) {
	srcEdge := g.GetVertex(source)
	sinkEdge := g.GetVertex(sink)
	if srcEdge == nil || sinkEdge == nil {
		err := fmt.Errorf("Invalid edge  (%v->%v)", source, sink)
		fmt.Println(err.Error())
	} else if g.contain(srcEdge.adjacent, sink) {
		err := fmt.Errorf("Existing edge  (%v->%v)", source, sink)
		fmt.Println(err.Error())
	} else {
		srcEdge.adjacent = append(srcEdge.adjacent, sinkEdge)
	}
}

func (g *Graph) printGraph() {
	for _, v := range g.vertices {
		fmt.Printf("\n Vertex %d :", v.key)
		for _, e := range v.adjacent {
			fmt.Printf("%v", e.key)
		}
	}
}

func (g *Graph) GetVertex(key int) *Vertex {
	for i, v := range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) contain(s []*Vertex, key int) bool {
	for _, v := range s {
		if v.key == key {
			return true
		}
	}
	return false
}
func main() {
	g := &Graph{}
	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 0)
	g.AddEdge(1, 3)
	g.AddEdge(2, 0)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 1)
	g.AddEdge(3, 2)
	g.AddEdge(3, 4)
	g.AddEdge(4, 2)
	g.AddEdge(4, 3)
	g.printGraph()
	fmt.Printf(" \n size of the graph is %d\n", len(g.vertices))
	bfs(g, len(g.vertices), 0)
	fmt.Println("\n")
	dfs(g, len(g.vertices), 0)
	fmt.Println("\n")
	shortestdistance(g, len(g.vertices), 0)
	cycleDetection()
}

func bfs(g *Graph, v, source int) {
	visited := make([]bool, v)
	visited[source] = true
	q := newQueue()
	q.enqueue(source)
	for !q.isEmpty() {
		u := q.front.data
		fmt.Printf("%d->", u)
		q.dequeue()
		for _, vertex := range g.GetVertex(u).adjacent {
			if visited[vertex.key] == false {
				q.enqueue(vertex.key)
				visited[vertex.key] = true

			}
		}
	}
}

func bfsDisconnected(g *Graph, v int) {
	visited := make([]bool, v)
	for i := 0; i < v; i++ {
		if visited[i] == false {
			//bfs(g, v, visited)
		}
	}
}

func dfs(g *Graph, v, s int) {
	visited := make([]bool, v)
	for i := 0; i < v; i++ {
		if !visited[i] {
			dfsRecursive(g, i, visited)
		}
	}
}

func dfsRecursive(g *Graph, s int, visited []bool) {
	visited[s] = true
	fmt.Printf("%d-->", s)
	for _, v := range g.GetVertex(s).adjacent {
		if !visited[v.key] {
			dfsRecursive(g, v.key, visited)
		}
	}
}

func shortestdistance(g *Graph, v, source int) {
	visited := make([]bool, v)
	distance := make([]int, v)
	for i, _ := range distance {
		distance[i] = math.MaxInt16
	}
	visited[source] = true
	distance[source] = 0
	q := newQueue()
	q.enqueue(source)
	for !q.isEmpty() {
		u := q.front.data
		q.dequeue()
		for _, vertex := range g.GetVertex(u).adjacent {
			if visited[vertex.key] == false {
				q.enqueue(vertex.key)
				distance[vertex.key] = distance[u] + 1
				visited[vertex.key] = true

			}
		}
	}
	fmt.Println(distance)
}

func cycleDetection() {
	g := &Graph{}
	for i := 0; i < 6; i++ {
		g.AddVertex(i)
	}
	g.AddEdge(0, 1)
	g.AddEdge(1, 0)
	g.AddEdge(1, 2)
	g.AddEdge(2, 1)
	g.AddEdge(1, 3)
	g.AddEdge(3, 1)
	g.AddEdge(2, 3)
	g.AddEdge(3, 2)
	g.AddEdge(2, 4)
	g.AddEdge(4, 2)
	g.AddEdge(4, 5)
	g.AddEdge(5, 4)
	fmt.Println(dfsCycle(g, 6))
}

func dfsCycleRec(g *Graph, visited []bool, source, parent int) bool {
	visited[source] = true
	for _, vertex := range g.GetVertex(source).adjacent {
		if !visited[vertex.key] {
			if dfsCycleRec(g, visited, vertex.key, source) {
				return true
			}
		} else if vertex.key != parent {
			return true
		}
	}
	return false
}

func dfsCycle(g *Graph, vertex int) bool {
	visited := make([]bool, vertex)
	for i := 0; i < len(visited); i++ {
		if !visited[i] {
			if dfsCycleRec(g, visited, 0, -1) {
				return true
			}
		}
	}
	return false
}
