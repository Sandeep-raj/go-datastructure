package recursion

import "log"

/*
Print all possible words from phone digits
Before the advent of QWERTY keyboards, texts and numbers were placed on the same key. For example, 2 has “ABC” if we wanted to write anything starting with ‘A’ we need to type key 2 once. If we wanted to type ‘B’, press key 2 twice and thrice for typing ‘C’.

Input number: 234
Output:
adg adh adi aeg aeh aei afg afh
afi bdg bdh bdi beg beh bei bfg
bfh bfi cdg cdh cdi ceg ceh cei
cfg cfh cfi
Explanation: All possible words which can be
formed are (Alphabetical order):
adg adh adi aeg aeh aei afg afh
afi bdg bdh bdi beg beh bei bfg
bfh bfi cdg cdh cdi ceg ceh cei
cfg cfh cfi
If 2 is pressed then the alphabet
can be a, b, c,
Similarly, for 3, it can be
d, e, f, and for 4 can be g, h, i.

Input number: 5
Output: j k l
Explanation: All possible words which can be
formed are (Alphabetical order):
j, k, l, only these three alphabets
can be written with j, k, l.
*/

var keymap map[string][]string

func keypad_comb(str string) []string {
	if len(str) == 1 {
		return keymap[str]
	}
	currChar := string(str[0])
	resArr := keypad_comb(str[1:])

	tempArr := make([]string, 0)

	for _, s := range resArr {
		for _, c := range keymap[currChar] {
			tempArr = append(tempArr, c+s)
		}
	}

	return tempArr
}

func TestKeypadComb() {
	keymap = make(map[string][]string)
	keymap["0"] = []string{".", ","}
	keymap["1"] = []string{"a", "b", "c"}
	keymap["2"] = []string{"d", "e", "f"}
	keymap["3"] = []string{"g", "h", "i"}
	keymap["4"] = []string{"j", "k", "l"}
	keymap["5"] = []string{"m", "n", "o"}
	keymap["6"] = []string{"p", "q", "r"}
	keymap["7"] = []string{"s", "t", "u"}
	keymap["8"] = []string{"v", "w", "x"}
	keymap["9"] = []string{"y", "z"}

	res := keypad_comb("78")

	for _, s := range res {
		log.Print(s)
	}
}
