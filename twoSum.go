// Solution O(n) - time complexity

func twoSum(nums []int, target int) []int {
    // declare an int slice which will have the answers 
    var ans []int
    // A map which will have number item as key and their location as value
    numToIdx := make(map[int]int)
    
    for idx, num := range nums{
        //get the expected number to present if this num is considered in result
        expected := target - num
        // check if expected item is present in the map, compare that two ids are not same
        if expectedIdx, exists := numToIdx[expected]; exists && expectedIdx != idx {
            ans = append(ans, idx, expectedIdx)
            break
        }
        // if number is not in the map then add it to the map as key
        numToIdx[num] = idx
    }   
    return ans
}
