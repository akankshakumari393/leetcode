// Hashing

func intersect(nums1 []int, nums2 []int) []int {
    var ans []int

    array1Content := make(map[int]int)
    for _, num := range nums1{
        if count, exist := array1Content[num]; exist {
            array1Content[num] = count +1  
        } else{
          array1Content[num] = 1  
        } 
    }
    for _, num := range nums2{
        if count, exist := array1Content[num]; exist {
            if count != 0 {
            array1Content[num] = count - 1
            ans = append(ans, num) 
            } 
        } 
    }
  return ans
}