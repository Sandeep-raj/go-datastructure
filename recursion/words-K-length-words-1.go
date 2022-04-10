package recursion

import (
	"fmt"
	"log"
)

/*
Words - K Length Words - 1

1. You are given a word (may have one character repeat more than once).
2. You are given an integer k.
3. You are required to generate and print all k length words (of distinct chars) by using chars of the
     word.

Sample Input
aabbbccdde
3

Sample Output
abc
abd
abe
acb
adb
aeb
acd
ace
adc
aec
ade
aed
bac
bad
bae
cab
dab
eab
cad
cae
dac
eac
dae
ead
bca
bda
bea
cba
dba
eba
cda
cea
dca
eca
dea
eda
bcd
bce
bdc
bec
bde
bed
cbd
cbe
dbc
ebc
dbe
ebd
cdb
ceb
dcb
ecb
deb
edb
cde
ced
dce
ecd
dec
edc
*/

func TestWordKSelection1() {
	inpStr := "aabbbccdde"
	k := 3

	fmap := make(map[string]int)

	for i := 0; i < len(inpStr); i++ {
		if fmap[string(inpStr[i])] == 0 {
			fmap[string(inpStr[i])] = 1
		}
	}

	wordKLengthSelection1(k, fmap, "")

}

func wordKLengthSelection1(k int, fmap map[string]int, res string) {

	if k == 0 {
		log.Print(res)
		return
	}

	for key, val := range fmap {
		if val > 0 {
			fmap[key]--
			wordKLengthSelection1(k-1, fmap, fmt.Sprintf("%s%s", res, key))
			fmap[key]++
		}
	}
}
