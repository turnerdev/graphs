package core

import "testing"

func TestGraphDetectCycles(t *testing.T) {

	t.Run("1 cycle", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 1)
		g.AddEdge(3, 4)

		if !g.IsCyclic() {
			t.Errorf("Cyclic graph not detected")
		}
	})

	t.Run("No cycle", func(t *testing.T) {
		g := NewGraph()
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		if g.IsCyclic() {
			t.Errorf("Cyclic graph not detected")
		}
	})

}

func TestDegree(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 3)
	g.AddEdge(2, 5)
	g.AddEdge(3, 5)

	t.Run("Test vertex degree", func(t *testing.T) {
		want := 5
		got := g.GetVertex(2).Degree()
		if want != got {
			t.Errorf("degree is %d ; want %d", got, want)
		}
	})

	t.Run("Test vertex indegree", func(t *testing.T) {
		want := 2
		got := g.GetVertex(2).Indegree()
		if want != got {
			t.Errorf("degree is %d ; want %d", got, want)
		}
	})

	t.Run("Test vertex outdegree", func(t *testing.T) {
		want := 3
		got := g.GetVertex(2).Outdegree()
		if want != got {
			t.Errorf("degree is %d ; want %d", got, want)
		}
	})

}
