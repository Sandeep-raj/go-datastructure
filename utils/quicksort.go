package utils

import "log"

type QSObject struct {
	Key     float32
	Payload interface{}
}

type QS struct {
	List  []QSObject
	Count int
}

func (qs *QS) QSSort() {
	n := len((*qs).List)
	processSort(qs, 0, n-1)
}

func processSort(qs *QS, low int, high int) {
	if low >= high {
		return
	}
	partitionKey := qs.partition(low, high, &qs.List)
	if partitionKey-1 >= 0 {
		processSort(qs, low, partitionKey-1)
	}

	if partitionKey+1 < qs.Count {
		processSort(qs, partitionKey+1, high)
	}
}

func (qs *QS) partition(low int, high int, list *[]QSObject) int {
	pivot := low
	currlow := low + 1
	currhigh := high

	for currlow < currhigh {
		for currlow < len(*list) && (*list)[currlow].Key < (*list)[pivot].Key {
			currlow++
		}

		for currhigh >= 0 && (*list)[currhigh].Key > (*list)[pivot].Key {
			currhigh--
		}

		if currlow < currhigh {
			qsswap(list, currlow, currhigh)
		}
	}

	qsswap(list, pivot, currhigh)
	return currhigh
}

func qsswap(list *[]QSObject, index1 int, index2 int) {
	temp := (*list)[index1]
	(*list)[index1] = (*list)[index2]
	(*list)[index2] = temp
}

func QuickSortTest() {
	inparr := []int{10, 60, 90, 20, 30, 80, 50}
	qs := QS{
		List:  []QSObject{},
		Count: 0,
	}
	for _, val := range inparr {
		qs.List = append(qs.List, QSObject{Key: float32(val), Payload: val})
		qs.Count++
	}

	qs.QSSort()

	for _, val := range qs.List {
		log.Print(val.Key)
	}
}
