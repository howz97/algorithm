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

const primeRK = 16777619 // must greater than 256

func IndexRabinKarp(s, substr string) int {
	if len(s) == 0 || len(substr) == 0 || len(s) < len(substr) {
		return -1
	}
	m := len(substr)
	if s[:m] == substr {
		return 0
	}
	hp, ht := hashStr(substr), hashStr(s[:m])
	rpm := pow(primeRK, uint32(m-1))
	for i := m; i < len(s); i++ {
		ht = (ht-uint32(s[i-m])*rpm)*primeRK + uint32(s[i])
		if ht == hp && s[i-m+1:i+1] == substr {
			return i - m + 1
		}
	}
	return -1
}

func hashStr(sep string) uint32 {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	return hash
}

func pow(x, y uint32) uint32 {
	if y == 0 {
		return 1
	}
	v := uint32(1)
	for i := uint32(0); i < y; i++ {
		v *= x
	}
	return v
}
