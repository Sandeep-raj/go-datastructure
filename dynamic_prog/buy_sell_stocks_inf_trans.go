package dynamicprog

import "log"

/*
Buy And Sell Stocks - Infinite Transactions Allowed

1. You are given a number n, representing the number of days.
2. You are given n numbers, where ith number represents price of stock on ith day.
3. You are required to print the maximum profit you can make if you are allowed infinite transactions.
Note - There can be no overlapping transaction. One transaction needs to be closed (a buy followed by a sell) before opening another transaction (another buy)

Sample Input
9
11
6
7
19
4
1
6
18
4

Sample Output
30
*/

func TestBuySellStocksInfTrans() {
	stockPrice := []int{11, 6, 7, 19, 4, 1, 6, 18, 4}
	log.Print(buySellInfTrans(stockPrice))
}

func buySellInfTrans(stockPrice []int) int {
	maxProfit := 0
	for i := 1; i < len(stockPrice); i++ {
		if stockPrice[i]-stockPrice[i-1] > 0 {
			maxProfit += stockPrice[i] - stockPrice[i-1]
		}
	}
	return maxProfit
}
