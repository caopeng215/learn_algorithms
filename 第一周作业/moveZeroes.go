func moveZeroes(nums []int)  {
  j := 0
  for i := 0; i < len(nums); i++ {
    if nums[i] != 0 {
      tmp := nums[i]
      nums[i] = nums[j]
      nums[j] = tmp
      j++
    }
  }
}
