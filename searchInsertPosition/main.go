package main

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if left < right {
		if nums[mid] > target {
			return binarySearch(nums, left, mid, target)
		} else {
			return binarySearch(nums, mid+1, right, target)
		}
	}

	if target > nums[right] {
		return right + 1
	}

	return right
}
