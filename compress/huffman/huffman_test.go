package huffman

import (
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	data, err := ioutil.ReadFile("..\\testdata\\tale_1.txt")
	if err != nil {
		t.Fatal(err)
	}
	compressed := Compress(data)
	t.Logf("compress performance %v", float64(len(compressed))/float64(len(data)))
	decompressed := Decompress(compressed)

	if len(data) != len(decompressed) {
		t.Fatalf("raw data length is %v, decompressed length is %v", len(data), len(decompressed))
	}
	for i, b := range data {
		if b != decompressed[i] {
			t.Log(data, string(data))
			t.Log(compressed)
			t.Log(decompressed, string(decompressed))
			t.Errorf("diffence at %d: raw data is %v, decompressed is %v", i, b, decompressed[i])
			break
		}
	}
}

func TestSimple(t *testing.T) {
	data := "zhang how 1997"
	result := string(Decompress(Compress([]byte(data))))
	if data != result {
		t.Errorf("%v not equal %v", data, result)
	}
}
