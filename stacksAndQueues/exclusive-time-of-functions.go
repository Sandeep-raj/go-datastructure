package stacksandqueues

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
Exclusive Time Of Functions

1: On a single-threaded CPU, we execute a program containing n functions. Each function has a unique ID between 0 and n-1.
2: You are required to find the total execution time of each function.
3: You are given len number of logs, where logs[i] represents the ith log message formatted as a string "{function_id}:{"start" | "end"}:{timestamp}", telling start or end time of function with id function_id.

Note that a function can be called multiple times, possibly recursively.

Sample Input
2
4
0:start:0
1:start:2
1:end:5
0:end:6

Sample Output
3
4
*/

func TestExclusiveTimeFunc() {
	fn := 2
	logs := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}
	ex_time := make([]int, fn)

	exTimeFunc(logs, ex_time)
}

func exTimeFunc(logs []string, ex_time []int) {
	stack := make([]string, 0)

	for i := 0; i < len(logs); i++ {
		currData := strings.Split(logs[i], ":")

		if currData[1] == "end" {
			tempRes := 0
			for len(stack) > 0 {
				tempCurrData := strings.Split(stack[len(stack)-1], ":")
				if tempCurrData[1] == "start" && tempCurrData[0] == currData[0] {
					index, _ := strconv.Atoi(tempCurrData[0])
					startTime, _ := strconv.Atoi(tempCurrData[2])
					endTime, _ := strconv.Atoi(currData[2])
					ex_time[index] = ex_time[index] + (endTime - startTime + 1) - tempRes
					stack = stack[:len(stack)-1]
					stack = append(stack, fmt.Sprintf("%d:%s:%d", -1, "res", (endTime-startTime+1)))
					break
				} else {
					val, _ := strconv.Atoi(tempCurrData[2])
					tempRes = tempRes + val
				}
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, logs[i])
		}
	}

	log.Printf("%+v", ex_time)
}
