// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compress

import (
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/util"
)

// LzwCompress data using LZW algorithm
func LzwCompress(data []byte) (out []byte) {
	table := search.NewHashMap[util.Str, uint16]()
	unused := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(util.Str([]byte{byte(b)}), unused)
		unused++
	}
	for len(data) > 0 {
		var code uint16
		i := 1
		for ; i <= len(data); i++ {
			c, ok := table.Get(util.Str(data[:i]))
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
		table.Put(util.Str(data[:i]), unused)
		unused++
		out = append(out, byte(code>>8))
		out = append(out, byte(code))
		data = data[i-1:]
	}
	return
}

// LzwDecompress data compressed by LZW algorithm
func LzwDecompress(data []byte) (out []byte) {
	table := search.NewHashMap[util.Int, []byte]()
	i := uint16(0)
	for b := 0; b <= 0xFF; b++ {
		table.Put(util.Int(i), []byte{byte(b)})
		i++
	}
	for {
		var code uint16
		code, data = readUint16(data)
		bytes, _ := table.Get(util.Int(code))
		out = append(out, bytes...)
		if len(data) == 0 {
			break
		}
		bytes1 := make([]byte, len(bytes))
		copy(bytes1, bytes)
		if e, ok := table.Get(util.Int(peekUint16(data))); ok {
			bytes1 = append(bytes1, e[0])
		} else {
			bytes1 = append(bytes1, bytes[0])
		}
		table.Put(util.Int(i), bytes1)
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
