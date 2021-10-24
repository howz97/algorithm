package main

import (
	"fmt"
	"github.com/howz97/algorithm/regexp"
)

const (
	Match   = 0
	UnMatch = 1
)

var match = map[string][2][]string{
	`(1(\\|\(|c)*2)`:     {Match: {`12`, `1c((\\2`}, UnMatch: {`1*2`}}, // \
	`(1(a|b|c|d)+2)`:     {Match: {`1a2`, `1aabbdcdc2`}, UnMatch: {`12`}},
	`(1a+2)`:             {Match: {`1a2`, `1aaaaa2`}, UnMatch: {`12`}}, // +
	`(1(a|b|c|d)?2)`:     {Match: {`12`, `1d2`}, UnMatch: {`1bc2`}},    // |
	`(1a?2)`:             {Match: {`12`, `1a2`}, UnMatch: {`1aa2`}},
	`(1(a|b|c|d){3}2)`:   {Match: {`1abc2`}, UnMatch: {`1ab2`, `1abcd2`}}, // {n}
	`(1a{3}2)`:           {Match: {`1aaa2`}, UnMatch: {`1aa2`, `1aaaa2`}},
	`(1(a|b|c|豪){0-3}2)`: {Match: {`12`, `1ab豪2`}, UnMatch: {`1abc豪2`}}, // // {n-m}
	`(1a{0-3}2)`:         {Match: {`12`, `1aaa2`}, UnMatch: {`1aaaa2`}},
}

func main() {
	for pattern, sli := range match {
		for _, str := range sli[Match] {
			if !regexp.IsMatch(pattern, str) {
				panic(fmt.Sprintf("%s not match %s", pattern, str))
			}
		}
		for _, str := range sli[UnMatch] {
			if regexp.IsMatch(pattern, str) {
				panic(fmt.Sprintf("%s match %s", pattern, str))
			}
		}
	}
}
