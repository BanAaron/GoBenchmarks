package main

import "math"

type signedNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func ReverseSignMultiplication[sn signedNumber](number sn) sn {
	number *= -1
	return number
}

func ReverseSignMinus[sn signedNumber](number sn) sn {
	number = -number
	return number
}

func ReverseSignChaos(number float64) float64 {
	if number < 0 {
		number = math.Abs(number)
	} else {
		number = number - (number * 2)
	}
	return number
}
