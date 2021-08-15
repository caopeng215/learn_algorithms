func merge(nums1 []int, m int, nums2 []int, n int)  {
  for  p, q, len := m-1, n-1, m+n-1; q >=0 || p >= 0; len-- {
      if p < 0 {
        nums1[len] = nums2[q]
        q--
      } else if q < 0 {
        nums1[len] = nums1[p]
        p--
      } else if nums1[p] < nums2[q] {
          nums1[len] = nums2[q]
          q--
      } else {
          nums1[len] = nums1[p]
          p--
      }
  }
}
