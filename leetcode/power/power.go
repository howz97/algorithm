package power

import (
	"errors"
)

func Power(base float64, exponent int) (float64, error) {
	if base == 0 && exponent < 0 {
		return 0, errors.New("invalid input")
	}
	var absExponent uint
	if exponent < 0 {
		absExponent = uint(-exponent)
	} else {
		absExponent = uint(exponent)
	}
	result := UnsignExponent(base, absExponent)
	if exponent < 0 {
		result = 1.0 / result
	}
	return result, nil
}

func UnsignExponent(base float64, exponent uint) float64 {
	if exponent == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	if exponent == 1 {
		return base
	}
	if (exponent & 1) == 0 {
		return UnsignExponent(base*base, exponent>>1)
	} else {
		return UnsignExponent(base*base, exponent>>1) * base
	}
}
