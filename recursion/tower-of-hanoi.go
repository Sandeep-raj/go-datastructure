package recursion

import "log"

func toh(n int, tA string, tB string, tC string) {
	if n == 0 {
		return
	}
	toh(n-1, tA, tC, tB)
	log.Printf("%d [%s -> %s]", n, tA, tB)
	toh(n-1, tC, tB, tA)
}

func TestToh() {
	toh(5, "A", "B", "C")
}
