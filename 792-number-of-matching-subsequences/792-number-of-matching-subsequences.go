// Using HashMap and Queue
func numMatchingSubseq(s string, words []string) int {
    slice := []byte(s)
    hashmap := make(map[byte][]int)
    for i, element := range slice {
        hashmap[element] = append(hashmap[element], i)
    }

    count := 0
    for _, word := range words {
        prev_pos := -1
        i := 0
        for ; i<len(word); i++ {
            if queue, ok := hashmap[word[i]]; ok {
                curr_pos := getCurPosGretThanPrev(prev_pos, queue)
                if curr_pos == -1 {
                    break
                }
                prev_pos = curr_pos            
            } else {
                break
            }
            
        }
        if i == len(word) {
            count++
        }
    }    
    return count
}


func getCurPosGretThanPrev(prev_pos int, arr []int) int {
    for _, elem := range arr {
        if elem > prev_pos {
            return elem
        }
    }
    return -1
}
