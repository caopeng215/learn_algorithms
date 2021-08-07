func trap(height []int) int {
  sum := 0
  left := 1
  right := len(height) - 2
  max_left := 0
  max_right := 0
  for i:=1; i <len(height)-1; i++ {
    if height[left-1] < height[right+1] {
        max_left = max(height[left-1], max_left)
        if max_left > height[left] {
            sum += (max_left - height[left])
        }
        left++
    } else {
        max_right = max(height[right+1], max_right)
        if max_right > height[right] {
            sum += (max_right - height[right])
        }
        right--
    }
  }
  return sum
}
func max(x,y int) int {
    if x > y {
        return x
    }
    return y
}
