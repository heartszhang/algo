package algo

type priority interface {
	less(rhs priority) bool
}

type prioq_item struct {
	val          priority
	heap_index   int
	origin_index int
}

type prioq []*prioq_item

// sort interface
func (h prioq) Len() int           { return len(h) }
func (h prioq) Less(i, j int) bool { return h[i].val.less(h[j].val) }
func (h prioq) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heap_index = i
	h[j].heap_index = j
}

// heap interface
func (h *prioq) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	v.heap_index = -1
	return v
}

func (h *prioq) Push(x interface{}) {
	i := x.(*prioq_item)
	i.heap_index = len(*h)
	*h = append(*h, i)
}
