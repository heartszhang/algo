package algo

import (
	"math/rand"
	"testing"
)

type graph_matrix [][]uint // weight[i][j] = the weight of edge from i to j

func (self graph_matrix) vertex_count() int {
	return len(self)
}

func (self graph_matrix) edges(v int) []edge {
	x := self[v]
	var val []edge
	for i, e := range x {
		if e > 0 {
			val = append(val, edge_ray{e, i})
		}
	}
	return val
}

type edge_ray struct {
	_weight uint
	end     int
}

func (self edge_ray) v0() int {
	return -1
}
func (self edge_ray) v1() int {
	return self.end
}
func (self edge_ray) weight() uint {
	return self._weight
}

func new_graph_matrix(vertex_count int) graph_matrix {
	v := make(graph_matrix, vertex_count)
	for i := 0; i < vertex_count; i++ {
		v[i] = make([]uint, vertex_count)
	}
	return v
}

var (
	g  graph_matrix
	vc int
)

func init() {
	vc = 10000
	g = new_graph_matrix(vc)
	for i := 0; i < vc; i++ {
		for j := 0; j < vc; j++ {
			if i != j {
				x := rand.Intn(2)
				if x == 0 {
					g[i][j] = uint(rand.Intn(12))
				}
			}
		}
	}
}

func TestDijkstra2(t *testing.T) {
	dijkstra(g, 3)
}
func TestDijkstra(t *testing.T) {
	dijkstra(g, 3)
}

func print_path(path []int, idx int, t *testing.T) {
	oi := idx
	x := []int{}
	for path[idx] >= 0 {
		x = append([]int{path[idx]}, x...)
		idx = path[idx]
	}
	t.Log(x, oi)
}

func BenchmarkDijkstra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dijkstra(g, 1)
	}
}
func BenchmarkDijkstra2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dijkstra2(g, 1)
	}
}
