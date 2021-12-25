package lzw

import (
	"github.com/howz97/algorithm/search/hash_map"
	. "github.com/howz97/algorithm/util"
)

func Compress(data []byte) (out []byte) {
	table := hash_map.New()
	var (
		key  []byte
		code uint16
	)
	for _, b := range data {
		key = append(key, b)
		e := table.Get(Str(key))
		if e != nil {
			continue
		}
		code++
		table.Put(Str(key), code)
		key = key[:0]
		out = append(out, byte(code))
		out = append(out, byte(code>>8))
	}
	return
}

func Decompress(data []byte) []byte {

}
