package core

import (
	"container/list"
	"errors"
	"fmt"
	"math"
)

// Edge structure
type Edge struct {
	v        *Vertex
	w        *Vertex
	capacity int
	flow     int
	weight   float32
}

// SetWeight sets the weight of a given edge
func (e *Edge) SetWeight(weight float32) {
	e.weight = weight
}

// SetCapacity sets the capacity of a flow network arc
func (e *Edge) SetCapacity(flow int) *Edge {
	e.capacity = flow
	return e
}

// SetFlow sets the flow of a flow network arc
func (e *Edge) SetFlow(flow int) *Edge {
	e.flow = flow
	return e
}

// Vertex structure
type Vertex struct {
	id  int
	in  map[int]*Edge
	out map[int]*Edge
}

func (v *Vertex) getEdge(id int) (*Edge, error) {
	if edge, ok := v.out[id]; ok {
		return edge, nil
	}
	return nil, errors.New("No edge found")
}

func (v *Vertex) addEdge(w *Vertex) *Edge {
	if _, ok := v.out[w.id]; !ok {
		edge := &Edge{
			v: v,
			w: w,
		}
		v.out[w.id] = edge
		w.in[v.id] = edge
	}
	return v.out[w.id]
}

// Degree (or valency) is the number of edges that are incident to the Vertex
func (v Vertex) Degree() int {
	return v.Indegree() + v.Outdegree()
}

// Indegree is the number of head ends adjacent to the Vertex
func (v Vertex) Indegree() int {
	return len(v.in)
}

// Outdegree is the number of tail ends adjacent to the Vertex
func (v Vertex) Outdegree() int {
	return len(v.out)
}

func (v Vertex) IsSource() bool {
	return v.Indegree() == 0
}

func (v Vertex) IsSink() bool {
	return v.Outdegree() == 0
}

// Graph structures
type Graph struct {
	vertices      map[int]*Vertex
	adjacencyList map[int][]int
	reversed      map[int][]int
	sinks         map[int]bool
	sources       map[int]bool
}

// NewGraph constructs a new Graph instance
func NewGraph() *Graph {
	return &Graph{
		vertices:      make(map[int]*Vertex),
		adjacencyList: make(map[int][]int),
		reversed:      make(map[int][]int),
		sinks:         make(map[int]bool),
		sources:       make(map[int]bool),
	}
}

// isCycle performs a recursive depth first search to identify cycles in the graph
func (g Graph) isCyclic(v int, visited []bool, stack []bool) bool {
	visited[v] = true
	stack[v] = true

	// If any neighbour is visited and in the stack, graph is cyclic
	for _, w := range g.adjacencyList[v] {
		if !visited[w] {
			if g.isCyclic(w, visited, stack) {
				return true
			}
		} else if stack[w] {
			return true
		}
	}
	// Pop the stack
	stack[v] = false
	return false
}

// IsCyclic returns true if the Graph has at least one cycle
func (g Graph) IsCyclic() bool {
	visited := make([]bool, len(g.vertices))
	stack := make([]bool, len(g.vertices))

	// With each vertex as a starting point, walk the graph
	for v := range g.adjacencyList {
		if !visited[v] {
			if g.isCyclic(v, visited, stack) {
				return true
			}
		}
	}
	return false
}

// AddVertex creates a Vertex with the given id
func (g Graph) AddVertex(v int) *Vertex {
	vertex := &Vertex{
		id:  v,
		in:  make(map[int]*Edge),
		out: make(map[int]*Edge),
	}
	g.vertices[v] = vertex
	g.adjacencyList[v] = make([]int, 0)
	return vertex
}

// GetVertex returns the Vertex with the given id
func (g Graph) GetVertex(v int) *Vertex {
	return g.vertices[v]
}

