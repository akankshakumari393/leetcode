

func maxSubArray(nums []int) int {
    
    	f := make([]int, len(nums))
	f[0] = nums[0]
    for i := 1; i < len(nums); i++{
		f[i] = max(f[i-1] + nums[i], nums[i])
	}
// sort a slice of Int
   sort.Ints(f)
  return f[len(f)-1]
}

func max(a, b int) int {
    if a < b {
        return b
    }
    return a
}
