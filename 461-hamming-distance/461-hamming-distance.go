func hammingDistance(x int, y int) int {
    // use XOR operation 
    // 00 - 0
    // 01 - 1
    // 10 - 1
    // 11 - 0
    result := 0
    for ;x!=0 || y!=0; {
        result = result + (x%2 ^ y%2)
        x = x>>1
        y = y>>1
    }
    return result
}