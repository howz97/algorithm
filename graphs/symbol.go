package graphs

func NewSymbolGraph() *Symbol {
	return &Symbol{
		syb2vet: make(map[string]int),
		vet2syb: nil,
	}
}

type Symbol struct {
	syb2vet map[string]int
	vet2syb []string
}

func (sg *Symbol) scanVertical(v string) {
	if _, ok := sg.syb2vet[v]; !ok {
		sg.syb2vet[v] = len(sg.vet2syb)
		sg.vet2syb = append(sg.vet2syb, v)
	}
}

func (sg *Symbol) SymbolOf(v int) string {
	return sg.vet2syb[v]
}

func (sg *Symbol) VetOf(s string) int {
	v, ok := sg.syb2vet[s]
	if !ok {
		return -1
	}
	return v
}
