func merge(intervals [][]int) [][]int {    
   	if len(intervals) == 0 {
		return [][]int{}
	}
    var result [][]int

    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
    })
    
    in := intervals[0] 
    for i:=1;i<len(intervals); i++ {
        if in[1] >= intervals[i][0] {
            in[1] = int(math.Max(float64(in[1]), float64(intervals[i][1])))
            
        } else {
            result = append(result, in)
            in = intervals[i]
        }
    }
    return append(result, in) 
}