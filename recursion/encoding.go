package recursion

import "log"

/*
Print Encodings
1. You are given a string str of digits. (will never start with a 0)
2. You are required to encode the str as per following rules
    1 -> a
    2 -> b
    3 -> c
    ..
    25 -> y
    26 -> z
3. Complete the body of printEncodings function - without changing signature - to calculate and print all encodings of str.
Use the input-output below to get more understanding on what is required
123 -> abc, aw, lc
993 -> iic
013 -> Invalid input. A string starting with 0 will not be passed.
103 -> jc
303 -> No output possible. But such a string maybe passed. In this case print nothing.
*/

var charMap map[string]string

func encoding(str string, currStr string, res *[]string) {
	if str == "" {
		*res = append(*res, currStr)
		return
	}

	if charMap[string(str[0])] != "" {
		encoding(str[1:], currStr+charMap[string(str[0])], res)
	} else {
		return
	}

	if len(str) > 1 && charMap[string(str[0:2])] != "" {
		encoding(str[2:], currStr+charMap[string(str[0:2])], res)
	} else {
		return
	}
}

func TestEncoding() {
	charMap = make(map[string]string)
	charMap["1"] = "A"
	charMap["2"] = "B"
	charMap["3"] = "C"
	charMap["4"] = "D"
	charMap["5"] = "E"
	charMap["6"] = "F"
	charMap["7"] = "G"
	charMap["8"] = "H"
	charMap["9"] = "I"
	charMap["10"] = "J"
	charMap["11"] = "K"
	charMap["12"] = "L"
	charMap["13"] = "M"
	charMap["14"] = "N"
	charMap["15"] = "O"
	charMap["16"] = "P"
	charMap["17"] = "Q"
	charMap["18"] = "R"
	charMap["19"] = "S"
	charMap["20"] = "T"
	charMap["21"] = "U"
	charMap["22"] = "V"
	charMap["23"] = "W"
	charMap["24"] = "X"
	charMap["25"] = "Y"
	charMap["26"] = "Z"

	res := make([]string, 0)
	encoding("993", "", &res)
	for _, s := range res {
		log.Print(s)
	}
}
