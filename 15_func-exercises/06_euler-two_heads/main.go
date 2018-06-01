package main

import "fmt"
import "math/rand"

//func loop() int {
//	//f := 0
//	//s := 0
//	var f int
//	var s int
//	var i int
//
//	for {
//		i++
//		if rand.Intn(2) == 0 {
//			f = i
//			for {
//				i++
//				if rand.Intn(2) == 0 && i + 1 == f {
//					s = i
//					return i
//				} else if rand.Intn(2) == 0 && i + 1 != f {
//					i = 0
//					loop()
//				}
//			}
//		}
//	}
//}

func main() {
	//tosses := loop()
	//fmt.Println(tosses / 10^18)
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

	 //0 is heads, 1 is tails
3
}