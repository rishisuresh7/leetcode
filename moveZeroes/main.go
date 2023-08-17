package main

func moveZeroes(nums []int)  {
    length := len(nums)
    for i := 0; i < length ; i++ {
        for j := 0; j < length -1; j++ {
            if nums[j] == 0 {
                nums[j], nums[j+1] = nums[j+1], nums[j]
            }
        }
    }
}
