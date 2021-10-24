package reverse

import (
	"math"

	"github.com/gbochora/leetcode/utils"
)

func reverse(x int) int {
	sign := utils.SignInt(x)
	xAbsVal := x * sign
	var revVal int
	for xAbsVal > 0 {
		if revVal > math.MaxInt32/10 || revVal*10 > math.MaxInt32-xAbsVal%10 {
			return 0
		}
		revVal = revVal*10 + xAbsVal%10
		xAbsVal /= 10
	}

	return sign * revVal
}
