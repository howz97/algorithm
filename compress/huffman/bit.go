package huffman

import (
	"io"
)

type bitReader struct {
	data []byte
	n    uint64
	bits uint
	err  error
}

func newBitReader(data []byte) *bitReader {
	return &bitReader{data: data}
}

// ReadBits64 reads the given number of bits and returns them in the
// least-significant part of a uint64. In the event of an error, it returns 0
// and the error can be obtained by calling Err().
func (br *bitReader) ReadBits64(bits uint) (n uint64) {
	for bits > br.bits {
		if len(br.data) == 0 {
			br.err = io.ErrUnexpectedEOF
			return 0
		}
		b := br.data[0]
		br.data = br.data[1:]
		br.n <<= 8
		br.n |= uint64(b)
		br.bits += 8
	}

	// br.n looks like this (assuming that br.bits = 14 and bits = 6):
	// Bit: 111111
	//      5432109876543210
	//
	//         (6 bits, the desired output)
	//        |-----|
	//        V     V
	//      0101101101001110
	//        ^            ^
	//        |------------|
	//           br.bits (num valid bits)
	//
	// This the next line right shifts the desired bits into the
	// least-significant places and masks off anything above.
	n = (br.n >> (br.bits - bits)) & ((1 << bits) - 1)
	br.bits -= bits
	return
}

func (br *bitReader) ReadBits(bits uint) (n int) {
	n64 := br.ReadBits64(bits)
	return int(n64)
}

func (br *bitReader) ReadBit() bool {
	n := br.ReadBits(1)
	return n != 0
}

func (br *bitReader) Err() error {
	return br.err
}

type bitWriter struct {
	bits   byte
	nBits  uint8
	output []byte
}

func (bw *bitWriter) WriteBit(bit bool) {
	if bit {
		bw.bits |= byte(1<<7) >> bw.nBits
	}
	bw.nBits++
	if bw.nBits == 8 {
		bw.output = append(bw.output, bw.bits)
		bw.bits = 0
		bw.nBits = 0
	}
}

func (bw *bitWriter) WriteBits(bits []bool) {
	for _, bit := range bits {
		bw.WriteBit(bit)
	}
}

func (bw *bitWriter) WriteByte(b byte) {
	for i := uint(7); i > 0; i-- {
		bw.WriteBit((b >> i & 1) == 1)
	}
	bw.WriteBit((b & 1) == 1)
}

func (bw *bitWriter) WriteUint32(n uint32) {
	for i := uint(31); i > 0; i-- {
		bw.WriteBit((n >> i & 1) == 1)
	}
	bw.WriteBit((n & 1) == 1)
}

func (bw *bitWriter) Close() {
	if bw.nBits > 0 {
		bw.output = append(bw.output, bw.bits)
		bw.bits = 0
		bw.nBits = 0
	}
}
