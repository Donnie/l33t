func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)

	l := len(nums)
	if l == 0 {
		return 0
	} else if l%2 == 0 {
		a := float64(nums[l/2-1])
		b := float64(nums[l/2])
		return (a + b) / 2
	} else {
		return float64(nums[l/2])
	}
}
