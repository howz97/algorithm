package alphabet

import (
	"fmt"
	"math"
)

// Alphabet represent an alphabet
type Alphabet struct {
	rtoi map[rune]int
	itor map[int]rune
}

// NewAlphabet assume every rune in s is unique
func NewAlphabet(s string) *Alphabet {
	a := &Alphabet{
		rtoi: make(map[rune]int),
		itor: make(map[int]rune),
	}
	for i, r := range s {
		a.rtoi[r] = i
		a.itor[i] = r
	}
	return a
}

// ToRune -
func (a *Alphabet) ToRune(i int) rune {
	r, exst := a.itor[i]
	if !exst {
		panic(fmt.Sprintf("index %v exceed range of Alphabet", i))
	}
	return r
}

// ToIndex -
func (a *Alphabet) ToIndex(r rune) int {
	i, exst := a.rtoi[r]
	if !exst {
		panic(fmt.Sprintf("rune %v do not belong to Alphabet", r))
	}
	return i
}

// Contains -
func (a *Alphabet) Contains(r rune) bool {
	_, exst := a.rtoi[r]
	return exst
}

// R -
func (a *Alphabet) R() int {
	return len(a.rtoi)
}

// lgR means the number of bits needed to represent R
func (a *Alphabet) lgR() int {
	logarithm := math.Log2(float64(a.R()))
	if logarithm > math.Logb(float64(a.R())) {
		return int(logarithm + 1)
	}
	return int(logarithm)
}

// ToIndeices equal to call ToIndex for every rune in s
func (a *Alphabet) ToIndeices(s string) []int {
	indices := make([]int, 0)
	for _, r := range s {
		indices = append(indices, a.ToIndex(r))
	}
	return indices
}

// ToRunes equal to call ToRune for every index in indices
func (a *Alphabet) ToRunes(indices []int) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, a.ToRune(indices[i]))
	}
	return runes
}
