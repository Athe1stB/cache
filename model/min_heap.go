package model

type Ordered interface {
	~int | ~int32 | ~int64 | ~float64
}

type MinHeap[F Ordered, S any] []Pair[F, S]

func (h MinHeap[F, S]) Less(i, j int) bool { return h[i].First < h[j].First }
func (h MinHeap[F, S]) Len() int           { return len(h) }
func (h MinHeap[F, S]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap[F, S]) Push(el any) {
	*h = append(*h, el.(Pair[F, S]))
}

func (h *MinHeap[F, S]) Pop() any {
	old := *h
	n := len(old)
	last_element := old[n-1]
	*h = old[:n-1]
	return last_element
}

func (h *MinHeap[F, S]) Update(index int, value Pair[F, S]) {
	(*h)[index] = value
}
