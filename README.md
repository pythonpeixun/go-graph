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


