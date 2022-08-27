func climbStairs(n int) int {
    // base cases
    if n <= 0 {
      return 0  
    }
    if n == 1 {
      return 1  
    }
    if n == 2 {
      return 2  
    }
    
    one_step_before := 2
    two_steps_before := 1
    all_ways := 0
    
    for i:=2; i<n; i++ {
    	all_ways = one_step_before + two_steps_before
    	two_steps_before = one_step_before
        one_step_before = all_ways
    }
    return all_ways
}