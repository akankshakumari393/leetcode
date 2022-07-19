// MakeMyTrip
func longestConsecutive(nums []int) int {
    hmap := make(map[int]int, len(nums))
    for _, ele := range nums {
        hmap[ele] = 1
    }
    fmt.Println(hmap)
    longestStreak := 0
    for num, _ := range hmap {
        if _, ok := hmap[num-1]; ok {
            continue
        }
        currentNum := num
        currentStreak := 1
        
        for ;true; {
            if hmap[currentNum+1] == 1 {
                  currentNum = currentNum + 1
                  currentStreak = currentStreak + 1
            }else {
                break
            }
        }
        if currentStreak > longestStreak {
            longestStreak = currentStreak
        }
    }
    return longestStreak
}