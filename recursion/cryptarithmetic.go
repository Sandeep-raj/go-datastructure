package recursion

import (
	"fmt"
	"log"
	"strconv"
)

/*
Cryptarithmetic

Sample Input

team
pep
toppr

Sample Output
a-3 e-9 m-4 o-1 p-2 r-6 t-0
a-3 e-9 m-5 o-1 p-2 r-7 t-0
a-3 e-9 m-6 o-1 p-2 r-8 t-0
a-4 e-9 m-2 o-1 p-3 r-5 t-0
a-4 e-9 m-5 o-1 p-3 r-8 t-0
a-5 e-9 m-2 o-1 p-4 r-6 t-0
a-5 e-9 m-3 o-1 p-4 r-7 t-0
a-6 e-9 m-2 o-1 p-5 r-7 t-0
a-6 e-9 m-3 o-1 p-5 r-8 t-0
a-7 e-9 m-2 o-1 p-6 r-8 t-0
*/

func TestCryptArthmatic() {
	strArr := []string{"team", "pep", "toppr"}

	charList := make([]string, 0)
	charSet := make(map[string]int)
	charMap := make(map[int]string)

	for _, str := range strArr {
		for _, c := range str {
			if charSet[string(c)] != -1 {
				charList = append(charList, string(c))
				charSet[string(c)] = -1
			}
		}
	}

	cryptArithmatic(charList, 0, strArr, charSet, charMap)

}

func cryptArithmatic(charList []string, idx int, strArr []string, setCharMap map[string]int, charMap map[int]string) {

	if idx == len(charList) {
		val1 := getNum(strArr[0], setCharMap)
		val2 := getNum(strArr[1], setCharMap)
		val3 := getNum(strArr[2], setCharMap)

		if val1+val2 == val3 {
			for k, v := range setCharMap {
				fmt.Printf("%s-%d ", k, v)
			}
			log.Println()
		}

		return
	}

	c := charList[idx]

	for i := 0; i <= 9; i++ {
		if charMap[i] == "" {
			setCharMap[c] = i
			charMap[i] = c
			cryptArithmatic(charList, idx+1, strArr, setCharMap, charMap)
			setCharMap[c] = -1
			delete(charMap, i)
		}
	}

}

func getNum(str string, charSet map[string]int) int {
	result := ""
	for _, c := range str {
		result = result + fmt.Sprint(charSet[string(c)])
	}
	val, _ := strconv.Atoi(result)
	return val
}
