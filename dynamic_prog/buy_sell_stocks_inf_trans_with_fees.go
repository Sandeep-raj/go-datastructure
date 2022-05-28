package dynamicprog

import "log"

/*
Buy And Sell Stocks With Transaction Fee - Infinite Transactions Allowed

1. You are given a number n, representing the number of days.
2. You are given n numbers, where ith number represents price of stock on ith day.
3. You are give a number fee, representing the transaction fee for every transaction.
4. You are required to print the maximum profit you can make if you are allowed infinite transactions, but has to pay "fee" for every closed transaction.
Note - There can be no overlapping transaction. One transaction needs to be closed (a buy followed by a sell) before opening another transaction (another buy).

Sample Input
12
10
15
17
20
16
18
22
20
22
20
23
25
3

Sample Output
13
*/

// salary is a drug that they give you to forget your dreams

func TestBuySellStocksInfTransFees() {
	stockPrice := []int{10, 15, 17, 20, 16, 18, 22, 20, 22, 20, 23, 25}
	log.Print(buySellStocksInfTransFees(stockPrice, 3))
}

func buySellStocksInfTransFees(stockPrice []int, transFees int) int {
	buyPrice := make([]int, len(stockPrice))
	sellPrice := make([]int, len(stockPrice))

	buyPrice[0] = stockPrice[0]

	for i := 1; i < len(stockPrice); i++ {
		buyPrice[i] = minVal([]int{buyPrice[i-1], stockPrice[i] - sellPrice[i-1]})
		sellPrice[i] = maxVal(sellPrice[i-1], stockPrice[i]-buyPrice[i-1]-transFees)
	}

	return sellPrice[len(stockPrice)-1]
}
