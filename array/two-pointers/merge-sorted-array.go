package two_pointers

/*
*
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
and the last n elements are set to 0 and should be ignored. nums2 has a length of n.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}

func Merge(nums1 []int, m int, nums2 []int, n int) {
	pt1, pt2, pt3 := m-1, n-1, m+n-1

	for ; pt1 >= 0 && pt2 >= 0; pt3-- {
		if nums1[pt1] >= nums2[pt2] {
			nums1[pt3] = nums1[pt1]
			pt1--
		} else {
			nums1[pt3] = nums2[pt2]
			pt2--
		}
	}

	if pt2 >= 0 {
		copy(nums1[:pt3+1], nums2[:pt2+1])
	}
}
