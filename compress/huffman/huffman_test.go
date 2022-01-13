package huffman

import (
	"io/ioutil"
	"testing"
)

func TestCompress(t *testing.T) {
	data, err := ioutil.ReadFile("..\\testdata\\tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	compressed := Compress(data)
	t.Logf("compress performance %.4f", float64(len(compressed))/float64(len(data)))
	decompressed, err := Decompress(compressed)
	if err != nil {
		t.Error(err)
	}

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
	de, err := Decompress(Compress([]byte(data)))
	if err != nil {
		t.Error(err)
	}
	result := string(de)
	if data != result {
		t.Errorf("%v not equal %v", data, result)
	}
}
