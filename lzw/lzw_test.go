package lzw

import "testing"

func TestCompress(t *testing.T) {
	data := []byte("abababaabskabaab")
	compressed := Compress(data)
	t.Logf("compress performance %v", float64(len(compressed))/float64(len(data)))
	decompressed := Decompress(compressed)

	if len(data) != len(decompressed) {
		t.Fatalf("raw data length is %v, decompressed length is %v", len(data), len(decompressed))
	}
	for i, b := range data {
		if b != decompressed[i] {
			t.Log(data)
			t.Log(compressed)
			t.Log(decompressed)
			t.Errorf("diffence at %d: raw data is %v, decompressed is %v", i, b, decompressed[i])
			break
		}
	}
}
