package main

import "fmt"

// ID interface opens up the possibilities of package user to
// implement own ID
//
// invariant: ensures keys are comparable (string).
type ID interface {
	String() string
}

// Vertex interface opens up the possibilities of package user to
// implement own Vertex that holds any contents.
//
// invariant: ID and ensure contents are printable.
type Vertex interface {
	ID() ID
	String() string
}

// stringID implements ID
type stringID string

func (s stringID) String() string {
	return string(s)
}

// vertex implements Vertex
type vertex struct {
	id string
}

func NewVertex(id string) Vertex {
	return &vertex{
		id: id,
	}
}

// ID returns vertex's id
func (v *vertex) ID() ID {
	return stringID(v.id)
}

// String returns vertex's content
func (v *vertex) String() string {
	return v.id
}

// Edge is defined as (u,v,w)
type Edge interface {
	Source() Vertex
	Target() Vertex
	Weight() int
}

type edge struct {
	src Vertex
	tgt Vertex
	wgt int
}

// NewEdge returns edge instance
func NewEdge(src, tgt Vertex, wgt ...int) Edge {

	if wgt == nil {
		wgt[0] = 0
	}

	return &edge{
		src: src,
		tgt: tgt,
		wgt: wgt[0],
	}
}

func (e *edge) Source() Vertex {
	return e.src
}
func (e *edge) Target() Vertex {
	return e.tgt
}
func (e *edge) Weight() int {
	return e.wgt
}

// Graph is defined as (V, E) mathematically.
type Graph interface {

	// AddVertex adds vertex to Graph, return false
	// if already exists.
	AddVertex(v Vertex) bool

	// AddEdge creates vertex u, v if not exist,
	// then form an edge via adjacency-list.
	//
	// return false if edge has been added.
	AddEdge(u, v Vertex, weight float64) bool
}

// Programming wise, its much more efficient to represent Graph
// with adjacency-list.
type graph struct {
	vtx map[ID]Vertex
	adj map[ID]map[ID]float64
}

// NewGraph returns graph instance.
func NewGraph() Graph {
	return &graph{
		vtx: make(map[ID]Vertex),
		adj: make(map[ID]map[ID]float64),
	}
}

func (g *graph) AddVertex(v Vertex) bool {
	if _, ok := g.vtx[v.ID()]; ok {
		return false
	}

	g.vtx[v.ID()] = v
	return true
}

func (g *graph) AddEdge(u, v Vertex, weight float64) bool {

	//if g.vtx[u]

	return false
}

func main() {
	var g Graph
	fmt.Println(g)
}
