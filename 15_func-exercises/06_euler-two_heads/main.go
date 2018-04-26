package main

import "fmt"

func main() {
	// Q(a/b, p) = q
	// a≡bq(modp)
	// Q(3/5,109)

	//(5 * q) ≡ 3 (mod 109)
	//q := 1;
	//
	//for {
	//	if 5 * q - 109 * 3 == 3 {
	//		fmt.Println(q, "is q, the smallest positive number")
	//		return
	//	} else {
	//		q++
	//	}
	//}

	//for {
	//	if 5 * q - 109 * 3 == 3 {
	//		fmt.Println(q, "is q, the smallest positive number")
	//		return
	//	} else {
	//		q++
	//	}
	//}

	q2 := 1;

	for {
		a := 9 % 109
		b := (31 * q2) % 109
		if a == b {
			fmt.Println(q2, "is q, the smallest positive number")
			return
		} else {
			q2++
		}
	}
}