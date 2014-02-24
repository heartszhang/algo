package algo

import "container/heap"

// path[i] : previous index of i, -2 means unreachable
func init_path(vcount, begin int) []int {
	path := make([]int, vcount)
	for i := 0; i < vcount; i++ {
		path[i] = -2
	}
	path[begin] = -1
	return path
}

type weight_priority uint

const (
	max_weight = ^weight_priority(0)
)

func (lhs weight_priority) less(rhs priority) bool {
	return lhs < rhs.(weight_priority)
}

func init_distance(vcount, begin int) []*prioq_item {
	dist := make([]*prioq_item, vcount)
	for i := 0; i < len(dist); i++ {
		dist[i] = &prioq_item{max_weight, i, i}
	}
	dist[begin].val = weight_priority(0)
	return dist
}

func dijkstra(g graph, source int) []int {
	path := init_path(g.vertex_count(), source)
	dist := init_distance(g.vertex_count(), source)
	distq := prioq(dist[:])
	heap.Init(&distq)

	for distq.Len() > 0 {
		m := heap.Pop(&distq).(*prioq_item)
		if m.val.(weight_priority) == max_weight {
			break
		}
		for _, edge := range g.edges(m.origin_index) {
			if dist[edge.v1()].heap_index < 0 {
				continue
			}
			relax(m.origin_index, edge.v1(), edge.weight(), dist, path, &distq)
		}

	}

	return path
}

func relax(v, u int, weight uint, dist []*prioq_item, path []int, q *prioq) {
	dv := uint(dist[v].val.(weight_priority))
	du := dist[u]
	if dv+weight < uint(du.val.(weight_priority)) {
		dist[u].val = weight_priority(dv + weight)
		heap.Fix(q, dist[u].heap_index)
		path[u] = v
	}
}
