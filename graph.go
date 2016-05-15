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
	// return false if edge has previously been added.
	AddEdge(u, v Vertex, weight float64) bool
}

// Programming wise, its much more efficient to represent Graph
// with adjacency-list.
type graph struct {
	vertices map[ID]Vertex
	adj      map[ID]map[ID]float64
}

// NewGraph returns graph instance.
func NewGraph() Graph {
	return &graph{
		vertices: make(map[ID]Vertex),
		adj:      make(map[ID]map[ID]float64),
	}
}

func (g *graph) AddVertex(v Vertex) bool {
	if _, ok := g.vertices[v.ID()]; ok {
		return false
	}

	g.vertices[v.ID()] = v
	return true
}

func (g *graph) AddEdge(u, v Vertex, weight float64) bool {

	g.AddVertex(u)
	g.AddVertex(v)

	edges := g.adj[u]

	if edges == nil {
		edges = make(map[ID]float64)
		g.adj[u] = edges
	}

	if _, found := edges[v]; found {
		return false
	}

	edges[v] = weight
	return true
}

func main() {
	var g Graph
	fmt.Println(g)
}
