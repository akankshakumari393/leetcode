func canFinish(numCourses int, prerequisites [][]int) bool {
    // figure out the length
	length := len(prerequisites)
    // return true if zero length
	if numCourses == 0 || length == 0 {
		return true
	}

	counter := make([]int, numCourses)
	for i := 0; i < length; i++ {
		counter[prerequisites[i][0]]++
	}
	fmt.Println("counter", counter)
	var queue []int
    // if no dependency or free node
	for i := 0; i < numCourses; i++ {
		if counter[i] == 0 {
			queue = append(queue, i)
		}
	}

	noPre := len(queue)
	for {
		if len(queue) == 0 {
			break
		}
        // pop one element from queue
		n := queue[0]
		queue = queue[1:]
         
		for i := 0; i < length; i++ {
            // since are visiting n so remove dependency wherever it is present
			if prerequisites[i][1] == n {
				counter[prerequisites[i][0]]--
                // if the dependent element are all visited add the element to queue
				if counter[prerequisites[i][0]] == 0 {
					noPre++
					queue = append(queue, prerequisites[i][0])
				}
			}
		}
	}

	return noPre == numCourses
}