package dynamicprog

import "log"

/*
Buy And Sell Stocks With Cooldown - Infinite Transaction Allowed

1. You are given a number n, representing the number of days.
2. You are given n numbers, where ith number represents price of stock on ith day.
3. You are required to print the maximum profit you can make if you are allowed infinite transactions, but have to cooldown for 1 day after 1 transaction
i.e. you cannot buy on the next day after you sell, you have to cooldown for a day at-least before buying again.
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

Sample Output
19
*/

func TestBuySellStockInfTransCooldown() {
	stockPrice := []int{10, 15, 17, 20, 16, 18, 22, 20, 22, 20, 23, 25}
	log.Print(buySellStockInfTransCooldown(stockPrice))
}

func buySellStockInfTransCooldown(stockPrice []int) int {
	buyPrice := make([]int, len(stockPrice))
	sellPrice := make([]int, len(stockPrice))
	cooldownPrice := make([]int, len(stockPrice))

	buyPrice[0] = stockPrice[0]

	for i := 1; i < len(stockPrice); i++ {
		buyPrice[i] = minVal([]int{buyPrice[i-1], stockPrice[i] - cooldownPrice[i-1]})
		sellPrice[i] = maxVal(sellPrice[i-1], stockPrice[i]-buyPrice[i-1])
		cooldownPrice[i] = sellPrice[i-1]
	}
	return sellPrice[len(stockPrice)-1]
}
