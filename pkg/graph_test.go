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
	g.AddEdge(3, 4)
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
			t.Errorf("degree is %d; want %d", got, want)
		}
	})

	t.Run("Test sinks", func(t *testing.T) {
		got := g.Sinks()
		if len(got) != 2 || got[0].id != 4 || got[1].id != 5 {
			t.Errorf("found %d sinks; want 2 (%d, %d)", len(got), got[0].id, got[1].id)
		}
	})

	t.Run("Test sources", func(t *testing.T) {
		got := g.Sources()
		if len(got) != 1 || got[0].id != 0 {
			t.Errorf("found %d sources; want 1 (%d)", len(got), got[0].id)
		}
	})

}
