# go-graph
Graph Theories in Go.

### Graph Search

It's about "exploring" a graph.

Given node s, t in a graph, and you'd like to find a path. For example,

> A particular state of a Rubik's Cube is given, and I want to know is there
  some path that gets me into a solve state?

Commonly, BFS is used (one type of exploration technique).

### Application
- web crawling.
- social networking.
- network broadcast.
- garbage collection.
- model checking.
- check mathematical conj.
- solving puzzles of game.

## Graph Representation

In general, G = {V, E}.

V = Set of vertices.
E = Set of edges.

 - e = {v, w} unordered pairs. *undirected*
 - e = (v, w) ordered pairs. *directed*

### 1. Naive Representation.

V = An array of nodes(vertices).
E = An array of (k,v) pairs (edges).

#### Why it's BAD?

To solve a graph problem, you need to explore. And the basic exploration technique
is to start to look-up from neighbours (adjacent vertices).

With *naive* representation, finding all adjecent vetices from a particular vertext u,
takes linear time on E array.

We need a better representation.

### 2. Adjacency Representation.

Adjacency lists: (it could be array, or hash)
- array adj of size |v| linked list.
- for each vertex `u ∈ V`, Adj[u] stores `u`'s neighbors, namely the vertices you can reach by one step from u.
- formally, Adj[u] = {v ∈ V | (u,v) ∈ E }, Adj[u] is a Set of vertices, v such that (u, v) is subset of E.
  - Adj[b] = { a, c }

At least for graph exploration proble, this is the representation you want. Because you're at some vertex, and you
want to know can I go next?


What about it's Space complexity?
`O( V + E )` for undirected graph,
`O( V + 2E )` for directed graph, but its essentially just `O( V + E )`.

Ideally, your algorithm will run in that much of time, as those are the noes you need to lookup. `O (V + E )`

### 3. Object-Oriented(Adjacency Representation)

> V.neighbors = Adj[v]


OO vs Non-OO:
- One graph vs Multiple graphs. As V object is essentially binded to ONE graph.



### 4. Implicit Representation

> Adj[u] is a function. v.neighbors() is a method.

It uses less space, 0 space.

The idea: start with some vertex, given a vertext, know how to compute the neighbors of the vertext, O(n) or O(1).


Don't have to build the whole graph, just build enough of it until you find your answer.

#### Scenarios?
> Rubik's cube. You never want to build the space, it has a **bajillion** states (vertices).

For n * n * n cube, in O(n) time, you can list all the O(n) next states, you can list O(n) neighbors, so you can keep
**exploring**, searching for your state.




---

### Breadth first search
- visit all nodes(vertices) reachable from given s ∈ V
- O(V+E) time.
- look at nodes reachable in O moves, 1 move, 2 moves, ..., until run out of graph.
- careful to avoid duplicates. (it would spend infinite times...)


#### Pseudo-code
```python
BFS(s, Adj):
	level = {s: 0}      // hash table.
	parent = {s: None}  // hash table.
	i = 1
	frontier = [s]      // append to array: frontier
	while frontier:
		next = []
		for u in frontier:
			for v in Adj[u]:
				if v not in level:
					level[v] = i
					parent[v] = u
					next.append(v)
		frontier = next
		i +=1
```
1. **next**. How do we compute `next`?

  We look at every node in the frontier, and we look all the nodes you can reach from those nodes, u.

2. How do we check for duplicates (visited nodes)?

  Look at the edges, Adj[u], then key thing is to check for duplicates, `for v not in level`. (have we seen the nodes before?)

3. The gist of the algorithm?

  To visit all the nodes reachable from s (initial node), in linear time.

4. As the side effects of BFS, the following is what you get for free:

  - # of moves with *level[u]*. From initial node S to to u ∈ V.
  - shortest path with *Π parent(u)*. From initial node S to u ∈ V.
