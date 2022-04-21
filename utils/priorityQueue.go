package utils

import "errors"

type PQObj struct {
	Key  int
	Data interface{}
}

type PQ struct {
	Count int
	Queue []PQObj
}

func InitPQ() *PQ {
	return &PQ{
		Count: 0,
		Queue: nil,
	}
}

func (pq *PQ) Add(obj PQObj) {
	if pq.Count == 0 {
		pq.Queue = make([]PQObj, 0)
		pq.Queue = append(pq.Queue, obj)
	} else {
		pq.Queue = append(pq.Queue, obj)
		currIndex := len(pq.Queue) - 1
		for currIndex > 0 {
			parentIndex := int((currIndex - 1) / 2)
			if pq.Queue[parentIndex].Key > pq.Queue[currIndex].Key {
				swap(pq.Queue, parentIndex, currIndex)
				currIndex = parentIndex
			} else {
				break
			}
		}
	}

	pq.Count++
}

func (pq *PQ) Remove() (*PQObj, error) {
	if pq.Count == 0 {
		return nil, errors.New("no element present")
	} else if pq.Count == 1 {
		currval := pq.Queue[0]
		pq.Queue = nil
		pq.Count--
		return &currval, nil
	} else {
		currVal := pq.Queue[0]

		lastVal := pq.Queue[len(pq.Queue)-1]
		pq.Queue = pq.Queue[:len(pq.Queue)-1]

		pq.Queue[0] = lastVal

		PercolateDown(pq.Queue)
		pq.Count--

		return &currVal, nil
	}
}

func swap(arr []PQObj, index1 int, index2 int) {
	temp := arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

func PercolateDown(arr []PQObj) {
	currIndex := 0

	for {
		leftChild := 2*currIndex + 1
		rightChild := 2*currIndex + 2

		if leftChild > len(arr)-1 {
			return
		} else if leftChild <= len(arr)-1 && rightChild > len(arr)-1 {
			if arr[currIndex].Key <= arr[leftChild].Key {
				return
			} else {
				swap(arr, currIndex, leftChild)
				currIndex = leftChild
			}
		} else {
			tempIndex := -1
			if arr[leftChild].Key < arr[rightChild].Key {
				tempIndex = leftChild
			} else {
				tempIndex = rightChild
			}

			if arr[currIndex].Key <= arr[tempIndex].Key {
				return
			} else {
				swap(arr, currIndex, tempIndex)
				currIndex = tempIndex
			}
		}
	}

}
