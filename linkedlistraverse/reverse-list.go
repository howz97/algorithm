package linkedlistreverse

type node struct {
	v    string
	next *node
}

func printListReverse(l *node) {
	if l == nil {
		return
	}
	if l.next == nil {
		println(l.v)
		return
	}
	printListReverse(l.next)
	println(l.v)
}
