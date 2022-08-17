func maxArea(height []int) int {
    area := 0
    for i, j:= 0, len(height)-1; i < j; {
        currentarea := 0
        if height[i] < height[j] {
            currentarea = height[i]* (j-i) 
            i++
        } else {
            currentarea = height[j]* (j-i)
            j--
        }
        if area < currentarea {
            area = currentarea
        }
    }
    return area
}