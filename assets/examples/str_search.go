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

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/howz97/algorithm/strings"
)

func demo_strSearch() {
	txt, err := ioutil.ReadFile("../strings/tale.txt")
	if err != nil {
		panic(err)
	}
	pattern := "It is a far, far better thing that I do, than I have ever done"
	searcher := strings.NewKMP(pattern)
	//searcher := strings.NewBM(pattern)
	i := searcher.Index(string(txt))
	//i := strings.IndexRabinKarp(string(txt), pattern)
	fmt.Println(string(txt[i-50 : i+100]))

	/*
		 story, with a tender and a faltering
		voice.

		"It is a far, far better thing that I do, than I have ever done;
		it is a far, far better rest that I
	*/
}