// AddEdge creates an edge between two vertices
func (g Graph) AddEdge(v int, w int) *Edge {
	vertexA, ok := g.vertices[v]
	if !ok {
		vertexA = g.AddVertex(v)
		g.sources[v] = true
	} else if vertexA.IsSink() {
		delete(g.sinks, v)
	}
	vertexB, ok := g.vertices[w]
	if !ok {
		vertexB = g.AddVertex(w)
		g.sinks[w] = true
	} else if vertexB.IsSource() {
		delete(g.sources, w)
	}
	edge := vertexA.addEdge(vertexB)

	g.adjacencyList[v] = append(g.adjacencyList[v], w)

	return edge
}

// Sources returns a list of vertices with an indegree of 0
func (g Graph) Sources() []*Vertex {
	keys := make([]*Vertex, len(g.sources))
	i := 0
	for k := range g.sources {
		keys[i] = g.vertices[k]
		i++
	}
	return keys
}

// Sinks returns a list of vertices with an outdegree of 0
func (g Graph) Sinks() []*Vertex {
	keys := make([]*Vertex, len(g.sources)+1)
	i := 0
	for k := range g.sinks {
		keys[i] = g.vertices[k]
		i++
	}
	return keys
}

func (g *Graph) Search(strategy SearchStrategy) {
	result := strategy.Search(g.Sources()[0], func(v *Vertex) bool {
		return false
	})

	fmt.Println(result)
}

type SearchResult struct {
	paths []*Vertex
}

type SearchStrategy interface {
	Search(g *Vertex, fn func(*Vertex) bool) SearchResult
}

type BFSStrategy struct{}

func (s BFSStrategy) Search(v *Vertex, fn func(v *Vertex) bool) SearchResult {
	// visited := make([]int)
	// queue = []*Vertex{v}

	// for len(queue) > 0 {
	// u := queue[]
	// }

	return SearchResult{
		// paths:
	}
}

func ResidualGraph(g *Graph) *Graph {
	r := NewGraph()
	for _, v := range g.vertices {
		for _, e := range v.out {
			// Set residual capacity on forward edge by subtracting flow from capacity
			r.AddEdge(v.id, e.w.id).SetCapacity(e.capacity - e.flow)
			if e.flow > 0 {
				// Set back edge
				r.AddEdge(e.w.id, v.id).SetCapacity(e.flow)
			}
		}
	}
	return r
}

// FordFulkerson algorithm returns the max flow of a given graph
func FordFulkerson(g *Graph) int {
	r := ResidualGraph(g)

	// TODO: Use synthetic source/sink where multiple exist
	s := r.GetVertex(g.Sources()[0].id)
	t := r.GetVertex(g.Sinks()[0].id)

	bfs := func(r *Graph, s *Vertex, t *Vertex) ([]int, bool) {
		visited := make(map[int]*Vertex)
		queue := list.New()
		queue.PushBack(s)
		visited[s.id] = s
		parent := make([]int, len(r.vertices))

		for queue.Len() > 0 {
			el := queue.Front()
			queue.Remove(el)
			v := el.Value.(*Vertex)

			for id, edge := range v.out {
				if _, ok := visited[id]; !ok && edge.capacity-edge.flow > 0 {
					parent[edge.w.id] = v.id
					if edge.w == t {
						return parent, false
					}
					queue.PushBack(edge.w)
					visited[edge.w.id] = edge.w
				}
			}
		}

		return nil, true
	}

	bfsIterator := func(r *Graph, v *Vertex, t *Vertex, f func([]int)) {
		for {
			results, done := bfs(r, v, t)
			if done {
				return
			}
			f(results)
		}
	}

	maxFlow := 0

	// Iterate over augmenting paths using breadth first search
	bfsIterator(r, s, t, func(path []int) {
		pathFlow := math.MaxInt32

		for w := t.id; w != s.id; w = path[w] {
			edge, _ := r.GetVertex(path[w]).getEdge(w)
			if pathFlow > edge.capacity-edge.flow {
				pathFlow = edge.capacity - edge.flow
			}
		}
		for w := t.id; w != s.id; w = path[w] {
			edge, _ := r.GetVertex(path[w]).getEdge(w)
			edge.flow += pathFlow
			edge.w.addEdge(edge.v).flow += pathFlow
		}
		maxFlow += pathFlow
	})

	return maxFlow
}

type DFSStrategy struct{}

// type queue struct
