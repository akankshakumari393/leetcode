## Solution 1:Brute force
​
* Approach: Take another dummy matrix of n*n, and then take the first row of the matrix and put it in the last column of the dummy matrix, take the second row of the matrix, and put it in the second last column of the matrix and so.
​
Space Complexity - O(N*N)
Time Complexity - O(N*N)
​
`rotated[j][n - i - 1] = matrix[i][j]`
​
## Solution 2: Optimized approach
​
* Intuition: By observation, we see that the first column of the original matrix is the reverse of the first row of the rotated matrix, so that’s why we transpose the matrix and then reverse each row, and since we are making changes in the matrix itself space complexity gets reduced to O(1).
​
* Approach: /
​
**Step1:** Transpose of the matrix. (transposing means changing columns to rows and rows to columns) /
​
**Step2:** Reverse each row of the matrix.
​