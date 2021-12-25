package lzw

import (
	"github.com/howz97/algorithm/search/hash_map"
	. "github.com/howz97/algorithm/util"
)

func Compress(data []byte) (out []byte) {
	table := hash_map.New()
	unused := uint16(1)
	for i := 0; i <= 0xFF; i++ {
		table.Put(Str([]byte{byte(i)}), unused)
		unused++
	}
	for len(data) > 0 {
		var code uint16
		i := 1
		for ; i <= len(data); i++ {
			e := table.Get(Str(data[:i]))
			if e == nil {
				break
			}
			code = e.(uint16)
		}
		if i > len(data) {
			out = append(out, byte(code))
			out = append(out, byte(code>>8))
			break
		}
		table.Put(Str(data[:i]), unused)
		unused++
		out = append(out, byte(code))
		out = append(out, byte(code>>8))
		data = data[i-1:]
	}
	return
}

func Decompress(data []byte) []byte {

}
