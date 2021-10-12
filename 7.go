package main

import "math"

func reverse(x int) int {
	sign := SignInt(x)
	xAbsVal := x * sign
	var revVal int
	for xAbsVal > 0 {
		if revVal > math.MaxInt32/10 || revVal*10 > math.MaxInt32-xAbsVal%10 {
			return 0
		}
		revVal = revVal*10 + xAbsVal%10
		xAbsVal /= 10
	}

	return revVal
}

func SignInt(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}
