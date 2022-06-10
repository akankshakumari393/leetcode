# Typical approach O(n2)
​
run two loop and compare every pair
​
# optimal apprach O(n)
## Approach:
​
* Create a variable maxProfit and mark it as 0.
* Create a variable minPrice and mark it as max_value.
* Run a for loop from 0 to n.
* Update the minPrice at if it greater than current element of the array
* Take the difference of the minPrice with the current element of the array and compare and maintain it in maxProfit.
* Return the maxProfit.
​