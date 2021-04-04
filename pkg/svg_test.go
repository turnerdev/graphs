package core

import (
	"encoding/xml"
	"strings"
	"testing"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func TestSVG(t *testing.T) {
	t.Run("Test SVG document", func(t *testing.T) {
		svg := New()
		svg.Rect(0, 0, 100, 100).
			Rect(0, 0, 100, 100).
			Rect(0, 0, 100, 100)

		var b strings.Builder

		_, err := svg.Write(&b)
		if err != nil {
			t.Error(err)
		}

		var node Node
		xml.Unmarshal([]byte(b.String()), &node)

		if node.XMLName.Local != "svg" {
			t.Errorf("Expected svg, got %q", node.XMLName.Local)
		}
		if len(node.Nodes) != 3 {
			t.Errorf("Expected 3 nodes, got %d", len(node.Nodes))
		}
	})

	t.Run("Test SVG enter/exit", func(t *testing.T) {
		svg := New()
		svg.Rect(0, 0, 100, 100).
			Enter().
			Enter(). // Redundant Enters should be ignored
			Rect(0, 0, 100, 100).
			Rect(0, 0, 100, 100).
			Rect(0, 0, 100, 100).
			Exit().
			Exit(). // Redundant Exits should be ignored
			Rect(0, 0, 100, 100)

		var b strings.Builder

		_, err := svg.Write(&b)
		if err != nil {
			t.Error(err)
		}

		var node Node
		xml.Unmarshal([]byte(b.String()), &node)

		if len(node.Nodes) != 2 || len(node.Nodes[0].Nodes) != 3 {
			t.Errorf("Expected 2-3, got %d-%d", len(node.Nodes), len(node.Nodes[0].Nodes))
		}
	})
}
