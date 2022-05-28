package dynamicprog

import "log"

/*
Buy And Sell Stocks - K Transactions Allowed

1. You are given a number n, representing the number of days.
2. You are given n numbers, where ith number represents price of stock on ith day.
3. You are given a number k, representing the number of transactions allowed.
3. You are required to print the maximum profit you can make if you are allowed k transactions at-most.
Note - There can be no overlapping transaction. One transaction needs to be closed (a buy followed by a sell) before opening another transaction (another buy).

Sample Input
6
9
6
7
6
3
8
1

Sample Output
5
*/

func TestBuySellStockKTrans() {
	stockPrice := []int{10, 15, 17, 20, 16, 18, 22, 20, 22, 20, 23, 25}
	log.Print(buySellStockKTrans(stockPrice, 100))
}

func buySellStockKTrans(stockPrice []int, transCount int) int {
	priceMapTransDays := make([][]int, transCount+1)
	for i := 0; i <= transCount; i++ {
		priceMapTransDays[i] = make([]int, len(stockPrice)+1)
	}

	for trans := 1; trans <= transCount; trans++ {
		for days := 1; days <= len(stockPrice); days++ {
			priceMapTransDays[trans][days] = priceMapTransDays[trans][days-1]

			for prevdays := days - 1; prevdays > 0; prevdays-- {
				priceMapTransDays[trans][days] = maxVal(priceMapTransDays[trans][days], priceMapTransDays[trans-1][prevdays]+(stockPrice[days-1]-stockPrice[prevdays-1]))
			}
		}
	}

	return priceMapTransDays[transCount][len(stockPrice)]
}
