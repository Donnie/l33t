func twoSum(nums []int, target int) (out []int) {
	seen := make(map[int]int)
	for plc1, num1 := range nums {
		complement := target - num1
		ind, found := seen[complement]
		if found {
			return []int{plc1, ind}
		}
        seen[num1] = plc1
	}
	return
}
