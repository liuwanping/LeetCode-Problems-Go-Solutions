package main

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		value, ok := m[target-num]
		if ok && value != i {
			return []int{i, value}
		}
		m[num] = i
	}
	return res
}
