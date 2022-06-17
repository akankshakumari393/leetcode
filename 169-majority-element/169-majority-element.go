func majorityElement(nums []int) int {
    count := 0
    candidate := 0

    for _ ,num := range nums {
        if (count == 0) {
            candidate = num
        }
        if num == candidate{
            count = count + 1             
        }else {
         count = count - 1   
        }
        fmt.Println(candidate)
    }
    return candidate
    
    
	// cnt := 0
	// for _, num := range nums {
	// 	if cnt == 0 {
	// 		candidate = num
	// 		cnt++
	// 	} else {
	// 		if candidate == num {
	// 			cnt++
	// 		} else {
	// 			cnt--
	// 		}
	// 	}
	// }
	// return candidate
}