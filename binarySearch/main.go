package main

func search(nums []int, target int) int {
	lenght := len(nums)
	if lenght < 0 {
		return -1
	}
	return binarySearch(nums, 0, lenght-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	key := nums[mid]
	if key == target {
		return mid
	}
	if left < right {
		if key < target {
			return binarySearch(nums, mid+1, right, target)
		} else {
			return binarySearch(nums, left, mid, target)
		}
	}

	return -1
}