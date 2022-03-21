// Brute Force approach -- Time limit exceeded

//Time complexity - O(n^2)
//Space complexity - O(1)
//Suppose Buying a stock at Ith day and selling it at jth day while i alwyas < j.
func maxProfit(prices []int) int {
 var maxProfit int    
for i:=0; i< len(prices); i++ {
      for j:=i+1; j < len(prices); j++ {
          profit := prices[j] - prices [i]
          if profit > maxProfit {
              maxProfit = profit
          } 
      }
   }
   return maxProfit
}

// Optimized approach

//Time complexity - O(n)
//Space complexity - O(n)

// By Traversing Backward get the maximum amount by selling stock on that particular day
func maxProfit(prices []int) int {
	var maxProfit int
	var profit int
	maxSell := make([]int, len(prices))
	maxSell[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {

		if maxSell[i+1] > prices[i] {
			maxSell[i] = maxSell[i+1]
		} else {
			maxSell[i] = prices[i]
		}
	}
	// By now we have a slice maxSell which specify that on a particular day what would be the maximum value of the stock in future if we plan to buy on that day
	// we now have to calculate the profit and keep the maximum in a variable
	for i := 0; i < len(prices); i++ {
		// calculate profit by buying on ith day, we already know the max price of that stock in future
		profit = maxSell[i] - prices[i]
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// Optimized approach

//Time complexity - O(n)
//Space complexity - O(1)
// store MinValue So Far in a var and compare it with every element
func maxProfit(prices []int) int {
    maxProfit := 0
    profit := 0 
    minSoFar := prices[0]
    for i := 1; i < len(prices); i++ {

       if prices[i] < minSoFar {
			minSoFar = prices[i]
		}
      profit = prices[i] - minSoFar
		if maxProfit < profit {
			maxProfit = profit
		}

 
	}
	return maxProfit
}

