package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSum(nums, target)
	fmt.Printf("Вход: nums = %v, target = %d\n", nums, target)
	fmt.Printf("Результат (Brute Force): %v\n", result)
}
