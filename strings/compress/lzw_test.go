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

func TestCompressLZW(t *testing.T) {
	data, err := os.ReadFile("../../assets/compress/tale_1.txt")
	if err != nil {
		t.Fatal(err)
	}
	compressed := LzwCompress(data)
	t.Logf("compress performance %v", float64(len(compressed))/float64(len(data)))
	decompressed := LzwDecompress(compressed)

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

func ExampleCompressLZW() {
	data := []byte("howhowhowhowhowhowhowhowhow")
	compressed := LzwCompress(data)
	fmt.Printf("performance %.4f\n", float64(len(compressed))/float64(len(data)))
	fmt.Println(string(LzwDecompress(compressed)))

	// Output:
	// performance 0.8889
	// howhowhowhowhowhowhowhowhow
}
