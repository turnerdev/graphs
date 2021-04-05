package core

// Edge structure
type Edge struct {
	v        *Vertex
	w        *Vertex
	capacity int
	flow     int
	weight   int
}

// Vertex structure
type Vertex struct {
	id  int
	in  []*Edge
	out []*Edge
}

func (v *Vertex) addEdge(w *Vertex) *Edge {
	edge := &Edge{
		v: v,
		w: w,
	}
	v.out = append(v.out, edge)
	w.in = append(w.in, edge)
	return edge
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

// Graph structures
type Graph struct {
	vertices      map[int]*Vertex
	adjacencyList map[int][]int
	reversed      map[int][]int
}

// NewGraph constructs a new Graph instance
func NewGraph() *Graph {
	return &Graph{
		vertices:      make(map[int]*Vertex),
		adjacencyList: make(map[int][]int),
		reversed:      make(map[int][]int),
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
		in:  make([]*Edge, 0),
		out: make([]*Edge, 0),
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
	}
	vertexB, ok := g.vertices[w]
	if !ok {
		vertexB = g.AddVertex(w)
	}
	edge := vertexA.addEdge(vertexB)

	g.adjacencyList[v] = append(g.adjacencyList[v], w)

	return edge
}
