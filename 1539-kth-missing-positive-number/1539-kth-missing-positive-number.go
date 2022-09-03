// Apply a binary search. Since the array is sorted we can find at any given index how many numbers are missing as arr[index] – (index+1)
func findKthPositive(arr []int, k int) int {
    low := 0
    high := len(arr) - 1
    var mid int
    for low <= high {
      mid = low + (high-low) / 2
        missing_number := arr[mid] - (1 + mid)   
      if missing_number >= k {  //A[m]-(m+1)   --> This gives umber of missing number before m'th index because m+1 is the actual value that should be there
          high = mid - 1
       } else {
          low = mid + 1   
       }
    }
    //return arr[high] + k - (arr[high] - (1 + high));   // or simply low+k
    return low+k
}