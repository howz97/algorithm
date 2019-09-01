package alphabet

import (
	"fmt"
	"math"
)

const (
	BINARY      = "01"
	DNA         = "ACTG"
	OCTAL       = "01234567"
	DECIMAL     = "0123456789"
	HEXADECIMAL = "0123456789ABCDEF"
	LOWERCASE   = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	BASE64      = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	ASCII       = " !\"#$%&'()*+,-./" + DECIMAL + ":;<=>?@" + UPPERCASE + "[\\]^_`" + LOWERCASE + "{|}~"
)

var (
	Binary      = NewAlphabet(BINARY)
	Dna         = NewAlphabet(DNA)
	Octal       = NewAlphabet(OCTAL)
	Decimal     = NewAlphabet(DECIMAL)
	Hexadecimal = NewAlphabet(HEXADECIMAL)
	LowerCase   = NewAlphabet(LOWERCASE)
	UpperCase   = NewAlphabet(UPPERCASE)
	Base64      = NewAlphabet(BASE64)
	Ascii       = NewAlphabet(ASCII)
	Unicode     = new(AlphabetUnicode)
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

// ToRune convert index to rune
func (a *Alphabet) ToRune(i int) rune {
	r, exst := a.itor[i]
	if !exst {
		panic(fmt.Sprintf("index %v exceed range of Alphabet", i))
	}
	return r
}

// ToIndex convert rune to index
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

// R is the size of this Alphabet
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

type AlphabetUnicode struct{}

func (u *AlphabetUnicode) ToRune(i int) rune {
	return rune(i)
}

func (u *AlphabetUnicode) ToIndex(r rune) int {
	return int(r)
}

func (u *AlphabetUnicode) R() int {
	return 0xFFFF
}

// ToIndeices equal to call ToIndex for every rune in s
func (u *AlphabetUnicode) ToIndeices(s string) []int {
	indices := make([]int, 0)
	for _, r := range s {
		indices = append(indices, u.ToIndex(r))
	}
	return indices
}

// ToRunes equal to call ToRune for every index in indices
func (u *AlphabetUnicode) ToRunes(indices []int) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, u.ToRune(indices[i]))
	}
	return runes
}
