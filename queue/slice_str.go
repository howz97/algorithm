package queue

type SliStr struct {
	*Slice
}

func NewSliStr(cap int) *SliStr {
	return &SliStr{Slice: NewSlice(cap)}
}

func (q *SliStr) Front() string {
	return q.Slice.Front().(string)
}

func (q *SliStr) PushBack(s string) {
	q.Slice.PushBack(s)
}

func (q *SliStr) PopAll() []string {
	var strings []string
	for q.Size() > 0 {
		strings = append(strings, q.Front())
	}
	return strings
}
