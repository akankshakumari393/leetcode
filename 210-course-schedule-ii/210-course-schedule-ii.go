func findOrder(numCourses int, prerequisites [][]int) []int {
    var ans []int
    
    // figure out the length
	length := len(prerequisites)
    // return true if zero length as no prerequisite
	if length == 0 {
      for i := 0; i < numCourses; i++ {
          ans = append(ans, i)
      }
    return ans
    }
    counter := make([]int, numCourses)
    for i:=0;i<len(prerequisites); i++ {
        counter[prerequisites[i][0]]++
    }
    var queue []int
    for i:=0;i<numCourses; i++ {
        if counter[i] == 0 {
            queue = append(queue, i)
        }
    }
        fmt.Print("printing queue", queue)
    
    noPre := len(queue)
    
    for {
    		if len(queue) == 0 {
			break
		}
    
        // pop from queue
        n:= queue[0]
        queue = queue[1:]
        // fmt.Print("printing ans", ans)
        for i:=0; i< len(prerequisites); i++ {
            if prerequisites[i][1] == n {
                counter[prerequisites[i][0]]--
                if counter[prerequisites[i][0]] == 0 {
                 queue = append(queue, prerequisites[i][0])
                 noPre++
                }
            }
        }
       ans = append(ans, n)

    }
    if noPre == numCourses {
        return ans
    }
    return []int{}

}