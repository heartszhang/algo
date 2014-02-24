package algo

type graph interface {
	vertex_count() int
	edges(v int) []edge
}

type edge interface {
	v0() int
	v1() int
	weight() uint
}

type astar_graph interface {
	graph
	heuristic_estimate(start, goal int) uint
}
