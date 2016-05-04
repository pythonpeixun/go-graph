package main

type node struct {
	val  int
	node *node
}

type graph struct {
	root *node
}
