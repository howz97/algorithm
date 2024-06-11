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
	"fmt"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	data, err := os.ReadFile("../../assets/compress/tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	compressed := HuffmanCompress(data)
	t.Logf("compress performance %.4f", float64(len(compressed))/float64(len(data)))
	decompressed, err := HuffmanDecompress(compressed)
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
	de, err := HuffmanDecompress(HuffmanCompress([]byte(data)))
	if err != nil {
		t.Error(err)
	}
	result := string(de)
	if data != result {
		t.Errorf("%v not equal %v", data, result)
	}
}

func Example() {
	data := []byte("zhang how, zhang how, zhang how, zhang how,")
	compressed := HuffmanCompress(data)
	fmt.Printf("performance %.4f\n", float64(len(compressed))/float64(len(data)))
	de, _ := HuffmanDecompress(compressed)
	fmt.Println(string(de))

	// Output:
	// performance 0.7674
	// zhang how, zhang how, zhang how, zhang how,
}
