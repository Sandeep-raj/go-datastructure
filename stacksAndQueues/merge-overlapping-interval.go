package stacksandqueues

import (
	"log"
	"math"
)

/*
Merge Overlapping Interval

1. You are given a number n, representing the number of time-intervals.
2. In the next n lines, you are given a pair of space separated numbers.
3. The pair of numbers represent the start time and end time of a meeting (first number is start time and second number is end time)
4. You are required to merge the meetings and print the merged meetings output in increasing order of start time.

E.g. Let us say there are 6 meetings
1 8
5 12
14 19
22 28
25 27
27 30

Then the output of merged meetings will belongs
1 12
14 19
22 30
*/

func TestMergeOverlapInterval() {
	meetings := make([][]int, 0)
	meetings = append(meetings, []int{1, 8})
	meetings = append(meetings, []int{5, 12})
	meetings = append(meetings, []int{14, 19})
	meetings = append(meetings, []int{22, 28})
	meetings = append(meetings, []int{25, 27})
	meetings = append(meetings, []int{27, 30})

	// meetings = append(meetings, []int{2, 4})
	// meetings = append(meetings, []int{5, 7})
	// meetings = append(meetings, []int{1, 3})
	// meetings = append(meetings, []int{6, 8})
	// meetings = append(meetings, []int{25, 27})
	// meetings = append(meetings, []int{27, 30})

	mergeOverlapIntervals(meetings)
}

func mergeOverlapIntervals(meetings [][]int) {
	meetingStack := make([][]int, 0)

	for i := 0; i < len(meetings); i++ {
		if len(meetingStack) == 0 {
			meetingStack = append(meetingStack, meetings[i])
		} else {
			if meetings[i][1]+1 < meetingStack[len(meetingStack)-1][0] {
				meetingStack = append(meetingStack, meetings[i])
			} else {
				tempStack := make([][]int, 0)
				for len(meetingStack) > 0 && meetings[i][0] > meetingStack[len(meetingStack)-1][1]+1 {
					tempStack = append(tempStack, meetingStack[len(meetingStack)-1])
					meetingStack = meetingStack[:len(meetingStack)-1]
				}

				if len(meetingStack) > 0 {
					if meetings[i][1]+1 < meetingStack[len(meetingStack)-1][0] {
						meetingStack = append(meetingStack, meetings[i])
					} else {
						tempmeet := meetingStack[len(meetingStack)-1]
						meetingStack = meetingStack[:len(meetingStack)-1]
						meetingStack = append(meetingStack, getLimits(tempmeet, meetings[i]))
					}
				} else {
					meetingStack = append(meetingStack, meetings[i])
				}

				for len(tempStack) > 0 {
					meetingStack = append(meetingStack, tempStack[len(tempStack)-1])
					tempStack = tempStack[:len(tempStack)-1]
				}
			}
		}
	}

	for len(meetingStack) > 0 {
		log.Print(meetingStack[len(meetingStack)-1])
		meetingStack = meetingStack[:len(meetingStack)-1]
	}
}

func getLimits(meet1 []int, meet2 []int) []int {
	lowerLimit := math.MaxInt
	higherLimit := -1 * math.MaxInt

	for i := 0; i < 2; i++ {
		if meet1[i] < lowerLimit {
			lowerLimit = meet1[i]
		}

		if meet1[i] > higherLimit {
			higherLimit = meet1[i]
		}

		if meet2[i] < lowerLimit {
			lowerLimit = meet2[i]
		}

		if meet2[i] > higherLimit {
			higherLimit = meet2[i]
		}
	}

	return []int{lowerLimit, higherLimit}
}
