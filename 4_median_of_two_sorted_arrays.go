package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 让n>=m, 保证i和j始终都是正数
	n, m := len(nums1), len(nums2)
	if m > n {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
		n, m = len(nums1), len(nums2)
	}
	minI, maxI := 0, m
	halfLen := (m + n + 1) / 2
	var res float64
	for minI <= maxI {
		i := (minI + maxI) / 2
		j := halfLen - i
		if i < m && nums1[j-1] > nums2[i] {
			minI = i + 1
		} else if i > 0 && nums2[i-1] > nums1[j] {
			maxI = i - 1
		} else {
			var maxLeft int
			if i == 0 {
				maxLeft = nums1[j-1]
			} else if j == 0 {
				maxLeft = nums2[i-1]
			} else {
				if nums1[j-1] > nums2[i-1] {
					maxLeft = nums1[j-1]
				} else {
					maxLeft = nums2[i-1]
				}
			}
			//一共奇数个元素
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			//一共偶数个元素
			var minRight int
			if i == m {
				minRight = nums1[j]
			} else if j == n {
				minRight = nums2[i]
			} else {
				if nums1[j] > nums2[i] {
					minRight = nums2[i]
				} else {
					minRight = nums1[j]
				}
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return res
}
