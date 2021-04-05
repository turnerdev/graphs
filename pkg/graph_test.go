package core

import "testing"

func TestGraphDetectCycles(t *testing.T) {

	t.Run("1 cycle", func(t *testing.T) {
		g := NewGraph(5)
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
		g := NewGraph(5)
		g.AddEdge(0, 1)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		g.AddEdge(3, 4)

		if g.IsCyclic() {
			t.Errorf("Cyclic graph not detected")
		}
	})

}
