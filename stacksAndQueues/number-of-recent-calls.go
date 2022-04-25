package stacksandqueues

import (
	"log"

	"github.com/Sandeep-raj/go-datastructure/utils"
)

/*
Number Of Recent Calls

You have a RecentCounter class which counts the number of recent requests within a certain time frame.

Implement the RecentCounter class:
1. RecentCounter() Initializes the counter with zero recent requests.
2. int ping(int t) Adds a new request at time t, where t represents some time in milliseconds, and returns the number of requests that has happened in the past 3000 milliseconds (including the new request). Specifically, return the number of requests that have happened in the inclusive range [t - 3000, t].

It is guaranteed that every call to ping uses a strictly larger value of t than the previous call.

Sample Input
1
100
3001
3002
3101

Sample Output
1
2
3
3
*/

type recentCounter struct {
	count int
	llist *utils.LinkedList
}

func recentCounterInit() *recentCounter {
	return &recentCounter{
		count: 0,
		llist: utils.InitLL(),
	}
}

func (rc *recentCounter) ping(t int) int {
	node := utils.InitLLNode(t, t)
	rc.llist.AddRear(node)
	rc.count++

	terminateTime := t - 3000

	for {
		val, _ := rc.llist.PeekFront()
		if val.Data.(int) < terminateTime {
			rc.llist.RemoveFront()
			rc.count--
			continue
		} else {
			break
		}
	}

	return rc.count
}

func TestNumRecentCalls() {
	rc := recentCounterInit()
	log.Print(rc.ping(1))
	log.Print(rc.ping(100))
	log.Print(rc.ping(3001))
	log.Print(rc.ping(3002))
	log.Print(rc.ping(3101))
}
