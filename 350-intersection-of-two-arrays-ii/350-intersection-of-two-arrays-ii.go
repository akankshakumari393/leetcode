// Hashing
//Time Complexity O(n)
// Space complexity - O(n) It uses a map for first array, and array to store answer  
// Populate a map which contains the element of first array as key and the count of it a value.
// Traverse through the second array, if the element exist in Map and it count >0 then add the //number to answer and decrese the count.


func intersect(nums1 []int, nums2 []int) []int {
    var ans []int

    array1Content := make(map[int]int)
    for _, num := range nums1{
        array1Content[num] = array1Content[num] +1
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