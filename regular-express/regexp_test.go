package regexp

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	if !IsMatch(`(1(\\|\(|c)*2)`, `1c((\\2`) { // transfer
		t.Fail()
	}
	if IsMatch(`(1(a|b|c|d)+2)`, `12`) ||
		!IsMatch(`(1(a|b|c|d)+2)`, `1a2`) ||
			!IsMatch(`(1(a|b|c|d)+2)`, `1aabbdcdc2`) { // +
		t.Fail()
	}
	if !IsMatch(`(1(a|b|c|d)?2)`, `12`) ||
		!IsMatch(`(1(a|b|c|d)?2)`, `1d2`) ||
		IsMatch(`(1(a|b|c|d)?2)`, `1bc2`) { // ?
		t.Fail()
	}
	if IsMatch(`(1(a|b|c|d){3}2)`, `1ab2`) ||
		!IsMatch(`(1(a|b|c|d){3}2)`, `1abc2`) ||
		IsMatch(`(1(a|b|c|d){3}2)`, `1abcd2`) { // {n}
		t.Fail()
	}
	if !IsMatch(`(1(a|b|c|d){0-3}2)`, `12`) ||
		!IsMatch(`(1(a|b|c|d){0-3}2)`, `1abc2`) ||
		IsMatch(`(1(a|b|c|d){0-3}2)`, `1abcd2`) { // {n-m}
		t.Fail()
	}
}
