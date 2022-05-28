package dynamicprog

import "log"

/*
Buy And Sell Stocks - One Transaction Allowed

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
17
*/

func TestBuyAndSell1Trans() {
	stockPrice := []int{11, 6, 7, 19, 4, 1, 6, 18, 4}
	log.Print(buySell1Trans(stockPrice))
}

func buyAndSell1TransStack(stockPrices []int) int {
	price_stack := make([]int, 0)
	maxProfit := -1
	for i := 0; i < len(stockPrices); i++ {
		if len(price_stack) > 0 {
			if stockPrices[i] >= price_stack[len(price_stack)-1] {
				price_stack = append(price_stack, stockPrices[i])
			} else {
				for len(price_stack) > 0 && price_stack[len(price_stack)-1] > stockPrices[i] {
					if maxProfit < (price_stack[len(price_stack)-1] - price_stack[0]) {
						maxProfit = (price_stack[len(price_stack)-1] - price_stack[0])
					}
					price_stack = price_stack[:len(price_stack)-1]
				}
			}
		}

		price_stack = append(price_stack, stockPrices[i])

	}

	for len(price_stack) > 0 {
		if maxProfit < price_stack[len(price_stack)-1]-price_stack[0] {
			maxProfit = price_stack[len(price_stack)-1] - price_stack[0]
		}
		price_stack = price_stack[:len(price_stack)-1]
	}
	return maxProfit
}

func buySell1Trans(stockPrice []int) int {

	lowerPrice := stockPrice[0]
	higherPrice := stockPrice[1]

	maxProfit := -1

	for i := 1; i < len(stockPrice); i++ {
		if stockPrice[i] >= higherPrice {
			higherPrice = stockPrice[i]
		} else if stockPrice[i] < lowerPrice {
			if maxProfit < higherPrice-lowerPrice {
				maxProfit = higherPrice - lowerPrice
			}
			lowerPrice = stockPrice[i]
			higherPrice = stockPrice[i]
		}
	}

	if maxProfit < higherPrice-lowerPrice {
		maxProfit = higherPrice - lowerPrice
	}

	return maxProfit
}
