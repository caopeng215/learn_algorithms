func twoSum(nums []int, target int) []int {
  numMap := make(map[int]int, len(nums))
  for i:=0; i < len(nums); i++ {
      numMap[nums[i]] = i
  }
  for i:=0; i<len(nums); i++ {
      if index, ok := numMap[target - nums[i]]; ok && i != index {
         return []int {i, index} 
      }
  }
  numMap = nil
  return []int{}
}
