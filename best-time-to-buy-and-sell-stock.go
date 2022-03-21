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
