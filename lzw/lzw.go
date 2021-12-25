package lzw

import (
	"github.com/howz97/algorithm/search/hash_map"
	. "github.com/howz97/algorithm/util"
)

func Compress(data []byte) (out []byte) {
	table := hash_map.New()
	unused := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(Str([]byte{byte(b)}), unused)
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
			out = append(out, byte(code>>8))
			out = append(out, byte(code))
			break
		}
		table.Put(Str(data[:i]), unused)
		unused++
		out = append(out, byte(code>>8))
		out = append(out, byte(code))
		data = data[i-1:]
	}
	return
}

func Decompress(data []byte) (out []byte) {
	table := hash_map.New()
	i := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(Int(i), []byte{byte(b)})
		i++
	}
	for {
		var code uint16
		code, data = readUint16(data)
		bytes := table.Get(Int(code)).([]byte)
		out = append(out, bytes...)
		if len(data) == 0 {
			break
		}
		code1 := peekUint16(data)
		if e := table.Get(Int(code1)); e != nil {
			bytes = append(bytes, e.([]byte)[0])
		} else {
			bytes = append(bytes, bytes[0])
		}
		table.Put(Int(i), bytes)
		i++
	}
	return
}

func readUint16(data []byte) (uint16, []byte) {
	return peekUint16(data), data[2:]
}

func peekUint16(data []byte) uint16 {
	return (uint16(data[0]) << 8) | uint16(data[1])
}
