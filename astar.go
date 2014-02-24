package algo

import "container/heap"

func astar(g astar_graph, start, goal int) []int {
	path := init_path(g.vertex_count(), start) // -1 : start point, -2 unreachable point

	scores := init_scores(g.vertex_count(), start, goal, g.heuristic_estimate(start, goal))
	q := prioq(scores[:])
	heap.Init(&q)
	for q.Len() > 0 {
		u := heap.Pop(&q).(*prioq_item)
		if u.val.(astar_item).weight == max_uint || u.origin_index == goal {
			break
		}
		for _, edge := range g.edges(u.origin_index) {
			if scores[edge.v1()].heap_index < 0 {
				continue
			}
			astar_relax(u.origin_index, edge.v1(), edge.weight(), scores, &q, path, g.heuristic_estimate(edge.v1(), goal))
		}
	}
	return path
}

func astar_relax(u, v int, weight uint, scores []*prioq_item, q *prioq, path []int, est uint) {
	du := scores[u].val.(astar_item).weight

	if du+weight < scores[v].val.(astar_item).weight {
		scores[v].val = astar_item{du + weight, du + weight + est}
		heap.Fix(q, scores[v].heap_index)
		path[v] = u
	}
}

const (
	max_uint = ^uint(0)
)

type astar_item struct {
	weight    uint
	estimated uint
}

func (lhs astar_item) less(rhs priority) bool {
	r := rhs.(astar_item)
	return lhs.estimated < r.estimated
}
func init_scores(vcount, start, goal int, startest uint) []*prioq_item {
	val := make([]*prioq_item, vcount)
	for i := 0; i < vcount; i++ {
		val[i] = &prioq_item{astar_item{max_uint, max_uint}, i, i}
	}
	val[start].val = astar_item{0, startest}
	return val
}
