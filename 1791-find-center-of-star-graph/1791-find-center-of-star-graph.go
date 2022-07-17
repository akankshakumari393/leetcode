func findCenter(edges [][]int) int {
    count := make(map[int]int, len(edges)+1)
    for _, edge := range edges {
        count[edge[0]] = count[edge[0]] + 1
        if count[edge[0]] == len(edges) {
            return edge[0]
        }
        count[edge[1]] = count[edge[1]] + 1
        if count[edge[1]] == len(edges) {
            return edge[1]
        }
    }
    return 0
}