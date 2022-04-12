package recursion

import (
	"fmt"
	"log"
	"strconv"
)

/*
Restore IP Addresses

A valid IP address consists of exactly four integers separated by single dots. Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.

Sample Input
101023

Sample Output
[1.0.10.23, 1.0.102.3, 10.1.0.23, 10.10.2.3, 101.0.2.3]

<1-255>.<0-255>.<0-255>.<0-255>
*/

func TestRestoreIpAddress() {
	inpStr := "101023"

	res := make([]string, 0)

	restoreIpAddress(4, inpStr, "", &res)

	for i := 0; i < len(res); i++ {
		log.Print(res[i])
	}

}

func restoreIpAddress(k int, inpStr string, resStr string, res *[]string) {

	if k == 0 && inpStr == "" {
		*res = append(*res, resStr[1:])
		return
	}

	for i := 0; i < len(inpStr); i++ {
		currStr := inpStr[0 : i+1]
		val, err := strconv.Atoi(currStr)

		if err == nil && currStr == fmt.Sprintf("%d", val) && val <= 255 {
			restoreIpAddress(k-1, inpStr[i+1:], fmt.Sprintf("%s.%s", resStr, currStr), res)
		}
	}
}
