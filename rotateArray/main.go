package main

func rotate(nums []int, k int) {
	for k > 0 {
		lastVal := nums[len(nums)-1]
		for i := len(nums) - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = lastVal
		k--
	}
}
