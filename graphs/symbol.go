// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
