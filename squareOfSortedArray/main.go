package main

func sortedSquares(nums []int) []int {
	length := len(nums)
	left, right := 0, length-1
	res := make([]int, length)
	for i := right; left <= right; i-- {
		rightVal := nums[right]
		leftVal := nums[left]
		if mod(rightVal) > mod(leftVal) {
			res[i] = rightVal * rightVal
			right--
		} else {
			res[i] = leftVal * leftVal
			left++
		}
	}

	return res
}

func mod(n int) int {
	if n < 0 {
		return -1 * n
	}

	return n
}
