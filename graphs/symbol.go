package graphs

type Symbol struct {
	str2int map[string]int
	int2str []string
}

func NewSymbolGraph() *Symbol {
	return &Symbol{
		str2int: make(map[string]int),
		int2str: nil,
	}
}

func (sg *Symbol) scanVertical(v string) {
	if _, ok := sg.str2int[v]; !ok {
		sg.str2int[v] = len(sg.int2str)
		sg.int2str = append(sg.int2str, v)
	}
}

func (sg *Symbol) SymbolOf(v int) string {
	return sg.int2str[v]
}

func (sg *Symbol) VetOf(s string) int {
	return sg.str2int[s]
}
