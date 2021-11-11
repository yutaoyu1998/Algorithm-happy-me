package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	size := len(nums)
	for i := 0; i < size; i++ {
		n := target - nums[i]
		elem, ok := numMap[n]
		if ok {
			return []int{elem, i}
		} else {
			numMap[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	/* 测试 */
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(nums, target))
}
