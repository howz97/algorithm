package lzw

import (
	"github.com/howz97/algorithm/search/hashmap"
	. "github.com/howz97/algorithm/util"
)

// Compress data using LZW algorithm
func Compress(data []byte) (out []byte) {
	table := hashmap.New[Str, uint16]()
	unused := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(Str([]byte{byte(b)}), unused)
		unused++
	}
	for len(data) > 0 {
		var code uint16
		i := 1
		for ; i <= len(data); i++ {
			c, ok := table.Get(Str(data[:i]))
			if !ok {
				break
			}
			code = c
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

// Decompress data compressed by LZW algorithm
func Decompress(data []byte) (out []byte) {
	table := hashmap.New[Int, []byte]()
	i := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(Int(i), []byte{byte(b)})
		i++
	}
	for {
		var code uint16
		code, data = readUint16(data)
		bytes, _ := table.Get(Int(code))
		out = append(out, bytes...)
		if len(data) == 0 {
			break
		}
		bytes1 := make([]byte, len(bytes))
		copy(bytes1, bytes)
		if e, ok := table.Get(Int(peekUint16(data))); ok {
			bytes1 = append(bytes1, e[0])
		} else {
			bytes1 = append(bytes1, bytes[0])
		}
		table.Put(Int(i), bytes1)
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
