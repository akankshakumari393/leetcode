func searchMatrix(matrix [][]int, target int) bool {
    lenRow := len(matrix)
    if lenRow == 1 {
    return binarySearch(matrix[0], target)        
    }
    
    for i := 0; i< lenRow-1;i++ {
        if target < matrix[i+1][0] {
            return binarySearch(matrix[i], target)
        }
    }
    return binarySearch(matrix[lenRow-1], target)
}

func binarySearch(arr []int, target int) bool {
    low := 0
    high := len(arr) -1
    for low <= high {
        mid := (low+high)/2

        if arr[mid] == target {
            return true
        }
        if target > arr[mid]{
            low= mid + 1
        }else if target < arr[mid] {
            high = mid-1
        }
    }
    return false
}