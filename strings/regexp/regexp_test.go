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

package regexp

import (
	"fmt"
	"regexp"
)

const (
	Yes = 0
	No  = 1
)

var match = map[string][2][]string{
	`(1(\\|\(|c)*2)`:     {Yes: {`12`, `1c((\\2`}, No: {`1*2`}}, // \
	`(1(a|b|c|d)+2)`:     {Yes: {`1a2`, `1aabbdcdc2`}, No: {`12`}},
	`(1a+2)`:             {Yes: {`1a2`, `1aaaaa2`}, No: {`12`}}, // +
	`(1(a|b|c|d)?2)`:     {Yes: {`12`, `1d2`}, No: {`1bc2`}},    // |
	`(1a?2)`:             {Yes: {`12`, `1a2`}, No: {`1aa2`}},
	`(1(a|b|c|d){3}2)`:   {Yes: {`1abc2`}, No: {`1ab2`, `1abcd2`}}, // {n}
	`(1a{3}2)`:           {Yes: {`1aaa2`}, No: {`1aa2`, `1aaaa2`}},
	`(1(a|b|c|豪){0-3}2)`: {Yes: {`12`, `1ab豪2`}, No: {`1abc豪2`}}, // {n-m}
	`(1a{0-3}2)`:         {Yes: {`12`, `1aaa2`}, No: {`1aaaa2`}},
	`.\..`:               {Yes: {`9.9`, `...`}, No: {`.豪.`, `.9`, `9.`, `..`}}, // .
	`.*`:                 {Yes: {``, `.`, `*`, `.*`, `any string`, `$%*^*^*&*)_`}, No: {}},
}

func Example() {
	for pattern, sli := range match {
		re, err := Compile(pattern)
		if err != nil {
			panic(err)
		}
		for _, str := range sli[Yes] {
			if !re.Match(str) {
				panic(fmt.Sprintf("%s not match %s", pattern, str))
			}
		}
		for _, str := range sli[No] {
			if re.Match(str) {
				panic(fmt.Sprintf("%s match %s", pattern, str))
			}
		}
	}

	invalid := []string{`\u12`, `{123)`}
	for _, p := range invalid {
		_, err := regexp.Compile(p)
		fmt.Println(p, err)
		if err == nil {
			panic(fmt.Sprintf("compile %s should not pass", p))
		}
	}
	fmt.Println("Passed!")

	// Output:
	// \u12 error parsing regexp: invalid escape sequence: `\u`
	// {123) error parsing regexp: unexpected ): `{123)`
	// Passed!
}
