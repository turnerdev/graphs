package core

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

// Element of an SVG document
type Element struct {
	tag        string
	attributes map[string]string
	children   []*Element
	parent     *Element
}

// SVG document
type SVG struct {
	Element
}

// New constructs an empty SVG document
func New() *SVG {
	return &SVG{
		Element{
			tag: "svg",
			attributes: map[string]string{
				"xmlns":       "http://www.w3.org/2000/svg",
				"xmlns:xlink": "http://www.w3.org/1999/xlink",
				"width":       "800",
				"height":      "800",
				"viewBox":     "0 0 800 800",
			},
			children: make([]*Element, 0),
		},
	}
}

// Enter moves the builder context to the last child of the current element
func (e *Element) Enter() *Element {
	if len(e.children) > 0 {
		return e.children[len(e.children)-1]
	}
	return e
}

// Exit moves the builder context to the parent of the current element
func (e *Element) Exit() *Element {
	if e.parent != nil {
		return e.parent
	}
	return e
}

// Rect creates an SVG rect element
func (e *Element) Rect(x int, y int, width int, height int) *Element {
	element := &Element{
		tag: "rect",
		attributes: map[string]string{
			"x":      fmt.Sprint(x),
			"y":      fmt.Sprint(y),
			"width":  fmt.Sprint(width),
			"height": fmt.Sprint(height),
		},
		parent:   e,
		children: make([]*Element, 0),
	}
	e.children = append(e.children, element)
	return e
}

// Render SVG document content to writer
func (s *SVG) Write(w io.Writer) (int, error) {
	n, err := w.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
	if err != nil {
		return n, err
	}
	nn, err := write(&s.Element, w, 0)
	return n + nn, err
}

func (e *Element) attributeString() string {
	var b bytes.Buffer

	for key, value := range e.attributes {
		b.WriteString(" ")
		b.WriteString(key)
		b.WriteString("=\"")
		xml.EscapeText(&b, []byte(value))
		b.WriteString("\"")
	}

	return b.String()
}

// write an svg element to a given writer. Depth is used to track the indentation
func write(e *Element, w io.Writer, depth int) (int, error) {
	isLeaf := len(e.children) == 0
	var closer string
	if isLeaf {
		closer = "/"
	}
	n, err := w.Write([]byte(fmt.Sprintf("%s<%s%s%s>\n", strings.Repeat("\t", depth), e.tag, e.attributeString(), closer)))
	if err != nil {
		return n, err
	}
	for _, child := range e.children {
		nn, err := write(child, w, depth+1)
		n += nn
		if err != nil {
			return n, err
		}
	}
	if isLeaf {
		return n, nil
	}
	nn, err := w.Write([]byte(fmt.Sprintf("%s</%s>\n", strings.Repeat("\t", depth), e.tag)))
	return n + nn, err
}
