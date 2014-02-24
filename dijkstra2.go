package algo

import "container/heap"

func init_distance2(vcount, begin int) []*prioq_item {
	dist := make([]*prioq_item, vcount)
	for i := 0; i < len(dist); i++ {
		dist[i] = &prioq_item{max_weight, -1, i}
	}
	dist[begin].val = weight_priority(0)
	return dist
}
func dijkstra2(g graph, source int) []int {
	path := init_path(g.vertex_count(), source)
	dist := init_distance2(g.vertex_count(), source)
	distq := &prioq{}
	heap.Init(distq)
	heap.Push(distq, dist[source])
	for distq.Len() > 0 {
		u := heap.Pop(distq).(*prioq_item)
		u.heap_index = -2
		if u.val.(weight_priority) == max_weight {
			break
		}
		for _, edge := range g.edges(u.origin_index) {
			if dist[edge.v1()].heap_index == -2 { // has been poped out
				continue
			}
			relax2(u.origin_index, edge.v1(), edge.weight(), dist, path, distq)
		}
	}
	return path
}

func relax2(u, v int, weight uint, dist []*prioq_item, path []int, q *prioq) {
	du := uint(dist[u].val.(weight_priority))

	if dist[v].heap_index == -1 {
		heap.Push(q, dist[v]) // make du into q
	}
	if du+weight < uint(dist[v].val.(weight_priority)) {
		dist[v].val = weight_priority(du + weight)
		heap.Fix(q, dist[v].heap_index)
		path[u] = u
	}
}
