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
	HEXADECIMAL = DECIMAL + `ABCDEF`
	LOWERCASE   = `abcdefghijklmnopqrstuvwxyz`
	UPPERCASE   = `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	BASE64      = UPPERCASE + LOWERCASE + DECIMAL + `+/`
	ASCII       = ` !"#$%&'()*+,-./` + DECIMAL + `:;<=>?@` + UPPERCASE + `[\]^_` + "`" + LOWERCASE + `{|}~`
)

var Unicode = new(unicodeImpl)

// map unicode to other alphabet
type Interface interface {
	ToRune(rune) rune
	ToIndex(rune) rune
	Contains(rune) bool
	R() int
}

type alphabetImpl struct {
	r2i map[rune]rune
	i2r []rune
}

func NewAlphabetImpl(s string) *alphabetImpl {
	a := &alphabetImpl{
		r2i: make(map[rune]rune),
		i2r: make([]rune, 0),
	}
	i := rune(0)
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
func (a *alphabetImpl) ToRune(i rune) rune {
	if int(i) >= len(a.i2r) {
		panic(fmt.Sprintf("index %v exceed range of alphabetImpl", i))
	}
	return a.i2r[i]
}

// ToIndex convert rune to index
func (a *alphabetImpl) ToIndex(r rune) rune {
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

// ToIndices equal to call ToIndex for every rune in s
func (a *alphabetImpl) ToIndices(s string) []rune {
	indices := make([]rune, 0)
	for _, r := range s {
		indices = append(indices, a.ToIndex(r))
	}
	return indices
}

// ToRunes equal to call ToRune for every index in indices
func (a *alphabetImpl) ToRunes(indices []rune) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, a.ToRune(indices[i]))
	}
	return runes
}

func (a *alphabetImpl) Rand() rune {
	return a.i2r[rand.Intn(len(a.i2r))]
}

func (a *alphabetImpl) RandString(l int) string {
	str := ""
	for i := 0; i < l; i++ {
		str += string(a.Rand())
	}
	return str
}

type unicodeImpl struct{}

func (u *unicodeImpl) ToRune(i rune) rune {
	return i
}

func (u *unicodeImpl) ToIndex(r rune) rune {
	return r
}

func (u *unicodeImpl) Contains(_ rune) bool {
	return true
}

func (u *unicodeImpl) R() int {
	return 0xFFFF
}

// ToIndices equal to call ToIndex for every rune in s
func (u *unicodeImpl) ToIndices(s string) []rune {
	indices := make([]rune, 0)
	for _, r := range s {
		indices = append(indices, u.ToIndex(r))
	}
	return indices
}

func (u *unicodeImpl) ToRunes(indices []rune) []rune {
	runes := make([]rune, 0)
	for i := range indices {
		runes = append(runes, u.ToRune(indices[i]))
	}
	return runes
}
