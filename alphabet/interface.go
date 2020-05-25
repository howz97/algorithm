package alphabet

type IAlphabet interface {
	ToRune(i int) rune
	ToIndex(r rune) int
	Contains(r rune) bool
	R() int
}
