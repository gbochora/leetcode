package binarysearch

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		mid := (len(nums2) - 1) / 2
		if len(nums2)%2 == 1 {
			return float64(nums2[mid])
		}
		return float64(nums2[mid]+nums2[mid+1]) / 2
	}
	if len(nums2) == 0 {
		mid := (len(nums1) - 1) / 2
		if len(nums1)%2 == 1 {
			return float64(nums1[mid])
		}
		return float64(nums1[mid]+nums1[mid+1]) / 2
	}

	return findMedianSortedArraysRec(nums1, nums2, 0, len(nums1)-1, 0, len(nums2)-1, (len(nums1)+len(nums2)-1)/2)
}

func findMedianSortedArraysRec(nums1 []int, nums2 []int, low1, high1, low2, high2, mid int) float64 {
	mid1 := (low1 + high1) / 2
	mid2 := (low2 + high2) / 2
	if low1 > high1 {
		mid2LeftCount := mid1 + mid2
		if nums2[mid2] > nums1[mid1] {
			mid2LeftCount++
		}
		if mid2LeftCount > mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, high1, low2, mid2-1, mid)
		}
		if mid2LeftCount < mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, high1, mid2+1, high2, mid)
		}

		if (len(nums1)+len(nums2))%2 == 1 {
			return float64(nums2[mid2])
		}
		next := nums1[mid1]
		if next < nums2[mid2] && mid1+1 < len(nums1) {
			next = nums1[mid1+1]
		}
		if next < nums2[mid2] || (mid2+1 < len(nums2) && next > nums2[mid2+1]) {
			next = nums2[mid2+1]
		}
		return float64(nums2[mid2]+next) / 2
	}
	if low2 > high2 {
		mid1LeftCount := mid1 + mid2
		if nums1[mid1] > nums2[mid2] {
			mid1LeftCount++
		}
		if mid1LeftCount > mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, mid1-1, low2, high2, mid)
		}
		if mid1LeftCount < mid {
			return findMedianSortedArraysRec(nums1, nums2, mid1+1, high1, low2, high2, mid)
		}

		if (len(nums1)+len(nums2))%2 == 1 {
			return float64(nums1[mid1])
		}
		next := nums2[mid2]
		if next < nums1[mid1] && mid2+1 < len(nums2) {
			next = nums2[mid2+1]
		}
		if next < nums1[mid1] || (mid1+1 < len(nums1) && next > nums1[mid1+1]) {
			next = nums1[mid1+1]
		}
		return float64(nums1[mid1]+next) / 2
	}

	if nums1[mid1] > nums2[mid2] {
		mid2LeftMaxCount := mid2 + mid1
		mid1LeftMinCount := mid1 + mid2 + 1
		if mid2LeftMaxCount < mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, high1, mid2+1, high2, mid)
		}
		if mid1LeftMinCount > mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, mid1-1, low2, high2, mid)
		}
	}
	if nums1[mid1] < nums2[mid2] {
		mid1MaxCount := mid2 + mid1
		mid2MinCount := mid1 + mid2 + 1
		if mid1MaxCount < mid {
			return findMedianSortedArraysRec(nums1, nums2, mid1+1, high1, low2, high2, mid)
		}
		if mid2MinCount > mid {
			return findMedianSortedArraysRec(nums1, nums2, low1, high1, low2, mid2-1, mid)
		}
	}
	mid1Count := mid1 + mid2
	if mid1Count == mid {
		return float64(nums1[mid1])
	}
	mid1Count++
	if mid1Count < mid {
		return findMedianSortedArraysRec(nums1, nums2, mid1+1, high1, mid2+1, high2, mid)
	}
	if mid1Count > mid {
		return findMedianSortedArraysRec(nums1, nums2, low1, mid1-1, low2, mid2-1, mid)
	}
	if (len(nums1)+len(nums2))%2 == 1 {
		return float64(nums1[mid1])
	}
	var next int
	if mid1+1 >= len(nums1) {
		next = nums2[mid2+1]
	} else {
		next = nums1[mid1+1]
	}
	if mid1+1 < len(nums1) && next > nums1[mid1+1] {
		next = nums1[mid1+1]
	}
	if mid2+1 < len(nums2) && next > nums2[mid2+1] {
		next = nums2[mid2+1]
	}
	return float64(nums1[mid1]+next) / 2
}
