package regexp

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	if !IsMatch(`(1(\\|\(|c)*2)`, `1c((\\2`) { // transfer
		t.Fail()
	}
	// +
	if IsMatch(`(1(a|b|c|d)+2)`, `12`) ||
		!IsMatch(`(1(a|b|c|d)+2)`, `1a2`) ||
		!IsMatch(`(1(a|b|c|d)+2)`, `1aabbdcdc2`) {
		t.Fail()
	}
	if IsMatch(`(1a+2)`, `12`) ||
		!IsMatch(`(1a+2)`, `1a2`) ||
		!IsMatch(`(1a+2)`, `1aaaaa2`) {
		t.Fail()
	}

	// ?
	if !IsMatch(`(1(a|b|c|d)?2)`, `12`) ||
		!IsMatch(`(1(a|b|c|d)?2)`, `1d2`) ||
		IsMatch(`(1(a|b|c|d)?2)`, `1bc2`) {
		t.Fail()
	}
	if !IsMatch(`(1a?2)`, `12`) ||
		!IsMatch(`(1a?2)`, `1a2`) ||
		IsMatch(`(1a?2)`, `1aa2`) {
		t.Fail()
	}

	// {n}
	if IsMatch(`(1(a|b|c|d){3}2)`, `1ab2`) ||
		!IsMatch(`(1(a|b|c|d){3}2)`, `1abc2`) ||
		IsMatch(`(1(a|b|c|d){3}2)`, `1abcd2`) {
		t.Fail()
	}
	if IsMatch(`(1a{3}2)`, `1aa2`) ||
		!IsMatch(`(1a{3}2)`, `1aaa2`) ||
		IsMatch(`(1a{3}2)`, `1aaaa2`) {
		t.Fail()
	}

	// {n-m}
	if !IsMatch(`(1(a|b|c|豪){0-3}2)`, `12`) ||
		!IsMatch(`(1(a|b|c|豪){0-3}2)`, `1ab豪2`) ||
		IsMatch(`(1(a|b|c|豪){0-3}2)`, `1abc豪2`) {
		t.Fail()
	}
	if !IsMatch(`(1a{0-3}2)`, `12`) ||
		!IsMatch(`(1a{0-3}2)`, `1aaa2`) ||
		IsMatch(`(1a{0-3}2)`, `1aaaa2`) {
		t.Fail()
	}

	//fmt.Println("==>> ", string(compile([]rune(`(1a{0-3}2)`))))
}
