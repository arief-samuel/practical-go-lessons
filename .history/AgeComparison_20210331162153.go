package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")
	fmt.Println("It is", ageJohn > agePaul, "that John is older than Paul")
	fmt.Println("It s", ageJohn == agePaul, " that John and Paul have same age")

	// if agePaul > ageJohn {
	// 	fmt.Println("Paul is older than John")
	// } else {
	// 	// fmt.Println("Paul is younger than John, or both have the same age")
	// 	if agePaul == ageJohn {
	// 		fmt.Println("Paul and John have the same age")
	// 	} else {
	// 		fmt.Println("Paul is younger than John")
	// // 	}
	// }

	// if without else
	// if agePaul > ageJohn {
	// 	fmt.Println("Paul is older than John")
	// }
	// fmt.Println("End of Program")

	// if/elseif/else
	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age ")
	} else {
		fmt.Println("Paul is younger than John")
	}

	// switch expression
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old ")
	case 20:
		fmt.Println("John is 20 years old ")
	case 100:
		fmt.Println("John is 100 years old ")
	default:
		fmt.Println("John has neither 10,20 nor 100 years old")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: //*\label{switchMulti}
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 time agePaul")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}
	// For statement with a single condition
	const emailToSend = 10000010
	emailSent := 10000000
	for emailSent <= emailToSend {
		fmt.Println("sending email..", emailSent)
		emailSent++ //*\label{forWithSingle2}
	}
	fmt.Println("end of program")
	// The initialization statement | The condition | The post statement
	for i := 0; i < ageJohn; i++ {
		fmt.Println("Iteration N", i)
	}
}
