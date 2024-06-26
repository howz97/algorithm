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

package util

import (
	"unsafe"
)

type Int int

func (v Int) Hash() uintptr {
	return uintptr(v)
}

type Float float64

func (v Float) Hash() uintptr {
	return *(*uintptr)(unsafe.Pointer(&v))
}

const littlePrime = 31

type Str string

func (v Str) Hash() uintptr {
	h := uintptr(0)
	for i := 0; i < len(v); i++ {
		h = h*littlePrime + uintptr(v[i])
	}
	return h
}
