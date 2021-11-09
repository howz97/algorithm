package alphabet

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	BINARY      = `01`
	DNA         = `ACTG`
	OCTAL       = `01234567`
	DECIMAL     = `0123456789`
	HEXADECIMAL = `0123456789ABCDEF`
	LOWERCASE   = `abcdefghijklmnopqrstuvwxyz`
	UPPERCASE   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	BASE64      = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/`
	ASCII       = ` !"#$%&'()*+,-./` + DECIMAL + `:;<=>?@` + UPPERCASE + `[\]^_` + "`" + LOWERCASE + `{|}~`
)

var (
	LowerCase   = NewAlphabetImpl(LOWERCASE)
	Ascii       = NewAlphabetImpl(ASCII)
	Unicode     = new(unicodeImpl)
)

type Interface interface {
	ToRune(int) rune
	ToIndex(rune) int
	Contains(rune) bool
	R() int
}

type alphabetImpl struct {
	r2i map[rune]int
	i2r []rune
}

func NewAlphabetImpl(s string) *alphabetImpl {
	a := &alphabetImpl{
		r2i: make(map[rune]int),
		i2r: make([]rune, 0),
	}
	i := 0
	for _, r := range s {
		if _, exist := a.r2i[r]; exist {
			continue
		}
		a.r2i[r] = i
		a.i2r = append(a.i2r, r)
		i++
	}
	return a
}

// ToRune convert index to rune
func (a *alphabetImpl) ToRune(i int) rune {
	if i >= len(a.i2r) {
		panic(fmt.Sprintf("index %v exceed range of alphabetImpl", i))
	}
	return a.i2r[i]
}

// ToIndex convert rune to index
func (a *alphabetImpl) ToIndex(r rune) int {
	i, exst := a.r2i[r]
	if !exst {
		panic(fmt.Sprintf("rune %v do not belong to alphabetImpl", string(r)))
	}
	return i
}

func (a *alphabetImpl) Contains(r rune) bool {
	_, exst := a.r2i[r]
	return exst
}

// R is the size of this Alphabet
func (a *alphabetImpl) R() int {
	return len(a.r2i)
}

// lgR means the number of bits needed to represent R
func (a *alphabetImpl) lgR() int {
	logarithm := math.Log2(float64(a.R()))
	if logarithm > math.Logb(float64(a.R())) {
		return int(logarithm + 1)
	}
	return int(logarithm)
}

// ToIndeices equal to call ToIndex for every rune in s
func (a *alphabetImpl) ToIndeices(s string) []int {
	indices := make([]int, 0)
	for _, r := range s {
		indices = append(indices, a.ToIndex(r))
	}
	return indices
}

// ToRunes equal to call ToRune for every index in indices
func (a *alphabetImpl) ToRunes(indices []int) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, a.ToRune(indices[i]))
	}
	return runes
}

func (a *alphabetImpl) Rand() rune {
	return a.i2r[rand.Intn(len(a.i2r))]
}

type unicodeImpl struct{}

func (u *unicodeImpl) ToRune(i int) rune {
	return rune(i)
}

func (u *unicodeImpl) ToIndex(r rune) int {
	return int(r)
}

func (u *unicodeImpl) Contains(_ rune) bool {
	return true
}

func (u *unicodeImpl) R() int {
	return 0xFFFF
}

// ToIndeices equal to call ToIndex for every rune in s
func (u *unicodeImpl) ToIndeices(s string) []int {
	indices := make([]int, 0)
	for _, r := range s {
		indices = append(indices, u.ToIndex(r))
	}
	return indices
}

// ToRunes equal to call ToRune for every index in indices
func (u *unicodeImpl) ToRunes(indices []int) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, u.ToRune(indices[i]))
	}
	return runes
}
