package hash_map

import (
	"github.com/howz97/algorithm/search"
)

type Key interface {
	Hash() uint
	search.Cmp
}
