package core

// Graph structures
type Graph struct {
	vertices      int
	adjacencyList map[int][]int
	reversed      map[int][]int
}

// NewGraph constructs a new Graph instance
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices:      vertices,
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
	stack[v] = false
	return false
}

// IsCyclic returns true if the Graph has at least one cycle
func (g Graph) IsCyclic() bool {
	visited := make([]bool, g.vertices)
	stack := make([]bool, g.vertices)

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

// AddEdge creates an edge between two vertices
func (g Graph) AddEdge(v int, w int) {
	if vs, ok := g.adjacencyList[v]; ok {
		g.adjacencyList[v] = append(vs, w)
	} else {
		g.adjacencyList[v] = []int{w}
	}
}
