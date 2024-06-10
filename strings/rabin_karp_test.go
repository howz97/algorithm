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

package strings

import "testing"

func TestIndexRabinKarp(t *testing.T) {
	txt := "kadskdkndkqbabc"
	pattern := "abc"
	i := IndexRabinKarp(txt, pattern)
	if txt[i:i+len(pattern)] != pattern {
		t.Fatalf("not match, i=%d", i)
	}
}
