package dynamicprog

import (
	"log"
	"math"
)

/*
Buy And Sell Stocks - Two Transactions Allowed

1. You are given a number n, representing the number of days.
2. You are given n numbers, where ith number represents price of stock on ith day.
3. You are required to print the maximum profit you can make if you are allowed two transactions at-most.
Note - There can be no overlapping transaction. One transaction needs to be closed (a buy followed by a sell) before opening another transaction (another buy).

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

func TestBuySellStocks2Trans() {
	stockPrice := []int{11, 6, 7, 19, 4, 1, 6, 18, 4}
	log.Print(buySellStock2Trans(stockPrice))
}

func buySellStock2Trans(stockPrice []int) int {

	firstBuy := math.MaxInt
	firstSell := 0
	secondBuy := math.MaxInt
	secondSell := 0

	for i := 0; i < len(stockPrice); i++ {
		firstBuy = minVal([]int{firstBuy, stockPrice[i]})
		firstSell = maxVal(firstSell, stockPrice[i]-firstBuy)
		secondBuy = minVal([]int{secondBuy, stockPrice[i] - firstSell})
		secondSell = maxVal(secondSell, stockPrice[i]-secondBuy)
	}

	return secondSell
}
