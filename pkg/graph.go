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
