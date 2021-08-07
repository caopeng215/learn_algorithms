func rotate(nums []int, k int)  {
  k %= len(nums)
  reverse(nums)
  reverse(nums[:k])
  reverse(nums[k:])
}

func reverse(nums []int) {
  for i, n := 0, len(nums); i < n / 2; i++ {
      nums[i],nums[n-1-i] = nums[n-1-i], nums[i]
  }
}
