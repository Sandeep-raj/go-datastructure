package stacksandqueues

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Moving Average From Data Stream

Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.

Implement the MovingAverage class:
1. MovingAverage(int size) Initializes the object with the size of the window size.
2. double next(int val) Returns the moving average of the last size values of the stream.

Sample Input
3
1
10
3
5

Sample Output
1.0
5.5
4.66667
6.0
*/

func TestMovingAvg() {
	aq := movingavg(3)
	aq.next(1)
	aq.next(10)
	aq.next(3)
	aq.next(5)
}

type avgqueue struct {
	capacity int
	count    int
	queue    *utils.Queue
	sum      int
}

func movingavg(k int) *avgqueue {
	return &avgqueue{
		count:    0,
		capacity: k,
		queue:    utils.InitQueue(),
		sum:      0,
	}
}

func (avgq *avgqueue) next(val int) {
	if avgq.count == avgq.capacity {
		currEle, _ := avgq.queue.Remove()
		avgq.sum = avgq.sum - currEle.(int)
		avgq.count--
	}

	avgq.queue.Add(val)
	avgq.sum = avgq.sum + val
	avgq.count++

	log.Print(float32(avgq.sum) / float32(avgq.count))

}
