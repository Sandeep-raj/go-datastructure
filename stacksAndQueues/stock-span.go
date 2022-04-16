package stacksandqueues

import "log"

/*
Stock Span

1. You are given a number n, representing the size of array a.
2. You are given n numbers, representing the prices of a share on n days.
3. You are required to find the stock span for n days.
4. Stock span is defined as the number of days passed between the current day and the first day before today when price was higher than today.

e.g.
for the array [2 5 9 3 1 12 6 8 7]
span for 2 is 1
span for 5 is 2
span for 9 is 3
span for 3 is 1
span for 1 is 1
span for 12 is 6
span for 6 is 1
span for 8 is 2
span for 7 is 1
*/

func TestStockSpan() {
	inpArr := []int{2, 5, 9, 3, 1, 12, 6, 8, 10}
	stockSpan(inpArr)
}

func stockSpan(inpArr []int) {
	indexArr := make([]int, 0)
	resArr := make([]int, len(inpArr))

	for i := 0; i < len(inpArr); i++ {
		for len(indexArr) > 0 && inpArr[indexArr[len(indexArr)-1]] < inpArr[i] {
			indexArr = indexArr[:len(indexArr)-1]
		}

		if len(indexArr) == 0 {
			resArr[i] = i + 1
		} else {
			resArr[i] = i - indexArr[len(indexArr)-1]
		}

		indexArr = append(indexArr, i)
	}

	log.Printf("%+v", resArr)
}
