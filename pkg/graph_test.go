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

func TestResidualGraph(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1).SetFlow(1).SetCapacity(3)
	g.AddEdge(0, 2).SetFlow(2).SetCapacity(2)
	g.AddEdge(2, 1).SetFlow(1).SetCapacity(3)
	g.AddEdge(1, 3).SetFlow(2).SetCapacity(2)
	g.AddEdge(3, 2).SetFlow(1).SetCapacity(1)
	g.AddEdge(3, 4).SetFlow(1).SetCapacity(3)
	g.AddEdge(3, 5).SetFlow(2).SetCapacity(3)
	g.AddEdge(4, 5).SetFlow(1).SetCapacity(2)

	r := ResidualGraph(g)

	if r.GetVertex(0).out[1].capacity != 2 {
		t.Errorf("Expected 2")
	}

	// TODO: implement getEdge(v, w) for more efficient testing
}

func TestFordFulkerson(t *testing.T) {
	g := NewGraph()
	g.AddEdge(0, 1).SetCapacity(11)
	g.AddEdge(0, 2).SetCapacity(12)
	g.AddEdge(2, 1).SetCapacity(1)
	g.AddEdge(1, 3).SetCapacity(12)
	g.AddEdge(2, 4).SetCapacity(11)
	g.AddEdge(4, 3).SetCapacity(7)
	g.AddEdge(3, 5).SetCapacity(19)
	g.AddEdge(4, 5).SetCapacity(4)

	want := 23
	got := FordFulkerson(g)

	if got != want {
		t.Errorf("%d; expected %d", got, want)
	}
}
